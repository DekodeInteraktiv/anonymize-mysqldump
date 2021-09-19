package log

import (
	"log"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	got := New()
	want := &log.Logger{}

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("Expected %+v got %+v", want, got)
	}
}

func TestFlags(t *testing.T) {
	got := New()
	want := &log.Logger{}
	want.SetFlags(log.Ldate | log.Ltime)

	if got.Flags() != want.Flags() {
		t.Errorf("Expected %+v got %+v", want, got)
	}
}

func TestPrefix(t *testing.T) {
	got := New()
	want := &log.Logger{}

	if got.Prefix() != want.Prefix() {
		t.Errorf("Expected %+v got %+v", want, got)
	}
}

func TestWriter(t *testing.T) {
	got := New()
	want := &log.Logger{}
	want.SetOutput(buf)

	if got.Writer() != want.Writer() {
		t.Errorf("Expected %+v got %+v", want, got)
	}
}
