package simple_factory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModel1(t *testing.T) {
	parser := NewIRuleConfigParser("json")
	want := "json parser"
	got := parser.parse([]byte{})
	assert.Equal(t, got, want)
}

func TestModel2(t *testing.T) {
	parser := NewIRuleConfigParser("yaml")
	want := "yaml parser"
	got := parser.parse([]byte{})
	assert.Equal(t, got, want)
}
