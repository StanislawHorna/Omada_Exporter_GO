package utils

import (
	"encoding/json"
)

func MapToStruct(data map[string]any, out any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonBytes, out)
}
