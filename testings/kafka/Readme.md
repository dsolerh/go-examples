# How to run a Kafka broker locally with Docker

## Set up

To run Kafka in Docker, you'll need to install Docker Desktop. Then, after you've installed Docker Desktop, you'll use Docker Compose to create your Kafka container. Docker Compose uses a YAML configuration file to manage your Docker components (services, volumes, networks, etc.) in an easy-to-maintain approach. Docker Desktop includes Docker Compose, so there are no additional steps you need to take. Finally, you'll need an account on Docker Hub so that Docker Desktop can pull the images specified in the Docker Compose file.

## The Docker Compose file

Before you get started, here's the Docker Compose file for the tutorial:

```yml
services:
  broker:
    image: apache/kafka:latest
    hostname: broker
    container_name: broker
    ports:
      - 9092:9092
    environment:
      KAFKA_BROKER_ID: 1
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT,CONTROLLER:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://broker:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
      KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
      KAFKA_PROCESS_ROLES: broker,controller
      KAFKA_NODE_ID: 1
      KAFKA_CONTROLLER_QUORUM_VOTERS: 1@broker:29093
      KAFKA_LISTENERS: PLAINTEXT://broker:29092,CONTROLLER://broker:29093,PLAINTEXT_HOST://0.0.0.0:9092
      KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT
      KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER
      KAFKA_LOG_DIRS: /tmp/kraft-combined-logs
      CLUSTER_ID: MkU3OEVBNTcwNTJENDM2Qk
```

Let's review some of the key parts of the YAML


```yml
broker:
    image: apache/kafka:latest
    hostname: broker
    container_name: broker
    ports:
      - 9092:9092
```

- `image` contains the organization, image name and the version to use. In this case it will always pull the latest one available.
- `ports` specifies the ports you use to connect to Kafka.

```yml
KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1
KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1
```

These settings affect topic replication and min-in-sync replicas and should only ever use these values when running a single Kafka broker on Docker. For a full explanation of these and other configurations, consult the Kafka documentation.

```yml
KAFKA_PROCESS_ROLES: broker,controller
```

Since the introduction of KRaft Kafka no longer requires Apache ZooKeeper® for managing cluster metadata, using Kafka itself instead. One advantage of the new KRaft mode is that you can have a single Kafka broker to handle both metadata and client requests in small, local development environment. The docker-compose.yml file for this tutorial uses this approach, leading to faster startup times and simpler configuration. Note that, in a production setting, you'll have distinct Kafka brokers for handling requests and operating as a cluster controller.

## Running Kafka on Docker

To run Kafka on Docker, first confirm your Docker Desktop is running (`docker ps`). Then execute the following command from this directory:

```bash
docker compose up -d
```

The `-d` flag runs the docker container in detached mode which is similar to running Unix commands in the background by appending `&`. To confirm the container is running, run this command:

```bash
docker logs broker
```

And if everything is running ok you'll see something like this at this at the end of screen output:

```
[2024-12-10 08:43:57,555] INFO Kafka version: 3.8.1 (org.apache.kafka.common.utils.AppInfoParser)
[2024-12-10 08:43:57,555] INFO Kafka commitId: 70d6ff42debf7e17 (org.apache.kafka.common.utils.AppInfoParser)
[2024-12-10 08:43:57,555] INFO Kafka startTimeMs: 1733820237555 (org.apache.kafka.common.utils.AppInfoParser)
[2024-12-10 08:43:57,556] INFO [KafkaRaftServer nodeId=1] Kafka Server started (kafka.server.KafkaRaftServer)
[2024-12-10 08:53:57,422] INFO [NodeToControllerChannelManager id=1 name=registration] Node 1 disconnected. (org.apache.kafka.clients.NetworkClient)
```

Now let's produce and consume a message! To produce a message, let's open a command terminal on the Kafka container:

```bash
docker exec -it -w /opt/kafka/bin broker sh
```

Then create a topic:

```bash
./kafka-topics.sh --create --topic my-topic --bootstrap-server broker:29092
```

The result of this command should be

```
Created topic my-topic.
```

### Important

Take note of the --bootstrap-server flag. Because you're connecting to Kafka inside the container, you use broker:29092 for the host:port. If you were to use a client outside the container to connect to Kafka, a producer application running on your laptop for example, you'd use localhost:9092 instead.

Next, start a console producer with this command:

```bash
./kafka-console-producer.sh  --topic my-topic --bootstrap-server broker:29092
```

At the prompt copy each line one at time and paste into the terminal hitting enter key after each one:

```
All streams
lead to Kafka
```

Then enter a CTRL-C to close the producer.
Now let's consume the messages with this command:

```bash
./kafka-console-consumer.sh --topic my-topic --from-beginning --bootstrap-server broker:29092
```

And you should see the following:

```
All streams
lead to Kafka
```

Enter a CTRL-C to close the consumer and then type exit to close the docker shell.

To shut down the container, run:

```bash
docker compose down -v
```
