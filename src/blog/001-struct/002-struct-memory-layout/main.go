package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type layouttest struct {
	b  byte
	v  int32
	f  float64
	v2 int32
}

type compactyouttest struct {
	f  float64
	v  int32
	v2 int32
	b  byte
}

func main() {

	var x layouttest
	var y compactyouttest

	fmt.Printf("Int alignment %v \n", unsafe.Alignof(10))
	fmt.Printf("Int8 aligment %v \n", unsafe.Alignof(int8(10)))
	fmt.Printf("Int16 aligment %v \n", unsafe.Alignof(int16(10)))
	fmt.Printf("Layoutest aligment %v ans size is %v \n", unsafe.Alignof(x), unsafe.Sizeof(x))
	fmt.Printf("Compactlayouttest aligment %v and size is %v \n", unsafe.Alignof(y), unsafe.Sizeof(y))

	fmt.Printf("Type %v has %v fields and takes %v bytes \n", reflect.TypeOf(x).Name(), reflect.TypeOf(x).NumField(), reflect.TypeOf(x).Size())

}
