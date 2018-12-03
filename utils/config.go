package utils

import "github.com/mitchellh/mapstructure"

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
