## Local "Dockerized" Go Application
This is the companion codebase for the article [Improve your local web app development with Docker](https://lawrence.okothodida.com/feed).

It is an HTTP server that simply says *"Hello, world!"*, with a counter for
how many hits it has received. These hits are counted in Redis and the server
is written in Go.

## Installation
* Install [Docker](https://docs.docker.com/get-docker/)
* Install [Docker-Compose](https://docs.docker.com/compose/install/)
* Clone this repository
* Create a `.env` file from the example:
    ```shell script
    cp .env.dist .env
    ```

## Commands
* Start the application and infrastructure
    ```shell script
    $ docker-compose up
    ```
* Run the web server (served on port `HTTP_PORT` in your [`.env`](.env.dist) file)
    ```shell script
    $ docker-compose exec app go run main.go
    ```
* Enter the container
    ```shell script
    $ docker-compose exec app sh
    ```
* Stop the application and infrastructure
    ```shell script
    $ docker-compose down
    ```

## TODO
- [ ] Add link to final article
- [ ] Add fix for permission issues when editing files created by the container
