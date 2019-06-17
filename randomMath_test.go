package base64Captcha

import (
	"fmt"
	"testing"
)

func TestRandColorZone(t *testing.T) {
	r := randSeed()
	for i := 0; i < 100; i++ {
		a := randColorZone(r, [][]int{
			[]int{55, 200,},
			[]int{55, 200,},
			[]int{55, 200,},
		}, nil)
		fmt.Printf("a=[%v]\n", a)
	}
}
