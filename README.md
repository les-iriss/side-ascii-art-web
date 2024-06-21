#ASCII-ART-WEB
##Description

ascii-art-web is a web application that creates and runs a server, allowing users to interact with a web GUI version of terminal-based ASCII art.
Authors

    Oussama Benali
    Ziad ...
    Hamza ...

##Usage

To start the server, run the following bash command:

```bash

go run main.go
```
Then, open your browser and navigate to:

```URL

http://localhost:8080
```
##Implementation Details

The project code is divided into the main package and two additional packages:

    funcs: Contains the server handlers and functions.
    fs: Contains the code that generates the ASCII art.

HTML files are placed in the templates directory, while the banners are in the static directory.
