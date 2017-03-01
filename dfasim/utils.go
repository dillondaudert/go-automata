//Utility structures and their functions for package
package dfasim

import (
	"bytes"
	"fmt"
)

// package structs ------------------------------------------------------------

type EquivSet map[State]struct{}

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

type EquivTable [][]int

// package methods ------------------------------------------------------------

func (et EquivTable) FormatTable(states []State) string {
	var et_string bytes.Buffer
	et_string.WriteString(fmt.Sprintf("    |"))
	for _, st := range states {
		et_string.WriteString(fmt.Sprintf(" %3s |", st.Name))
	}
	et_string.WriteString("\n")
	for i, st1 := range states {
		et_string.WriteString(fmt.Sprintf("%3s |", st1.Name))
		for k, _ := range states {
			if k >= i {
				et_string.WriteString(fmt.Sprintf("     |"))
			} else {
				et_string.WriteString(fmt.Sprintf(" %3d |", et[i][k]))
			}
		}
		et_string.WriteString("\n")
	}

    return et_string.String()
}



func (es EquivSet) RandomMember() (State) {
    for k, _ := range es {
        return k
    }
    return State{}
}

func (es EquivSet) IsMember(st State) bool {
    _, ok := es[st]
    return ok
}

func (es *EquivSet) AddMember(st State) {
    m := *es
    m[st] = *new(struct{})
}

func (es EquivSet) Members() ([]State) {
    membs := make([]State, 0, len(es))
    for memb := range es {
        membs = append(membs, memb)
    }
    return membs
}

//Return a new Equivalence Table with size numStates x numStates
func MakeET(numStates int) EquivTable {
	et := make([][]int, numStates, numStates)
	for i := 0; i < numStates; i++ {
		et[i] = make([]int, numStates, numStates)
	}

	for i := 0; i < numStates; i++ {
		for k := 0; k < numStates; k++ {
			et[i][k] = -1
		}
	}
	return et
}

//Set a pair of states as distinguished in the Equivalence Table
func (et EquivTable) SetDistinguished(p int, q int, round int) {
	et[p][q] = round
	et[q][p] = round
	return
}

//Get whether a pair of states are distinguished in the Equivalence Table
func (et EquivTable) Distinguished(p int, q int) bool {
	return et[p][q] != -1
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
