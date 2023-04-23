# Gossip Gloomers
Solutions to [Gossip Gloomer](https://fly.io/dist-sys/) challenges a set of distributed systems challenges.

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
  "body": <json_body>
}
```

## Echo Challenge
More of a "getting started" guide" to get the hang of working with Maelstrom in Go. 

## Unique ID Generator
In this challenge we need to implement a globally-unique ID generation system that runs against Maelstrom's unique-ids workload. The service should be totally available, meaning that it can continue to operate even in the face of network partitions.

## Broadcast
On this challenge, we need to implement a broadcast system that gossips messages between all nodes in the cluster. Gossiping is a common way to propagate information across a cluster when you don't need strong consistency guarantees.

This challenge is broken up in multiple sections so that you can build out your system incrementally.

### Single-Node Broadcast
First, we start out with a single-node broadcast system. That may sound like an oxymoron but this lets us get our message handlers working correctly in isolation before trying to share messages between nodes.

### Multi-Node Broadcast
In this challenge, we build on our Single-Node Broadcast implementation and replicate our messages across a cluster that has no network partitions.

The node should propagate values it sees from broadcast messages to the other nodes in the cluster. It can use the topology passed to your node in the topology message or you can build your own topology.

The simplest approach is to simply send a node's entire data set on every message, however, this is not practical in a real-world system. Instead, try to send data more efficiently as if you were building a real broadcast system.

Values should propagate to all other nodes within a few seconds.

### Fault Tolerant Broadcast
In this challenge, we build on our Multi-Node Broadcast implementation, however, this time we'll introduce network partitions between nodes so they will not be able to communicate for periods of time.

The node should propagate values it sees from broadcast messages to the other nodes in the cluster—even in the face of network partitions! Values should propagate to all other nodes by the end of the test. Nodes should only return copies of their own local values.


### Efficient Broadcast Pt. 1

### Efficient Broadcast Pt. 2