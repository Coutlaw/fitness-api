# fitness-api
Its an API to track fitness stuff

For details and docs, check out the wiki [here](https://github.com/Coutlaw/fitness-api/wiki)


## Build & Run (local) outside of docker

Run the DB in docker: `docker-compose up` (you will need to initialize the db with the tables in the api's sql repo)

Build: `go build -o build`

Run the API: `build/fitness-api`

Access pgAdmin: `http://localhost:5050/browser/#`

clean your db volumes: `docker volume rm fitness-api_database_postgres`

All other details for running the API in docker can be found in the docs
