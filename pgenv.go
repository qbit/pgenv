// Package pgenv provides a quick way to create db connection strings
// from the running users environment variables.
package pgenv

import (
	"os"
	"strings"
)

// ConnStr string map to relate the key values pairs of a pg connection string
type ConnStr map[string]string

/*
Options are from: http://www.postgresql.org/docs/current/static/libpq-envars.html

Future plans are to support more options.
*/

func getVar(envVal string, def string) string {
	var ret string
	ret = os.Getenv(envVal)
	if ret == "" {
		ret = def
	}
	return ret
}

// SetDefaults pulls common environment variables and sets them for later use.
func (c ConnStr) SetDefaults() ConnStr {
	c["user"] = getVar("PGUSER", getVar("USER", "postgres"))
	c["dbname"] = getVar("PGDATABASE", getVar("USER", "postgres"))
	c["host"] = getVar("PGHOST", "localhost")
	c["port"] = getVar("PGPORT", "5432")
	c["password"] = getVar("PGPASSWD", "")
	c["sslmode"] = getVar("PGSSLMODE", "disable")

	return c
}

// ToString converts the string map to a sql.Open() compatable connection string.
func (c ConnStr) ToString() string {
	var s []string
	for key, val := range c {
		if val != "" {
			s = append(s, key+"="+val)
		}
	}

	return strings.Join(s, " ")
}
