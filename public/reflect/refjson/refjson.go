package refjson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func init() {

}

type AnyJson interface{}

func Json2Inf(str string) (interface{}, error) {
	var ijson interface{}
	bstr := []byte(str)
	err := json.Unmarshal(bstr, &ijson)
	if err != nil {
		fmt.Println(err)
		return ijson, err
	} else {
		return ijson, err
	}

}
func Inf2Json(i interface{}) (string, error) {
	bstr, err := json.Marshal(i)
	return string(bstr), err

}
func Intf2Mapstr(key string, i interface{}) map[string]interface{} {
	var mps = make(map[string]interface{})
	mps[key] = i
	return mps

}

//AssertAndSetVal 断言将格式化的json串对象赋新值。
// i: json转成的任意类型，vvalue: 当i成员下的key==vvalue中的key时候，且类型一致，将vvalue值 赋给i子孙key的值。
func AssertAndSetVal(i *interface{}, vvalue map[string]interface{}) string {
	switch t := (*i).(type) {
	case map[string]interface{}:
		for k, v := range t {
			switch t1 := v.(type) {
			case map[string]interface{}:
				value1, ok := vvalue[k]
				if ok {
					switch value1.(type) {
					case map[string]interface{}:
						(*i).(map[string]interface{})[k] = value1
					default:
						(*i).(map[string]interface{})[k] = nil
					}

				} else {
					addint := ((*i).(map[string]interface{})[k])
					AssertAndSetVal(&addint, vvalue)
					(*i).(map[string]interface{})[k] = addint

				}
			case []interface{}:
				value2, ok2 := vvalue[k]
				if ok2 {
					switch value2.(type) {
					case []map[string]interface{}:
						(*i).(map[string]interface{})[k] = value2
					case []interface{}:
						(*i).(map[string]interface{})[k] = value2
					default:
						(*i).(map[string]interface{})[k] = nil
					}

				} else {
					for k1, v1 := range t1 {
						switch v1.(type) {
						case map[string]interface{}:
							addint2 := (*i).(map[string]interface{})[k].([]interface{})[k1]
							AssertAndSetVal(&addint2, vvalue)
							(*i).(map[string]interface{})[k].([]interface{})[k1] = addint2
						case []interface{}:
							addint3 := (*i).(map[string]interface{})[k].([]interface{})[k1]
							AssertAndSetVal(&addint3, vvalue)
							(*i).(map[string]interface{})[k].([]interface{})[k1] = addint3
						}
					}
				}
			default:
				value1, ok := vvalue[k]
				if ok {
					(*i).(map[string]interface{})[k] = value1
				}
			}
		}
	case []interface{}:
		for k1, v1 := range t {
			switch v1.(type) {
			case map[string]interface{}:
				addint2 := (*i).([]interface{})[k1]
				AssertAndSetVal(&addint2, vvalue)
				(*i).([]interface{})[k1] = addint2
			case []interface{}:
				AssertAndSetVal(&(*i).([]interface{})[k1], vvalue)
			default:

			}
		}

	}
	ret, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(ret)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
