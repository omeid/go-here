//Package here provides source code location of the file it is being called from!
package here

import (
	"path/filepath"
	"runtime"
)

// Dir provides the base path of the caller filepath.
// So if you call here.Dir from $GOPATH/blah/bleh.go you will get $GOPATH/blah
func Dir() string {
	_, filename, _, _ := runtime.Caller(1)

	return filepath.Dir(filename)
}

// Abs provides the absolute path of the of the provided path relative to the caller.
// So if you call `here.Abs("../testdata/data.json")` from $GOPATH/project/helper/bleh.go you will get $GOPATH/project/testdata/data.json
func Abs(path string) string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return ""
	}

	return filepath.Join(filepath.Dir(filename), path)
}
