# Product catalog Manager

## Description
This is a study project to learn event driven architecture and microservices. The project is a simple product catalog manager.
The project is divided in 3 parts:
- Product catalog crud (mongodb)
- Message Brooker (kafka, rabbitmq and SQS)
- AWS integration for JSON storage (S3)

## Initial setup
To run the project you need to have docker and docker-compose installed. Infra is started in daemon mode.
```shell
docker-compose -f docker/infra.yml up -d
```
and after initialization (a few seconds after, or minutes if it's the first time) run:
```shell
docker-compose -f docker/app.yml up
```
Application docker is separated from infra to allow reload without restarting the entire infrastructure services.

## Kafka
Kafka Management (control center):
```shell
http://localhost:9021
```

## RabbitMQ
RabbitMQ Management:
```shell
http://localhost:15672

username: guest
password: guest
```

