package abstractfactory

import (
	"reflect"
	"testing"
)

func TestAdidasShoe(t *testing.T) {
	adidasFactory, _ := NewSportFactory("adidas")
	adidasShoe := adidasFactory.makeShoe()
	logo := adidasShoe.getLogo()
	size := adidasShoe.getSize()
	if !reflect.DeepEqual(logo, "adidas") {
		t.Errorf("expected:%#v, got:%#v", "adidas", logo)
	}
	if !reflect.DeepEqual(size, 14) {
		t.Errorf("expected:%#v, got:%#v", "nike", 16)
	}
}
