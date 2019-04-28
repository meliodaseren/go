package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println("哈囉！世界！")
	fmt.Println("----------")
	timeNow()
	fmt.Println("----------")
	randIntn()
	fmt.Println("----------")
	mathNextAfter()
	fmt.Println("----------")
	mathFuncCheck()
}

func timeNow() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
}

func randIntn() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

func mathNextAfter() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}

func mathFuncCheck() {
	//fmt.Println(unsafe.Sizeof(true))
	//fmt.Println(unsafe.Sizeof(true))
	fmt.Printf("uint8 : 0 ~ %d\n", math.MaxUint8)
	fmt.Printf("uint16: 0 ~ %d\n", math.MaxUint16)
	fmt.Printf("uint32: 0 ~ %d\n", uint32(math.MaxUint32))
	fmt.Printf("uint64: 0 ~ %d\n", uint64(math.MaxUint64))
	fmt.Printf("int8  : %d ~ %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("int16 : %d ~ %d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32 : %d ~ %d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64 : %d ~ %d\n", int64(math.MinInt64), int64(math.MaxInt64))
	fmt.Printf("Integer default type: %s\n", reflect.TypeOf(1))
}
