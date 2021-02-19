package Week_05

import (
	"fmt"
	"reflect"
	"testing"
)

func Test005(t *testing.T) {
	ttt := myTry{
		name: "123",
	}

	ttt.changeForMethod()
	fmt.Println(ttt.name)
}

type myTry struct {
	name string
}

func (ttt *myTry) changeForMethod() {
	ttt.name = "tqh"
	fmt.Println(ttt.name)
}

func Test006(t *testing.T) {
	t.Log(reflect.TypeOf(byte('a')))
}

//func change(x *myTry) {
//	x.name = "tqh"
//	fmt.Println(x.name)
//}
