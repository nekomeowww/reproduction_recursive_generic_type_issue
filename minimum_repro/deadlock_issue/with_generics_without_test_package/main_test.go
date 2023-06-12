package with_generics_without_test_package

import (
	"fmt"
	"testing"
)

func TestT1(t *testing.T) {
	if fmt.Sprintf("%T", T1[any]{}) != "with_generics_without_test_package.T1[interface {}]" {
		t.Error("T1[any] is not with_generics_without_test_package.T1[interface {}]")
	}
}
