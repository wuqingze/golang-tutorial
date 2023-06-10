package util

import (
    "testing"
)

func TestAdd(t *testing.T) {
    testCases := []struct {
        a        int
        b        int
        expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {3, -3, 0},
    }

    for _, tc := range testCases {
        result := add(tc.a, tc.b)
        if result != tc.expected {
            t.Errorf("add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
        }
    }
}

func TestError(t *testing.T){
    if 1==2{
        t.Errorf("this is from ./util/helper_00_test.go:TestError")
    }
}
