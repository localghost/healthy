package utils

import "github.com/mitchellh/mapstructure"

func WithDefault(m map[string]interface{}, k string, d interface{}) interface{} {
	if v, ok := m[k]; ok {
		return v
	}
	return d
}

func Decode(input interface{}, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   output,
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		ZeroFields: true,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}
