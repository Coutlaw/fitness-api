# Fitness API
This is an API I built to be the backend to a personal trainers fitness app, the API has been deprecated in favor of one in Typescript and I chose to make it Public to show off some cool features I developed. 

# About
This is Rest API with that has Jwt authentication for users, and can assign workout programs to users. The Database is Postgres and I have the schema in a private repository, there are tables for Users, Workout Programs and Auth. Trainers can create a workout program that can be assigned to multiple users. The users can view their workouts per day, make unique changes to that workout, or make comments. This API was consumed by a React front end.

# Docs
Some API docs are hosted in github pages, but they need some love [here](https://coutlaw.github.io/fitness-api-docs/)
For details check out the wiki [here](https://github.com/Coutlaw/fitness-api/wiki)

# Hosting
This project is set up with GitHub Actions to automatically deploy to Microsoft Azure as a new docker image any time a new release version is made in Github.
The database was also hosted in Azure.

## Build & Run (local) outside of docker

Run the DB in docker: `docker-compose up` (you will need to initialize the db with the tables in the api's sql repo)

Build: `go build -o build`

Run the API: `build/fitness-api`

Access pgAdmin: `http://localhost:5050/browser/#`

clean your db volumes: `docker volume rm fitness-api_database_postgres`

All other details for running the API in docker can be found in the docs
