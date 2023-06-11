# golang
<h2>A place for my Go tutorials, etc.</h2>

The Go <a href="https://go.dev/doc/tutorial/" target="golangtuts">tutorials</a> demonstrated separately how to create and access modules, how to connect to a postgres database, and how to create api endpoints. I am combining the three lessons for the "Albums" project.
<br /><br />This application will create a schema named ``tutorial_sandbox`` in the postgres database indicated by the operating system environment variables described below.  The schema will contain a single table named ``tutorial_sandbox.album``, which we can query and modify via the API endpoints we are creating.
<br /><br />***Prerequisites***: 
<br />1. <a href="https://go.dev/doc/install" target="goinstall">Go must be installed</a>
<br />2. The user must have write access credentials to a <a href="https://www.postgresql.org/download/" target="pgdl">postgres</a> instance.
<br /><br />***Usage instructions***:
<br />1. Set the following operating system environment variables to values that are appropriate for the postgres instance you will be using (using environment variables avoids the need for hard-coding the server connection info):
<br />&nbsp;&nbsp;&nbsp;&nbsp;``PQHOST`` (defaults to "localhost" if the ``PQHOST`` environment variable is empty or not found)
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``PQUSER`` (defaults to "postgres" if the ``PQUSER`` environment variable is empty or not found)
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``PQPW`` (defaults to "password" if the ``PQPW`` environment variable is empty or not found)
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``PQDB`` (defaults to "postgres" if the ``PQDB`` environment variable is empty or not found)
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``PQSSL`` (defaults to "disable" if the ``PQSSL`` environment variable is empty or not found)
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``PQPORT`` (defaults to 5432 if the ``PQPORT`` environment variable is empty or not found)
    <br /><br />To set the environment variables, open a Terminal or Command Prompt window.  Linux/Bash Terminal uses the ``export`` command, and Windows Command Prompt uses the ``set`` command.
<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;***Examples for setting environment variables in Windows and in Linux/Mac***:
<br />&nbsp;&nbsp;&nbsp;&nbsp;Set the ``PQHOST`` environment variable to "mypostgres.example.com", ``PQUSER`` to "test.username", and ``PQPW`` to "test.pw":
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In Windows, you would type the following into a Command Prompt window after the ``>`` prompt:
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``set PQHOST=mypostgres.example.com``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``set PQUSER=test.username``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``set PQPW=test.pw``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In a Linux/Bash Terminal (Mac), you would type the following into a Terminal window after the ``$`` prompt:
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``export PQHOST=mypostgres.example.com``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``export PQUSER=test.username``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``export PQPW=test.pw``
    <br /><br />2. After setting all of the necessary environment variables, in the same Terminal or Command Prompt window, navigate to ``./albums/web-service-gin`` and type the following commands at the prompt (``$`` or ``>``):
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go mod init ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go work use . ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go mod edit -replace example.com/dbaccess=../dbaccess ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go mod tidy ``
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``go run . ``
<br /><br />3. Use **curl** in a ***separate*** Terminal or Command Prompt window to select and modify values.
<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;***Examples for selecting and modifying data with curl***:
<br />&nbsp;&nbsp;&nbsp;&nbsp;Get all albums:
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``curl http://localhost:8080/albums --header "Content-Type: application/json" --request "GET" ``
<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;Get the album having ID 3:
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``curl http://localhost:8080/albums/3 --header "Content-Type: application/json" --request "GET" ``
<br /><br />&nbsp;&nbsp;&nbsp;&nbsp;Post an insert/update:
    <br />&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;``curl http://localhost:8080/albums  --include  --header "Content-Type: application/json" --request "POST"  --data "{ \"id\": 4,\"title\": \"The Modern Sound of Betty Carter\",\"artist\": \"Betty Carter\",\"price\": 56.78}" ``
<br /><br />4. In a postgres session connected to the instance indicated by the environment variable settings that were configured in step 1 above, run the following query to verify any inserts/updates:
    <br />&nbsp;&nbsp;&nbsp;&nbsp;``select * from tutorial_sandbox.album;``
<br /><br />Notes about the jazz album project itself: &nbsp;a real record store would have far more properties for albums, as the price would be dependent on the condition of the vinyl, the condition of the cover/sleeve/insert, whether or not it is an original pressing, etc.  There is not even a quantity.  I decided to stop overthinking it after a while.
