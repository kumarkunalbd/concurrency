package distchar

import (
	"fmt"
	"testing"
)

func TestCountDistChars(t *testing.T) {
	inputs := []struct {
		input       string
		expectedout int
	}{
		{
			input:       "aBaBbeCNnO",
			expectedout: 8,
		},
		{
			input:       "aaaAAA",
			expectedout: 2,
		},
	}

	dst := NewDistChar(1, false)
	for i, aninput := range inputs {
		actres := dst.CountDistChar(aninput.input)
		if actres == aninput.expectedout {
			fmt.Printf("test case num=%v passed\n", i)
		} else {
			t.Errorf("error. Expeted:%v <-> got:%v\n", aninput.expectedout, actres)
		}
	}
}

func TestCountDistCharsCAseInsensitive(t *testing.T) {
	inputs := []struct {
		input       string
		expectedout int
	}{
		{
			input:       "aBaBbeCNnO",
			expectedout: 6,
		},
		{
			input:       "aaaAAA",
			expectedout: 1,
		},
	}

	dst := NewDistChar(2, false)
	for i, aninput := range inputs {
		actres := dst.CountDistChar(aninput.input)
		if actres == aninput.expectedout {
			fmt.Printf("test case num=%v passed\n", i)
		} else {
			t.Errorf("error. Expeted:%v <-> got:%v\n", aninput.expectedout, actres)
		}
	}
}

func TestFilter(t *testing.T) {
	inputs := []struct {
		input       string
		expectedout int
	}{
		{
			input:       "aBaB@!~beCNnO",
			expectedout: 6,
		},
		{
			input:       "aa!!!#aAAA",
			expectedout: 1,
		},
	}

	dst := NewDistChar(2, true)
	for i, aninput := range inputs {
		actres := dst.CountDistChar(aninput.input)
		if actres == aninput.expectedout {
			fmt.Printf("test case num=%v passed\n", i)
		} else {
			t.Errorf("error. Expeted:%v <-> got:%v\n", aninput.expectedout, actres)
		}
	}
}
