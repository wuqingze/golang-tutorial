package util

import (
    "testing"
)

func TestGood(t *testing.T){
    if 1==2{
        t.Errorf("this is from ./util/helper_00_test.go:TestError")
    }
}
