package helper

import (
	"encoding/json"
	"os"
)

func MarshaUp(obj interface{}) []byte {
	data, err := json.Marshal(obj)
	if err != nil {
		return []byte{}
	}
	return data
}

func UnmarshaUp(jsonString string, result interface{}) {
	json.Unmarshal([]byte(jsonString), &result)
}

func SetLocalVar(key string, localVar *string, defaultValue string) {
	if temp := os.Getenv(key); temp != "" {
		*localVar = temp
	} else {
		*localVar = defaultValue
	}
}
