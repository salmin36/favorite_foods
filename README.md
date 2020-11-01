# Favorite foods
This is a dockerised system for keeping track of your favorite foods

It is designed to have 3 docker images to run the entire application in with docker-compose.

These 3 components are :
- Redis database
- Frontend made with NodeJS + React
- Backend made wiht Golang 

## Prerequisites
One has to have installed
- docker 
- docker-compose 

## Running the app
```
    docker-compose up

```

Access the website via http://127.0.0.1:8080

## Usefull 

In order to manaually go to redis container to check what exists.
First run all the services up with 
```
docker-compose up
```
Then find what is the redis container id:
```
docker ps |grep redis | awk '{print $1}'
```
Then open shell inside redis container (replace <redis_id> with the id from last command):
```
docker exec -it <redis_id> /bin/sh
```
When inside container just run the redis cli:
```
redis-cli
```
When inside redis cli then authenticate via (replace <password>  with real redis password given in docker-compose):
```
 AUTH <password>
```
