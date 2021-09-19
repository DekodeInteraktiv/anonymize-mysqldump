package config

import (
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	version := "1.0.0"
	commit := "master"
	now := time.Now()
	date := now.Format("2006-01-02")
	got := New(version, commit, date)
	want := &Config{}

	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Errorf("Expected %+v got %+v", want, got)
	}
}
