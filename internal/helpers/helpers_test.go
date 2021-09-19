package helpers

import (
	"reflect"
	"testing"

	"github.com/xwb1989/sqlparser"
)

func TestNew(t *testing.T) {
	got := GetFakerFuncs()
	want := map[string]func(*sqlparser.SQLVal) *sqlparser.SQLVal{}

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("Expected %+v got %+v", want, got)
	}
}
