package dfasim

import (
	"bytes"
	"fmt"
)

//Utility structures and their functions for package

// interfaces -----------------------------------------------------------------
type StateSet interface {
    RandomMember() State
    IsMember(State) bool
    Members() []State
}

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


type NFATransPair struct {
    States  StateSet
    Symbol  string
}

type Trace []TransPair

type EquivTable [][]int

type TransTable map[TransPair]EquivSet

// package methods ------------------------------------------------------------

func (es *EquivSet) String() string{
    var es_string bytes.Buffer
    es_string.WriteString(fmt.Sprintf("Set: {"))
    for _, el := range es.Members() {
        es_string.WriteString(fmt.Sprintf("%v,", el.Name))
    }
    es_string.WriteString(fmt.Sprintf("}\n"))

    return es_string.String()
}

func (tt *TransTable) AddTransition(in State, a string, out State) {
    if out_set, ok := (*tt)[TransPair{in, a}]; ok {
        //A transition from this state on this input exists, add to the set
        out_set.AddMember(out)
        (*tt)[TransPair{in, a}] = out_set
    } else {
        (*tt)[TransPair{in, a}] = EquivSet{out:*new(struct{}),}
    }
}

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

func SetEqual(s1 StateSet, s2 StateSet) bool {
    if len(s1.Members()) != len(s2.Members()) {
        return false
    }
    for _,memb := range s1.Members() {
        if !s2.IsMember(memb) {
            return false
        }
    }
    return true
}

// IsSubset checks if s1 is a subset of s2
func IsSubset(s1 StateSet, s2 StateSet) bool {
    for _,memb := range s1.Members() {
        if !s2.IsMember(memb) {
            return false
        }
    }
    return true
}

// Return the union of two sets of states
func Union(s1 *EquivSet, s2 *EquivSet) EquivSet {
    res := make(EquivSet)
    
    for _,memb := range (*s1).Members() {
        res.AddMember(memb)
    }
    for _,memb := range (*s2).Members() {
        res.AddMember(memb)
    }
    return res
}

func (es EquivSet) RandomMember() State {
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

func (es *EquivSet) DelMember(st State) {
    m := *es
    delete(m, st)
}

func (es EquivSet) Members() []State {
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
