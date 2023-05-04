package basic

import (
    "fmt"
    "testing"
)

type MyStruct struct {
    Value int
}

type MyInterface interface {
    GetValue() int
}

func (x *MyStruct) GetValue() int {
    return x.Value
}

func TestNil(t *testing.T) {
    var ptr *MyStruct = (*MyStruct)(nil)
    var val MyInterface = ptr    
    fmt.Println(ptr)  
    fmt.Println(val)   
    fmt.Println(val != nil) // true
}