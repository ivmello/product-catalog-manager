# Product catalog Manager

## Description
This is a study project to learn event driven architecture and microservices. The project is a simple product catalog manager.
The project is divided in 3 parts:
- Product catalog crud (mongodb)
- Message Brooker (kafka, rabbitmq and SQS)
- AWS integration for JSON storage (S3)

## Initial setup
To run the project you need to have docker and docker-compose installed.
```shell
docker-compose -f docker-compose.infra.yml up
```
and after initialization run:
To run the project you need to have docker and docker-compose installed.
```shell
docker-compose -f docker-compose.app.yml up
```
Application docker is separated from infra to allow reload without restarting the entire infra.

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

