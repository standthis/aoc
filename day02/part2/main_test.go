package main

import (
	"fmt"
	"io/ioutil"
	"testing"
    "os"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	realInput, err := ioutil.ReadFile("../input.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }


	tests := []struct {
		input string
		want  int
	}{
		{
			input: string(realInput),
			want:  19690720,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := run(test.input)

            if test.want != got {
                t.Errorf("Solution was incorrect, got: %d, want: %d.", got, test.want)
            }
		})
	}
}
