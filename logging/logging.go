package logging

import (
	"encoding/json"
	"fmt"
)

func StructToPrettyString(structure interface{}) (*string, error) {
	b, err := json.MarshalIndent(structure, "", "  ")
	if err == nil {
		str := string(b)
		return &str, nil
	} else {
		return nil, err
	}
}

func StructToString(structure interface{}) (*string, error) {
	b, err := json.Marshal(structure)
	if err == nil {
		str := string(b)
		return &str, nil
	} else {
		return nil, err
	}
}

func PrintStruct(prefix string, structure interface{}) {
	fmt.Printf("%s: %#v\n", prefix, structure)
}
