package json

import jsoniter "github.com/json-iterator/go"

var jiter jsoniter.API

func init() {
	jiter = jsoniter.Config{
		IndentionStep:                 4,
		EscapeHTML:                    true,
		UseNumber:                     true,
		ObjectFieldMustBeSimpleString: true, // do not unescape object field
	}.Froze()
}

func FromStr(str string, v interface{}) error {
	return jiter.UnmarshalFromString(str, v)
}

func FromBytes(byte []byte, v interface{}) error {
	err := jiter.Unmarshal(byte, &v)
	return err
}

func ToStr(v interface{}) (string, error) {
	str, err := jiter.MarshalToString(v)
	return str, err
}
