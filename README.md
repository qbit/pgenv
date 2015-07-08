Simple utility to import PostgreSQL environment variables and create a connection string from them.

# pgenv
--
    import "github.com/qbit/pgenv"

Package pgenv provides a quick way to create db connection strings from the
running users environment variables.

## Usage

#### type ConnStr

```go
type ConnStr map[string]string
```

ConnStr string map to relate the key values pairs of a pg connection string

#### func (ConnStr) SetDefaults

```go
func (c ConnStr) SetDefaults() ConnStr
```
SetDefaults pulls common environment variables and sets them for later use.

#### func (ConnStr) ToString

```go
func (c ConnStr) ToString() string
```
ToString converts the string map to a sql.Open() compatable connection string.

Example:

```
var c = pgenv.ConnStr{}

c.SetDefaults() 		// Grab env vars from the env with smart defaults
c["dbname"] = "potato"		// Manually set a var to a new value
c.ToString()			// Dump out a string that can be passed to sql.Open

```
