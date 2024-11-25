// Package uncheckedconversions can be used to manually promote unsafe strings to safesql.String.
// Uses of this package should be carefully vetted and allowlisted. A security review should be necessary
// for all calls to this package.
package uncheckedconversions

import (
	"github.com/empijei/safesql"
	"github.com/empijei/safesql/internal/raw"
)

var safesqlStringCtor = raw.StringCtor.(func(string) safesql.String)

// KnownSafeString riskily promotes the given string to a trusted string.
// Uses of this function should be carefully reviewed to make sure that no user input
// can ever be passed to it.
//
// Examples of safe usages are to promote strings stored in external query storages
// under the programmer control, or startup flags.
func KnownSafeString(trusted string) safesql.String {
	return safesqlStringCtor(trusted)
}
