package common

import "encoding/json"

func ToJsonString(input interface{}) string {
	data, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	return string(data)
}
func ToJsonObject(inputJSON string, inputStruct interface{}) {
	_ = json.Unmarshal([]byte(inputJSON), inputStruct)
}
