package json

import (
	"github.com/beego/beego/v2/adapter/logs"
	jsoniter "github.com/json-iterator/go"
)

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

func ToJSONStr(v interface{}) string {
	str, err := jiter.MarshalToString(v)
	if err != nil {
		logs.Error("interface{} to string has error {}", err)
		return ""
	}
	return str
}
