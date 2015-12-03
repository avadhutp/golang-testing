package main

import (
	"encoding/json"
)

func decodeJSON(raw string) (ret map[string]string) {
	err := json.Unmarshal([]byte(raw), &ret)

	if err != nil {
		return
	}

	return
}
