package main

import (
    "net/http"
    "fmt"
    "strconv"
    "example.com/dbaccess"

    "github.com/gin-gonic/gin"
)

func main() {
    result, err := dbaccess.InitSchema()
    if err != nil {
      fmt.Println(err.Error())
    }
    if result < 1 {
      fmt.Println("Error in InitSchema")
    }
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    albums, err := dbaccess.AllAlbums()
    if err != nil {
      c.IndentedJSON(http.StatusNotFound, gin.H{"getAlbums message": "Error retrieving AllAlbums"})
      return
    }
    c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")
    idInt, err := strconv.Atoi(id)
    if err != nil {
      c.IndentedJSON(http.StatusBadRequest, gin.H{"getAlbumById message": "id not an integer"})
      return
    }
    alb, err := dbaccess.AlbumById(idInt)
    if err != nil {
      c.IndentedJSON(http.StatusNotFound, gin.H{"getAlbumById message": "Error retrieving AlbumById"})      
    }
    c.IndentedJSON(http.StatusOK, alb)
    return
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum dbaccess.Album
    albums, err := dbaccess.AllAlbums()
    if err != nil {
      c.IndentedJSON(http.StatusNotFound, gin.H{"postAlbums message": "Error retrieving AllAlbums"})
      return
    }
    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    id, operation, err := dbaccess.UpsertAlbum(newAlbum)
    if err != nil {
      if operation == "insert" {
        c.IndentedJSON(http.StatusFailedDependency, gin.H{"postAlbums message": "Error inserting album"})
        return
      }
      c.IndentedJSON(http.StatusFailedDependency, gin.H{"postAlbums message": "Error updating album"})
      return
    }
    if id < 0 {
      if operation == "insert" {
        c.IndentedJSON(http.StatusFailedDependency, gin.H{"postAlbums message": "Error inserting album - invalid id"})
        return
      }
      c.IndentedJSON(http.StatusFailedDependency, gin.H{"postAlbums message": "Error updating album - invalid id"})
      return
    }
    c.IndentedJSON(http.StatusCreated, newAlbum)
    return
}