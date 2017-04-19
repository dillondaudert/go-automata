//Error types and functions for the dfasim package
package regex

import "fmt"

// package constants ----------------------------------------------------------

const (
	DFAMissingParams = "Missing parameters"
	DFAInvalidParams = "Invalid parameters"
)

// package structs ------------------------------------------------------------
type DFAError struct {
	Type    string
	Message string
}

// package methods ------------------------------------------------------------

func (e *DFAError) Error() string {
	return fmt.Sprintf("DFA error encountered: %s: %s\n", e.Type, e.Message)
}
