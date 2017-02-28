//Utility structures and their functions for package
package dfasim

import "fmt"

// package structs ------------------------------------------------------------

type State struct {
	Name  string
	Final bool
}

type TransPair struct {
	State  State
	Symbol string
}

type Trace []TransPair

type StatePair struct {
    State1 State
    State2 State
}

// package methods ------------------------------------------------------------

func (tr TransPair) String() string {
	return fmt.Sprintf("%v, \"%v\"", tr.State.Name, tr.Symbol)
}

/* AddComputation adds a State, string pair to the array slice of
 * transitions.
 */
func (t *Trace) addComputation(st State, w string) {
	*t = append(*t, TransPair{st, w})
}

func (t Trace) String() string {
	var transitions string
	for i, pair := range t {
		if i == 0 {
			transitions += fmt.Sprintf("%v\n", pair)
		} else {
			transitions += fmt.Sprintf("\t-> %v\n", pair)
		}
	}
	return fmt.Sprintf("%v", transitions)
}
