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

type EquivTable [][]bool

// package methods ------------------------------------------------------------

func MakeET(numStates int) (EquivTable) {
    et := make([][]bool, numStates, numStates)
    for i := 0; i < numStates; i++ {
        et[i] = make([]bool, numStates, numStates)
    }
    return et
}

//Set a pair of states as distinguished in the Equivalence Table
func (et EquivTable) SetDistinguished(p int, q int) {
    et[p][q] = true
    et[q][p] = true
    return
}

//Get whether a pair of states are distinguished in the Equivalence Table
func (et EquivTable) Distinguished(p int, q int) bool {
    if p > q {
        k := p
        p = q
        q = k
    }
    return et[p][q]
}

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
