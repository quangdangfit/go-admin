package utils

import (
	jsoniter "github.com/json-iterator/go"
)

// JSON
var (
	json              = jsoniter.ConfigCompatibleWithStandardLibrary
	JSONMarshal       = json.Marshal
	JSONUnmarshal     = json.Unmarshal
	JSONMarshalIndent = json.MarshalIndent
	JSONNewDecoder    = json.NewDecoder
	JSONNewEncoder    = json.NewEncoder
)

// JSONMarshalToString JSON
func JSONMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}
