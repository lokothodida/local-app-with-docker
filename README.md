## Local "Dockerized" Go Application
This is the companion codebase for the article [*Improve your local web app development with Docker*](https://lawrence.okothodida.com/feed).

It is an HTTP server that simply says *"Hello, world!"*, with a counter for
how many hits it has received. These hits are counted in [Redis](https://redis.io/) and the server
is written in [Go](https://golang.org/).

## Installation
* Install [Docker](https://docs.docker.com/get-docker/)
* Install [Docker Compose](https://docs.docker.com/compose/install/)
* Clone this repository

## Commands
* Build the container (with your user permissions)
    ```shell script
    $ make container
    ```
* Start the application and infrastructure
    ```shell script
    $ make start
    ```
* Run the web server in the container (served on port `HTTP_PORT` in your [`.env`](.env.dist) file)
    ```shell script
    $ make web
    ```
* Or run the web server from the host machine:
    ```shell script
    $ source .env && go run main.go
    ```
* Enter the `go` container
    ```shell script
    $ make sh
    ```
* Stop the application and infrastructure
    ```shell script
    $ make stop
    ```

## TODO
- [ ] Add link to final article
- [x] Add fix for permission issues when editing files created by the container
