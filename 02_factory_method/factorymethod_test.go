package factorymethod

import (
	"reflect"
	"testing"
)

func TestNewGunType(t *testing.T) {
	tests1 := []struct {
		gunType   string
		wantName  string
		wantPower int
	}{
		{"ak47", "ak47 gun", 4},
		{"musket", "musket gun", 1},
	}

	for _, tt := range tests1 {
		t.Run(tt.gunType, func(t *testing.T) { // 使用t.Run()执行子测试
			gun := NewGunType(tt.gunType)
			gotGunName := gun.getName()
			if !reflect.DeepEqual(gotGunName, tt.wantName) {
				t.Errorf("expected:%#v, got:%#v", tt.wantName, gotGunName)
			}
			gotGunPower := gun.getPower()
			if !reflect.DeepEqual(gotGunPower, tt.wantPower) {
				t.Errorf("expected:%#v, got:%#v", tt.wantPower, gotGunPower)
			}
		})
	}

	tests2 := []struct {
		gunType   string
		setName   string
		setPower  int
		wantName  string
		wantPower int
	}{
		{"ak47", "AK47 gun", 6, "AK47 gun", 6},
		{"musket", "Musket gun", 2, "Musket gun", 2},
	}
	for _, tt := range tests2 {
		t.Run(tt.gunType, func(t *testing.T) { // 使用t.Run()执行子测试
			gun := NewGunType(tt.gunType)
			gun.setName(tt.setName)
			gotGunName := gun.getName()
			if !reflect.DeepEqual(gotGunName, tt.wantName) {
				t.Errorf("expected:%#v, got:%#v", tt.wantName, gotGunName)
			}
			gun.setPower(tt.setPower)
			gotGunPower := gun.getPower()
			if !reflect.DeepEqual(gotGunPower, tt.wantPower) {
				t.Errorf("expected:%#v, got:%#v", tt.wantPower, gotGunPower)
			}
		})
	}
}
