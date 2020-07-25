package main
import (
	"bytes"
	"testing"
)
func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "danizavtz")

	got := buffer.String()
	want := "Hello, danizavtz"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}