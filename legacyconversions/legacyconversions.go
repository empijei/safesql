// Package legacyconversions can be used to atomically switch to safesql.
// This allows to easily automate the rewrite of all uses of database/sql into safesql.
// After that step is done, imports of database/sql should be banned, and legacyconversions gradually removed.
// An allowlist of legacyconversions should be created to make sure no new ones are added.
package legacyconversions

import (
	"github.com/empijei/safesql"
	"github.com/empijei/safesql/internal/raw"
)

var safesqlStringCtor = raw.StringCtor.(func(string) safesql.String)

// UnsafeSQLString riskily promotes the given string to a trusted string.
// Uses of this function should only be introduced to begin a migration to safesql
// and should eventually be removed.
func UnsafeSQLString(trusted string) safesql.String {
	return safesqlStringCtor(trusted)
}
