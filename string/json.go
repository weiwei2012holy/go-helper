package string

import "encoding/json"

// JsonEncode json_encode
func JsonEncode(data any) (string, error) {
    jsons, err := json.Marshal(data)
    return string(jsons), err
}

//JsonDecode json_decode
func JsonDecode[V any](s string, data *V) error {
    return json.Unmarshal([]byte(s), data)
}
