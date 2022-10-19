package factory_method

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonConfigParserFactory_CreateParser(t *testing.T) {
	f := NewIRuleConfigParserFactory("json")
	got := f.CreateParser().parse([]byte(""))
	want := "json parse"
	assert.Equal(t, got, want)
}

func TestYamlConfigParserFactory_CreateParser(t *testing.T) {
	f := NewIRuleConfigParserFactory("yaml")
	got := f.CreateParser().parse([]byte(""))
	want := "yaml parse"
	assert.Equal(t, got, want)
}
