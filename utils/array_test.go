package utils

import (
	"reflect"
	"testing"
)

func TestInterfaceSliceConversion(t *testing.T) {
	result := InterfaceSliceConversion([]interface{}{"create:items", "update:items", "delete:items"})
	if reflect.TypeOf(result) != reflect.TypeOf([]string{}) {
		t.Errorf("The interface type has not been converted to a string type.")
	}
}

func TestEvery(t *testing.T) {
	result := Every([]string{"create:items", "update:items"}, []string{"create:items", "update:items"})
	if !result {
		t.Errorf("The user does not have permissions.")
	}
}

func TestSome(t *testing.T) {
	result := Some([]string{"create:items", "update:items"}, []string{"update:items", "delete:items"})
	if !result {
		t.Error("The user does not have permissions.")
	}
}
