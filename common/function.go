package common

import "encoding/json"

func JsonEncode(v any) (res []byte) {
	data, _ := json.Marshal(v)
	return data
}

func JsonDecode(v string, res any) {
	err := json.Unmarshal([]byte(v), res)
	if err != nil {
		panic(err)
	}
}
