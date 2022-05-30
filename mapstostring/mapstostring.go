package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var testmap = map[string]map[string]int{}

	testmap["test1"] = map[string]int{"test1": 100}
	testmap["test2"] = map[string]int{"test2": 200}

	jsonStr, err := json.Marshal(testmap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
