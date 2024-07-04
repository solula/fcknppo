package message

import "encoding/json"

type Message struct {
	Payload json.RawMessage
	Headers map[string]interface{}
}
