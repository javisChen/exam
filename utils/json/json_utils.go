package json

import jsoniter "github.com/json-iterator/go"

func FromStr(str string, v interface{}) error {
	return jsoniter.UnmarshalFromString(str, v)
}

func FromBytes(byte []byte, v interface{}) error {
	err := jsoniter.Unmarshal(byte, &v)
	return err
}

func ToStr(v interface{}) (string, error) {
	str, err := jsoniter.MarshalToString(v)
	return str, err
}
