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

Output:

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

Output:

```bash
Topic: products.prices.changelog        TopicId: 4sGDOBjwTauvZ-J8usBWWw PartitionCount: 1     ReplicationFactor: 1     Configs:
  Topic: products.prices.changelog        Partition: 0    Leader: 5       Replicas: 5   Isr: 5   Elr:    LastKnownElr:
```

Brokers 1-3 are controllers only and do not host any partitions. Broker 4-6 are
data brokers.

| Broker 1 | Broker 2 | Broker 3 | Broker 4 | Broker 5 | Broker 6 |
| -------- | -------- | -------- | -------- | -------- | -------- |
|          |          |          |          | P0 (L)   |          |

`Topic: products.prices.changelog`: This is the name of the Kafka topic being
described.

`TopicId: 4sGDOBjwTauvZ-J8usBWWw`: This is a unique identifier assigned to the
topic by Kafka that ensures that even if the topic name is changed, the topic
can still be uniquely identified.

`PartitionCount: 1`: This indicates that the topic is divided into 1 partition.

`ReplicationFactor: 1`: This means that each partition of the topic has only one
replica.A replication factor of 1 provides no fault tolerance. If the broker
hosting this single replica fails, the data in this topic becomes unavailable.
For production environments, a replication factor of 3 is a common
recommendation to ensure high availability.

`Configs:`: This section would list any custom configuration settings that have
been applied to this specific topic, such as `retention.ms` or `cleanup.policy`.
In this case, no custom configurations are listed.

`Partition: 0`: This line provides details for a specific partition within the
topic. Since `PartitionCount` is 1, there is only one partition, numbered 0.

`Leader: 5`: This is the broker ID of the Kafka broker that is currently the
leader for this partition.

`Replicas: 5`: This is a list of all brokers that hold a replica of this
partition. In this case, only broker 5 holds a replica, which is consistent with
the `ReplicationFactor` of 1.

`Isr: 5`: This stands for In-Sync Replicas. It's a list of the brokers that have
a replica of this partition and are fully caught up with the leader. The leader
is always part of the ISR. Since the replication factor is 1, only the leader
(broker 5) is in the ISR.

`Elr:`: This stands for Eligible Leader Replicas and contains the list of
brokers that are also ISR eligible to become the leader for that partition if
the current leader fails.

`LastKnownElr:`: Contains the list of previously eligible replicas that had an
unclean shutdown.

## List Kafka Topics

```bash
./kafka-topics.sh \
  --list \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Output:

```bash
__consumer_offsets
products.prices.changelog
```

## Delete Kafka Topic

```bash
./kafka-topics.sh \
  --delete \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

## Create, Customize, and Delete Topics

Create a topic:

```bash
./kafka-topics.sh \
  --create \
  --topic products.prices.changelog \
  --replication-factor 2 \
  --partitions 2 \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Describe the topic:

```bash
./kafka-topics.sh \
  --describe \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Output:

```bash
Topic: products.prices.changelog        TopicId: 8zaM8ODHS6Gy_7LIhcATow PartitionCount: 2       ReplicationFactor: 2    Configs:
        Topic: products.prices.changelog        Partition: 0    Leader: 4       Replicas: 4,5   Isr: 4,5        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 1    Leader: 5       Replicas: 5,6   Isr: 5,6        Elr:    LastKnownElr:
```

Brokers 1-3 are controllers only and do not host any partitions. Broker 4-6 are
data brokers.

| Broker 1 | Broker 2 | Broker 3 | Broker 4 | Broker 5 | Broker 6 |
| -------- | -------- | -------- | -------- | -------- | -------- |
|          |          |          | P0 (L)   | P0 (R)   |          |
|          |          |          |          | P1 (L)   | P1 (R)   |

Increase partitions to 3:

```bash
./kafka-topics.sh \
  --alter \
  --topic products.prices.changelog \
  --partitions 3 \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Describe the topic:

```bash
./kafka-topics.sh \
  --describe \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Output:

```bash
Topic: products.prices.changelog        TopicId: 8zaM8ODHS6Gy_7LIhcATow PartitionCount: 3       ReplicationFactor: 2    Configs:
        Topic: products.prices.changelog        Partition: 0    Leader: 4       Replicas: 4,5   Isr: 4,5        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 1    Leader: 5       Replicas: 5,6   Isr: 5,6        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 2    Leader: 5       Replicas: 5,6   Isr: 5,6        Elr:    LastKnownElr:
```

Brokers 1-3 are controllers only and do not host any partitions. Broker 4-6 are
data brokers.

| Broker 1 | Broker 2 | Broker 3 | Broker 4 | Broker 5 | Broker 6 |
| -------- | -------- | -------- | -------- | -------- | -------- |
|          |          |          | P0 (L)   | P0 (R)   |          |
|          |          |          |          | P1 (L)   | P1 (R)   |
|          |          |          |          | P2 (L)   | P2 (R)   |

Increase partitions to 5:

```bash
./kafka-topics.sh \
  --alter \
  --topic products.prices.changelog \
  --partitions 5 \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Describe the topic:

```bash
./kafka-topics.sh \
  --describe \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

Output:

```bash
Topic: products.prices.changelog        TopicId: 8zaM8ODHS6Gy_7LIhcATow PartitionCount: 5       ReplicationFactor: 2    Configs:
        Topic: products.prices.changelog        Partition: 0    Leader: 4       Replicas: 4,5   Isr: 4,5        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 1    Leader: 5       Replicas: 5,6   Isr: 5,6        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 2    Leader: 5       Replicas: 5,6   Isr: 5,6        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 3    Leader: 6       Replicas: 6,4   Isr: 6,4        Elr:    LastKnownElr:
        Topic: products.prices.changelog        Partition: 4    Leader: 4       Replicas: 4,5   Isr: 4,5        Elr:    LastKnownElr:
```

Brokers 1-3 are controllers only and do not host any partitions. Broker 4-6 are
data brokers.

| Broker 1 | Broker 2 | Broker 3 | Broker 4 | Broker 5 | Broker 6 |
| -------- | -------- | -------- | -------- | -------- | -------- |
|          |          |          | P0 (L)   | P0 (R)   |          |
|          |          |          |          | P1 (L)   | P1 (R)   |
|          |          |          |          | P2 (L)   | P2 (R)   |
|          |          |          | P3(R)    |          | P3 (L)   |
|          |          |          | P4 (L)   | P4 (R)   |          |

Reducing the number of partitions is not possible because it could lead to data
loss and reordering of messages. Therefore, Kafka does not allow decreasing the
number of partitions for an existing topic.

Delete the topic:

```bash
./kafka-topics.sh \
  --delete \
  --topic products.prices.changelog \
  --bootstrap-server kafka-broker-1:19092,kafka-broker-2:19092,kafka-broker-3:19092
```

## Messages

Max message size is 1 MB by default. To produce messages larger than 1 MB, you
need to increase the `message.max.bytes` broker configuration and the
`max.request.size` producer configuration. However, for optimal performance, it
is recommended to keep messages smaller than 1 MB.

Note: Kafka records always contain a timestamp.

### Message Types

#### States

State messages represent the current status or condition of an entity at a
specific point in time. They provide a snapshot of the entity's attributes,
allowing systems to understand its present situation. State messages are often
used in scenarios where the latest information is crucial, such as monitoring
systems, user profiles, or inventory levels.

```txt
{"user_id": "12345", "status": "active", "last_login": "2024-10-01T12:34:56Z", "account_balance": 250.75}
{"user_id": "67890", "status": "inactive", "last_login": "2024-09-15T08:22:10Z", "account_balance": 0.00}
{"user_id": "54321", "status": "active", "last_login": "2024-10-02T14:20:30Z", "account_balance": 100.00}
```

#### Deltas

Delta messages capture changes or updates to an entity's state, focusing on what
has changed rather than the entire state. They are particularly useful in
scenarios where bandwidth or storage is a concern, as they minimize the amount
of data transmitted by only including the differences. Delta messages are
commonly used in synchronization processes, event sourcing, and real-time
updates.

```txt
{"user_id": "12345", "account_balance": 300.00}
{"user_id": "67890", "account_balance": 50.00}
{"user_id": "54321", "last_login": "2024-10-03T09:15:00Z"}
{"user_id": "12345", "status": "inactive"}
```

#### Events

Event messages represent significant occurrences or actions that have taken
place within a system. They are typically used to signal that something
noteworthy has happened, such as a user action, system change, or external
trigger. Event messages often include a timestamp and relevant metadata,
allowing systems to react to these occurrences in real-time or through
event-driven architectures.

```txt
{"event_type": "user_signup", "user_id": "12345", "payload": {"referral_code": "XYZ123"}}
{"event_type": "purchase", "user_id": "67890", "payload": {"item_id": "A1B2C3", "amount": 49.99}}
{"event_type": "password_reset", "user_id": "54321", "payload": {"reset_method": "email"}}
```

#### Commands

Command messages are directives sent to a system or component, instructing it to
perform a specific action or operation. They are typically used in scenarios
where a system needs to be controlled or manipulated, such as in remote
procedure calls, task scheduling, or workflow management. Command messages often
include payload data that provides additional context or parameters for the
action to be executed.

```txt
{"command": "create_user", "payload": {"username": "new_user", "email": "example@mail.com"}}
{"command": "update_balance", "payload": {"user_id": "12345", "amount": 100.00}}
{"command": "deactivate_account", "payload": {"user_id": "67890"}}
```

### Message Structure

Kafka Record:

- Key: Optional, used for partitioning and message identification (Byte array).
  Keys serve as identifiers for messages, allowing for efficient partitioning
  and retrieval. They can be used to group related messages together, ensuring
  that messages with the same key are sent to the same partition. This is
  particularly useful for maintaining order and consistency in message
  processing.
- Value: The actual message payload (Byte array).
- Headers: Optional, key-value pairs for additional metadata (Array of key-value
  pairs). It is used for technical metadata, not business data.
- Timestamp: The time the record was created or received by the broker (Long).
