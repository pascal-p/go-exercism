// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(str string) string {
	rexpr := regexp.MustCompile(`[\s,\.:;\-_"]`) 
	v := rexpr.Split(str, -1) // -1 to tell no limits for the number of substrings

	new_str := ""
	for _, s := range v {
		if strings.TrimSpace(s) == "" {
			continue
		}
		new_str += strings.ToUpper(string(s[0]))		
	}	
	return new_str
}
