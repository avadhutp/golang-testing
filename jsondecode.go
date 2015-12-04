package main

import (
	"encoding/json"
	"fmt"
)

func decodeJSON(raw string) (ret map[string]string) {
	err := json.Unmarshal([]byte(raw), &ret)

	if err != nil {
		return
	}

	return
}

func main() {
	input := "{\"obi\":\"wan\"}"
	fmt.Println("Input: " + input)
	fmt.Println("Ouput: ")
	fmt.Println(decodeJSON(input))
}
