package dbaccess

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func getPGDB() (*sql.DB, error) {
	pqhost, exists := os.LookupEnv("PQHOST")
	if pqhost == "" || !exists {
		pqhost = "localhost"
	}
	pquser, exists := os.LookupEnv("PQUSER")
	if pquser == "" || !exists {
		pquser = "postgres"
	}
	pqpw, exists := os.LookupEnv("PQPW")
	if pqpw == "" || !exists {
		pqpw = "password"
	}
	pqdb, exists := os.LookupEnv("PQDB")
	if pqdb == "" || !exists {
		pqdb = "postgres"
	}
	pqssl, exists := os.LookupEnv("PQSSL")
	if pqssl == "" || !exists {
		pqssl = "disable"
	}
	pqport, exists := os.LookupEnv("PQPORT")
	if pqport == "" || !exists {
		pqport = "5432"
	}
	pqportInt, err := strconv.Atoi(pqport)
	if err != nil {
		pqportInt = 5432
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		pqhost, pqportInt, pquser, pqpw, pqdb, pqssl)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}

func AllAlbums() ([]Album, error) {
	var albums []Album
	db, err := getPGDB()
	defer db.Close()
	rows, err := db.Query("select * from tutorial_sandbox.album")
	if err != nil {
		return nil, fmt.Errorf("AllAlbums %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AllAlbums %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AllAlbums %v", err)
	}
	return albums, nil
}

func AlbumsByArtist(artist string) ([]Album, error) {
	var albums []Album
	db, err := getPGDB()
	defer db.Close()
	rows, err := db.Query("select * from tutorial_sandbox.album where artist = $1", artist)
	if err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", artist, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumsByArtist %q: %v", artist, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", artist, err)
	}
	return albums, nil
}

func findByArtistAndTitle(artist string, album string) int {
	db, err := getPGDB()
	defer db.Close()
	var ct int
	var id int
	row := db.QueryRow("select count(*) as ct, max(id) as id from tutorial_sandbox.album where artist = $1 and title = $2", artist, album)
	if err != nil {
		return 0
	}
	switch err := row.Scan(&ct, &id); err {
	case sql.ErrNoRows:
		return 0
	case nil:
		return id
	default:
		return id
	}
}

func AlbumById(id int) (Album, error) {
	var alb Album
	db, err := getPGDB()
	defer db.Close()
	row := db.QueryRow("select * from tutorial_sandbox.album WHERE id = $1", id)
	if err != nil {
		return alb, fmt.Errorf("AlbumById %q: %v", id, err)
	}
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("AlbumById %d: no such album", id)
		}
		return alb, fmt.Errorf("AlbumById %d: %v", id, err)
	}
	return alb, nil
}

func InitSchema() (int, error) {
	strSQL := "create schema if not exists tutorial_sandbox;"
	db, err := getPGDB()
	defer db.Close()
	_, err = db.Exec(strSQL)
	if err != nil {
		return 0, fmt.Errorf("InitSchema create schema %v", err)
	}
	strSQL = "drop table if exists tutorial_sandbox.album;"
	_, err = db.Exec(strSQL)
	if err != nil {
		return 0, fmt.Errorf("InitSchema drop table %v", err)
	}
	strSQL = "create table if not exists tutorial_sandbox.album ( id serial not null, title varchar(128) not null, artist varchar(255) not null, price decimal(5,2) not null, constraint pk_album primary key ( id ) );"
	_, err = db.Exec(strSQL)
	if err != nil {
		return 0, fmt.Errorf("InitSchema create table %v", err)
	}
	strSQL = "insert into tutorial_sandbox.album ( title, artist, price ) values ('Blue Train', 'John Coltrane', 56.99), ('Giant Steps', 'John Coltrane', 63.99), ('Jeru', 'Gerry Mulligan', 17.99), ('Sarah Vaughan', 'Sarah Vaughan', 34.98);"
	_, err = db.Exec(strSQL)
	if err != nil {
		return 0, fmt.Errorf("InitSchema insert %v", err)
	}
	return 1, nil
}

func UpsertAlbum(alb Album) (int, string, error) {
	var id int
	operation := "insert"
	existingId := findByArtistAndTitle(alb.Artist, alb.Title)
	db, err := getPGDB()
	defer db.Close()
	if existingId == 0 {
		result := db.QueryRow("insert into tutorial_sandbox.album (title, artist, price) values ($1, $2, $3) returning id", alb.Title, alb.Artist, alb.Price).Scan(&id)
		if err != nil {
			return 0, operation, fmt.Errorf("UpsertAlbum insert: %v", err)
		}
		if id != 0 {
			return id, operation, nil
		}
		if result == nil {
			return 0, operation, nil
		}
		return id, operation, nil
	}
	result := db.QueryRow("update tutorial_sandbox.album set price = $1 where id = $2 returning id", alb.Price, existingId).Scan(&id)
	operation = "update"
	if err != nil {
		return 0, operation, fmt.Errorf("UpsertAlbum update: %v", err)
	}
	if id != 0 {
		return id, operation, nil
	}
	if result == nil {
		return 0, operation, nil
	}
	return id, operation, nil
}
