package main

import (
	"encoding/json"
	"fmt"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Topology struct {
	Type     string              `json:"type,omitempty`
	Topology map[string][]string `json:"topology,omitempty`
}

func PushMessage(n *maelstrom.Node, initNode string, lastNode string, node any, message any) error {
	dest, ok := node.(string)
	if !ok {
		return fmt.Errorf("message is not a string")
	}
	body := map[string]any{"init_node": initNode, "message": message, "type": "push", "last_node": lastNode}
	return n.Send(dest, body)
}

func InStore(store *[]any, value any) bool {
	for _, v := range *store {
		if v == value {
			return true
		}
	}

	return false
}

func main() {
	n := maelstrom.NewNode()
	store := make([]any, 0)
	var topology Topology

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		// Unmarshall the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// save message to be read later
		message := body["message"]
		store = append(store, message)
		delete(body, "message")

		// update message type to ok
		body["type"] = "broadcast_ok"
		initNode := msg.Dest

		if topology.Topology != nil {
			for _, node := range topology.Topology[initNode] {
				if err := PushMessage(n, initNode, initNode, node, message); err != nil {
					return err
				}
			}
		}

		// Echo the original message back with the update message type.
		return n.Reply(msg, body)
	})

	n.Handle("push", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		message := body["message"]
		if inStore := InStore(&store, message); !inStore {
			store = append(store, message)
		}

		initNode, ok := body["init_node"].(string)
		if !ok {
			return fmt.Errorf("init_node is not a string")
		}
		lastNode, ok := body["last_node"].(string)
		if !ok {
			return fmt.Errorf("last_node is not a string")
		}

		self := msg.Dest
		if topology.Topology != nil {
			for _, node := range topology.Topology[self] {
				if node == initNode || node == lastNode {
					continue
				}

				if err := PushMessage(n, initNode, self, node, message); err != nil {
					return err
				}
			}
		}

		return nil
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		// Unmarshall the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back

		body["type"] = "read_ok"
		body["messages"] = store

		// Echo the original message back with the update message type.
		return n.Reply(msg, body)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		// Unmarshall the topology map and return.
		if err := json.Unmarshal(msg.Body, &topology); err != nil {
			return err
		}

		// Update the message type to return back
		body := map[string]string{"type": "topology_ok"}
		delete(body, "topology")
		// Echo the original message back with the update message type.
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
