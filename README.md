# opensky-rabbit
RabbitMQ Tutorial extended to pull sample data from an OpenSky Network API

## Building
Run ```go build```

## Prerequisites
Create an OpenSkyNetwork account, and save your username and password under the following env vars:
Username: OPENSKY_USERNAME
Password: OPENSKY_PASSWORD

## Running 
Start rabbitmq docker container: ```docker compose up```

To run in consumer mode run ```./opensky-rabbit consumer```
To run in producer mode run ```./opensky-rabbit producer```
