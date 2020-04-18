# fitness-api
Its an API to track fitness stuff

For details and docs, check out the wiki [here](https://github.com/Coutlaw/fitness-api/wiki)


## Build & Run (local)

Run the DB in docker: `docker-compose up`
Build: `go build -o build`
Run the API: `build/fitness-api`
Access pgAdmin: `http://localhost:5050/browser/#`
clean your db volumes: `docker volume rm fitness-api_database_postgres`

All other details can be found in the docs
