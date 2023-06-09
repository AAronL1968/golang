# golang
<h2>A place for my Go tutorials, etc.</h2>

The Go <a href="https://go.dev/doc/tutorial/" target="golangtuts">tutorials</a> demonstrated separately how to create and access modules, how to connect to a postgres database, and how to create api endpoints. I am combining the three lessons for the "Albums" project.

Usage (assuming <a href="https://go.dev/doc/install" target="goinstall">Go is installed</a>):<br /><br />
Edit ``./albums/dbaccess/dbaccess.go`` and modify the constants to appropriate values for your postgres instance (``line 10``)<br />
    <h4>&nbsp;&nbsp;``const ( ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``host = "your.hostname.here" ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``port = 5432 ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``user = "your.username" ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``password = "your.password" ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``dbname = "your.database.name" ``
   <br />``)``</h4>  
In a Terminal or Command Prompt window, navigate to ``./albums/web-service-gin`` and type the following at the prompt (``$`` or ``>``):
    <h4>&nbsp;&nbsp;&nbsp;&nbsp;``go work use . ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go mod tidy ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go run . ``</h4>
Use curl in another Terminal or Command Prompt window to select and modify values, then verify with database selects.
<br /><br />Get all albums:
    <h4>&nbsp;&nbsp;``curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET" ``</h4>
<br /><br />Get the album having ID 3:
    <h4>&nbsp;&nbsp;``curl http://localhost:8080/albums/3 --header "Content-Type: application/json" --request "GET" ``</h4>
<br /><br />Post an insert/update:
    <h4>&nbsp;&nbsp;``curl http://localhost:8080/albums  --include  --header "Content-Type: application/json" --request "POST"  --data "{ \"id\": 4,\"title\": \"The Modern Sound of Betty Carter\",\"artist\": \"Betty Carter\",\"price\": 49.99}" ``</h4>
<br /><br />In a postgres session connected to your instance (indicated in ``./albums/dbaccess/dbaccess.go``) run the following query:
    <h4>&nbsp;&nbsp;``select * from tutorial_sandbox.album;``</h4>
