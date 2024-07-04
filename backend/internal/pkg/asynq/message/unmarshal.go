package message

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	ErrInvalidUnmarshal = "ошибка распаковки данных"
)

// ExtractHeaders извлекает хэдеры из буфера и возвращает мапу из них
func ExtractHeaders(buf []byte) (map[string]interface{}, error) {
	// Распаковываем буфер в объект Message
	msg, err := unmarshalMessage(buf)
	if err != nil {
		return nil, err
	}

	return msg.Headers, nil
}

// Unmarshal читает payload сообщения в объект v
func Unmarshal(buf []byte, v interface{}) error {
	// В любом случае распаковываем буфер в объект Message
	msg, err := unmarshalMessage(buf)
	if err != nil {
		return err
	}

	// Если v не указан, то ничего не делаем
	if v == nil {
		return nil
	}

	err = unmarshalPayload(msg.Payload, v)
	if err != nil {
		return err
	}

	return nil
}

func unmarshalMessage(buf []byte) (Message, error) {
	msg := Message{}
	err := json.Unmarshal(buf, &msg)
	if err != nil {
		return msg, fmt.Errorf("%s: %w", ErrInvalidUnmarshal, err)
	}

	return msg, nil
}

func unmarshalPayload(buf json.RawMessage, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("%s: передан некорректный тип объекта <%s>", ErrInvalidUnmarshal, reflect.TypeOf(v))
	}
	rv = rv.Elem()

	//switch v.(type) {
	//case *int, *int8, *int16, *int32, *int64:
	//	val, err := strconv.ParseInt(string(buf), 10, 64)
	//
	//	if err != nil {
	//		return fmt.Errorf("ошибка конвертации значения int: %w", err)
	//	}
	//	rv.SetInt(val)
	//case *uint, *uint8, *uint16, *uint32, *uint64:
	//	val, err := strconv.ParseUint(string(buf), 10, 64)
	//
	//	if err != nil {
	//		return fmt.Errorf("ошибка конвертации значения uint: %w", err)
	//	}
	//	rv.SetUint(val)
	//case *[]byte:
	//	rv.SetBytes(buf)
	//case *string:
	//	rv.SetString(string(buf))
	//default:
	err := json.Unmarshal(buf, v)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrInvalidUnmarshal, err)
	}
	//}

	return nil
}
