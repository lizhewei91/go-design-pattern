package builder

import (
	"reflect"
	"testing"
)

func TestBuilder(t *testing.T) {
	builder := PersonBuilder{}
	person := builder.SetName("lzw").
		SetAge("18").
		SetGender("man").
		SetAddress("xj").
		Build()
	gotName := person.GetName()
	if !reflect.DeepEqual(gotName, "lzw") {
		t.Errorf("expected:%#v, got:%#v", "lzw", gotName)
	}
}
