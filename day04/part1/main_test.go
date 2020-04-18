package main

import (
	"fmt"
	//"io/ioutil"
	//"os"
	"testing"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	//realInput, err := ioutil.ReadFile("../input.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	tests := []struct {
		input string
		want  bool
	}{
		{
			input: "111111",
			want:  true,
		},
		{
			input: "223550",
			want:  false,
		},
		{
			input: "123789",
			want:  false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := accepted(test.input)

			if test.want != got {
				t.Errorf("Solution was incorrect, got: %t, want: %t - %s", got, test.want, test.input)
			}
		})
	}
}
