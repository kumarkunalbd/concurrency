package infoblox

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	//inputstr := "xyzabyzxsayxzyerfxyasz"
	//tstr := "xyz"

	////

	inputstr1 := "ababababa"
	tstr2 := "ab"

	ans := numberOfAnagram(inputstr1, tstr2)
	fmt.Printf("ans=%v", ans)
}
