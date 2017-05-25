package here

import (
	"os"
	"path/filepath"
	"testing"
)

var base = filepath.Join(os.Getenv("GOPATH"), "src/github.com/omeid/go-here")

func TestDir(t *testing.T) {

	expect := filepath.Join(base, "here_test.go")

	path := Dir()

	if path != expect {
		t.Fatalf("Expected Path:\n%s\nGot:\n%s", expect, path)
	}
}

func TestAbs(t *testing.T) {

	expect := filepath.Join(base, "../go-yarn")

	path := Abs("../go-yarn")

	if path != expect {
		t.Fatalf("Expected Path:\n%s\nGot:\n%s", expect, path)
	}

}
