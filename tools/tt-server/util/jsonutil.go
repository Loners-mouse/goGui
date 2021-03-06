package util

import (
	"encoding/json"
	"fmt"
)

func Json2Map(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
	
	return m, nil
}

func Map2Json(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}
	
	return string(jsonByte), nil
}