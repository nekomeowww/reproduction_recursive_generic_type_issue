package recursive_generic_type_issue_reproduction_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nekomeowww/recursive_generic_type_issue_reproduction"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestTypeA_GetValue(t *testing.T) {
	t.Parallel()

	valueFromTypeA := recursive_generic_type_issue_reproduction.NewTypeA[int]().WithValue(1).GetValue()
	assert.Equal(t, "random name: 1", valueFromTypeA)
	valueFromTypeB := recursive_generic_type_issue_reproduction.NewTypeB[int]().WithValue(1).GetValue()
	assert.Equal(t, "TypeB: 1", valueFromTypeB)
}
