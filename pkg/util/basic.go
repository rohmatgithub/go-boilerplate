package util

import "encoding/json"

func ObjectToString(object interface{}) string {
	json, err := json.Marshal(object)
	if err != nil {
		return ""
	}
	return string(json)
}
