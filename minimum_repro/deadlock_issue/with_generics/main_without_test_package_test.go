package with_generics_test

import (
	"os"
	"testing"

	"github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestT1(t *testing.T) {
	_ = with_generics.T1[any]{}
}
