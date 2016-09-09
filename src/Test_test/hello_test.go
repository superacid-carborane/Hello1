package hello

import "testing"
import "flag"

var (
	path = flag.String("path", "", "The path to the file")
)

func TestHello(t *testing.T) {
	expected := ".txt"
	actual := hello(".txt")
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestString(t *testing.T) {
	if *path != "hello" {
		t.Error("string flag should be `hello`, is ", *path)
	}
}
