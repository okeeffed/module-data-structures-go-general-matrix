package caps

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	for _, tt := range testCases {
		res := genMatrix(tt.input)

		fmt.Println(res)

		for i, v := range res {
			for j := range v {
				if res[i][j] != tt.expected[i][j] {
					t.Errorf("FAIL: Expected %+v but got %+v", tt.expected, res)
					return
				}
			}
		}

		t.Logf("PASS: Expected %+v and got %+v", tt.expected, res)
	}
}
