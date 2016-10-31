// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package customer

import (
	"html/template"
	"path"
	"path/filepath"
	"runtime"
)

// T contains the set of cached templates for use.
var T *template.Template

func init() {

	// Load the templates from the specified location.
	pattern := filepath.Join(currentDirectory(), "templates", "*.html")
	T = template.Must(template.ParseGlob(pattern))
}

// Returns the current directory we are running in.
func currentDirectory() string {

	// Locate the current directory for the site.
	_, fn, _, _ := runtime.Caller(1)
	return path.Dir(fn)
}
