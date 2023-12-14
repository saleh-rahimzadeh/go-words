/*┄─────────────────────────────────────────────────────────────────────────────────────────────╮
┊   Copyright (C) 2023 Saleh Rahimzadeh                                                         │
│   https://github.com/saleh-rahimzadeh/go-words                                                │
│   All rights reserved.                                                                        ┊
╰─────────────────────────────────────────────────────────────────────────────────────────────┄*/

package gowords

//──────────────────────────────────────────────────────────────────────────────────────────────────

// Words the interface to specifiy required methods to get and find words
type Words interface {
	// Get search for a name then return value if found, else return empty string
	Get(string) string
	// Find search for a name then return value and `true` if found, else return empty string and `false`
	Find(string) (string, bool)
}
