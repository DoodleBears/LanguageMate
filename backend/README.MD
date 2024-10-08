# Language Mate backend

This project is using goframe as the base framework. See [documentation](https://goframe.org/) to learn more about goframe.

## System Requirements
- Operating System:
  - Windows
  - Linux
  - MacOS
- Docker >= 20.10
- Optional (click to see details)
  - [Golang](https://golang.org/doc/install) >= 1.18
  - GoFrame >= 2.7.0
  - [Mariadb](https://mariadb.org/download/)

_You could run the project with docker-compose without setting optional environment._

## Development flows
### Start Up

### Setup environment and config
Copy and rename `manifest\config\config.example` to `manifest\config\config.yaml`.

Copy and rename `.env.example` to `.env`, this file is used to set some environment variables for your local development in docker.

The `nginx` docker compose file is used for prod related development, remove it if you don't need.

### Run project
#### Simple
use command below to build and run the project.
```
docker compose -f docker-compose.yaml -f docker-compose.nginx.yaml up -d --build
```
#### Advanced
1. Revise `.env` file `COMPOSE_FILE` variable for your OS
2. Use `docker compose up -d --build` to build and run the project.

The default exposed ports for different containers are as follows:
- web:8000
- nginx:80
- mariadb:3306

___It is recommended to use WSL2 to run this project on Windows or hot reload is not available. There are some [existing problems](https://github.com/microsoft/WSL/issues/4739) when using WSL.____

### Initialize Database
_We are using [Atlas](https://atlasgo.io/) as our database schema manager._

Database would be auto migrated when starting containers, you could use `atlas` in web container if other operation is needed.
