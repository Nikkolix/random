package github.com/Nikkolix/random

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAny(t *testing.T) {
	for i := 0; i < 100; i++ {
		v := Any(10)
		fmt.Println(i)
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(v)
		fmt.Println()
	}
}
