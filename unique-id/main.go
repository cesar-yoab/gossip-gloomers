package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		// Unmarshall the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		uid := fmt.Sprintf("%s-%s-%s-%d", msg.Src, msg.Dest, body["msg_id"], time.Now().Unix())
		// Update the message type to return back
		body["type"] = "generate_ok"
		body["id"] = uid
		// Echo the original message back with the update message type.
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
