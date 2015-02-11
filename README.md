Simple utility to import PostgreSQL environment variables and create a connection string from them.

Example:

```
var c = pgenv.ConnStr{}

c.SetDefaults() 		// Grab env vars from the env with smart defaults
c["dbname"] = "potato"		// Manually set a var to a new value
c.ToString()			// Dump out a string that can be passed to sql.Open

```
