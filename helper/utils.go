package helper

import (
	"encoding/json"
	"os"
	"net/http"
)


type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

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

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}
