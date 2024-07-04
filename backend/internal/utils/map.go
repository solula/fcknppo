package utils

import "github.com/mitchellh/mapstructure"

func MapToStruct[T any](m map[string]interface{}) (*T, error) {
	var res T
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &res,
		TagName:  "json",
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(m)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
