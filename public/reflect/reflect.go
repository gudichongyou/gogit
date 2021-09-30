package reflect

import (
	"fmt"
	"reflect"
)

func GetColnameandvale(mpa map[interface{}]interface{}) string {
	for indx, va := range mpa {
		t1 := reflect.TypeOf(indx)
		t2 := reflect.TypeOf(va)
		fmt.Println(t1, t2)

	}
	return ""
}
