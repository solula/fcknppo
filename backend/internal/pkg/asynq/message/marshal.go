package message

import (
	"encoding/json"
	"fmt"
)

const (
	ErrInvalidMarshal = "ошибка упаковки данных"
)

// Marshal маршалит payload и headers в буфер
func Marshal(payload interface{}, headers map[string]interface{}) ([]byte, error) {
	var payloadBytes []byte

	// Маршалим payload только если он указан
	if payload != nil {
		var err error
		payloadBytes, err = marshalPayload(payload)
		if err != nil {
			return nil, err
		}
	}

	msg := Message{
		Payload: payloadBytes,
		Headers: headers,
	}

	msgBytes, err := marshalMessage(msg)
	if err != nil {
		return nil, err
	}

	return msgBytes, nil
}

func marshalMessage(msg Message) ([]byte, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrInvalidMarshal, err)
	}

	return bytes, nil
}

func marshalPayload(payload interface{}) (json.RawMessage, error) {
	var dataBytes []byte

	//switch typedData := payload.(type) {
	//case int, int8, int16, int32, int64:
	//	intData := reflect.ValueOf(payload).Int()
	//	dataBytes = []byte(strconv.FormatInt(intData, 10))
	//case uint, uint8, uint16, uint32, uint64:
	//	uintData := reflect.ValueOf(payload).Uint()
	//	dataBytes = []byte(strconv.FormatUint(uintData, 10))
	//case []byte:
	//	dataBytes = typedData
	//case string:
	//	dataBytes = []byte(typedData)
	//default:
	var err error
	dataBytes, err = json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	//}

	return dataBytes, nil
}
