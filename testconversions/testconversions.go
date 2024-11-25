//go:build test

package testconversions

import (
	"github.com/empijei/safesql"
	"github.com/empijei/safesql/internal/raw"
)

var safesqlStringCtor = raw.StringCtor.(func(string) safesql.String)

// SQLStringForTests riskily promotes the given string to a trusted string.
// Uses of this function should only be introduced in test files.
func SQLStringForTests(trusted string) safesql.String {
	return safesqlStringCtor(trusted)
}
