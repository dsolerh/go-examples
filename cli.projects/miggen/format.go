package main

// Config holds configuration for
// generating the migration file
type Config struct {
	// DateFormat is the format used for the migration
	// filename Ex. "2006-01-02T15:04:05<separator><migration-name>.<extension>"
	DateFormat string
	// Separator is the string used for separating the
	// migration name and the date
	Separator string
	// Extension is the extension for the
	// migration filename
	Extension string
	// Types is an colection of MigType
	Types MigType
}

// MigType is used for hold a set of keys (types)
// and values, the values can be the body of the migration
// skeleton or a filename that points to a template
type MigType map[string]string

const (
	dateFormat = "2006-01-02T15:04:05"
	separator  = "-"
	extension  = "js"
)
