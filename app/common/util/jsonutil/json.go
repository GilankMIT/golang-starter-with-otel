package jsonutil

import "encoding/json"

func ToJsonString(data any) string {
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}
