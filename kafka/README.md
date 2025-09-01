# Kafka Labs

## Run Kafka Cluster

```bash
docker-compose up -d
```

## Stop Kafka Cluster

```bash
docker-compose down
```

## Access Kafka Broker

```bash
docker exec --workdir /opt/kafka/bin/ -it kafka-broker-1 sh
```

## Create Kafka Topic

```bash
./kafka-topics.sh \
  --create \
  --topic products.prices.changelog \
  --partitions 1 \
  --replication-factor 1 \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

## Produce Messages to Topic

```bash
echo "coffee pads 10" | ./kafka-console-producer.sh \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

## Consume Messages from Topic

Press Ctrl-C to cancel.

```bash
./kafka-console-consumer.sh \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

```bash
Processed a total of 0 messages
```

## Consume Messages from Topic (from beginning)

Press Ctrl-C to cancel.

```bash
./kafka-console-consumer.sh \
  --topic products.prices.changelog \
  --from-beginning \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

```bash
coffee pads 10
Processed a total of 1 messages
```

## Run Kafka Producer

Press Ctrl-C to cancel.

```bash
./kafka-console-producer.sh \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

## Run Kafka Consumer

Press Ctrl-C to cancel.

```bash
./kafka-console-consumer.sh \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

## Describe Kafka Topic

```bash
./kafka-topics.sh \
  --describe \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

```bash
Topic: products.prices.changelog        TopicId: 4sGDOBjwTauvZ-J8usBWWw PartitionCount: 1     ReplicationFactor: 1     Configs:
  Topic: products.prices.changelog        Partition: 0    Leader: 5       Replicas: 5   Isr: 5   Elr:    LastKnownElr:
/opt/kafka/bin
```
