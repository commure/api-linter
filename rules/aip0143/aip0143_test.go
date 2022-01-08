package aip0143

import (
	"testing"

	"github.com/commure/api-linter/lint"
)

func TestAddRules(t *testing.T) {
	if err := AddRules(lint.NewRuleRegistry()); err != nil {
		t.Errorf("AddRules got an error: %v", err)
	}
}
