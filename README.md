# ASCII-ART-WEB 

## Description 
ascii-art-web is a web application consists in creating and running a server, in which it will be possible to user a web GUI version of terminal base ascii-art

## Authors
    - Oussama Benali
    - Ziad ...
    - Hamza ...

## Usage
    to start the server you need to run the bash command 
    ```bash
        go run main.go
    ```
    and then fire the browser on http://localhost on port 8080
    ```URL 
        http://localhost:8080

## Implementation details
    the projects code is devided into the main package and two other packages, 
    the package funcs that has the server handlers and function
    and the package fs where the code that generates the ascii art is 
    html files are placed in the templates derectory while the banners are in the static derectory 
    
