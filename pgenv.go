// Package pgenv provides the ability to get default postgres
// credentials from the users env
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

// SetDefaults sets a few common default variables based on os env
func (c ConnStr) SetDefaults() ConnStr {
	c["user"] = getVar("PGUSER", "postgres")
	c["dbname"] = getVar("PGDATABASE", "postgres")
	c["host"] = getVar("PGHOST", "localhost")
	c["port"] = getVar("PGPORT", "5432")
	c["password"] = getVar("PGPASSWD", "")
	c["sslmode"] = getVar("PGSSLMODE", "disabled")

	return c
}

// ToString returns a string that can be dropped into sql.Open
func (c ConnStr) ToString() string {
	var s []string
	for key, val := range c {
		if val != "" {
			s = append(s, key+"="+val)
		}
	}

	return strings.Join(s, ", ")
}