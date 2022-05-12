# How to start
1. Spuštění přes composer -> docker-compose up 
2. Spuštění standalone container -> docker run -d -p 8080:8080 jirituryna/incicounter:latest
3. Pokud udělám změnu -> docker build -t jirituryna/incicounter .
-> vygeneruje na locale nový docker image
4. Push do Docker Hubu -> docker hub push jirituryna/incicounter:latest

# API Commands
## Incident Counter
- docker image - 
jirituryna/incicounter:latest


## Get all counters
- vrací aktuální seznam counterů

curl localhost:8080/counters

## Add a counter
- bere na vstupu .json file -> v repozitáři předem vytvořen .json file "data1.json" (při spouštění commandu je nutno být ve složce s /data1.json)

curl localhost:8080/counters --include --header "Content-Type: application/json" -d @data1.json --request "POST"

## INCREMENT / DECREMENT a counter
- incrementuje counter s definovaným id (localhost:8080/counters/1)

- curl localhost:8080/counters/inc/0 --request "PATCH"
- curl localhost:8080/counters/dec/0 --request "PATCH"


## Delete a counter
- smaže counter se specifikovaným id

curl localhost:8080/counters/0 --request "DELETE"

## Delete all counters
- smaže všechny countery

curl localhost:8080/counters/del/all --request "DELETE"

## Reset Counter
- smaže counter podle specifikovanýho id

curl localhost:8080/counters/res/0 --request "PATCH"

## SET Counter
- Nastaví counter se spec. id na konkrétní hodnotu

curl localhost:8080/counters/set/0/20 --request "PATCH"

# Docker Knowhow: Images
## Create Image se zvoleným image name
- spouští dockerfile a ukládá na localhost docker image

docker build -t jirituryna/incicounter .

## Run docker image s published portem 8080:8080 
- v main appce router.Run(":8080"), v Dockerfile EXPOSE 8080

docker run -d -p 8080:8080 jirituryna/incicounter:latest

## Přehled všech uložených imagů na localu
docker image ls

## Přehled všech currently running containerů
docker ps

## Access container 
- /bin/bash (Dockerfile has to contain RUN apk update && apk add bash)

docker exec -it <containerid> /bin/bash

## Přidání packagů v containeru
apk add ...
