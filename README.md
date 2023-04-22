# Gossip Gloomers
Solutions to Gossip Gloomer challenges a set of distributed systems challenges.

1. Echo Challenge
2. Unique ID Generator
3. Broadcast
4. Grow-Only Counter Challenge
5. Kafka-Style Logging Challenge
6. Totally-Available Transactions


## Message structure
All messages are sent in json format with the following structure:

```json
{
  "src": "c1",
  "dest": "n1",
  "body": {
    "type": "echo",
    "msg_id": 1,
    "echo": "Please echo 35"
  }
}
```
