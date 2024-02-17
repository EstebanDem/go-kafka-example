# Microservices with Kafka as Broker - Example

## Architecture diagram

![Microservices with kafka - golang](https://github.com/EstebanDem/go-kafka-example/assets/83315171/55403910-db9d-44ba-8be6-7a4f8f5263c3)

- `registration-service`: handles the client request and register users in the DB, it produces to `new-user` topic.it listens to events that modify user's status by http
- `broker-service`: listens to messages in every topic and redirects them to interested apps
- `email-service`: validate email request from broker and, if valid, produces to `email-validator` topic

## Run Locally

Initiate Kafka with Zookeeper

*I used kafka `2.11-0.11.0.1` but it should work any other higher versions,
you can download it from https://kafka.apache.org/downloads*

Run Zookeeper

```bash
  bin/zookeeper-server-start.sh config/zookeeper.properties
```

Run Kafka

```bash
  bin/kafka-server-start.sh config/server.properties
```

Create topics

```bash
  # New user topic
  bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic new-user
  # Email Validation topic
  bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic email-validation
```

Clone the project

```bash
  git clone https://github.com/EstebanDem/go-kafka-example
```

Go to the project directory

```bash
  cd go-kafka-example
```

Run every microservice

```bash
  cd broker-service
  go run cmd/main.go
```

```bash
  cd registration-service
  go run cmd/main.go
```

```bash
  cd email-service
  go run cmd/main.go
```

## Demo

https://github.com/EstebanDem/go-kafka-example/assets/83315171/63fdbae8-9897-4cb9-b049-2d207aa8b0de


