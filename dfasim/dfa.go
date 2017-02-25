/*Deterministic Finite Automaton
*Simulate the behavior of a DFA
 */
package dfasim

import "fmt"

// A struct that defines a particular deterministic finite automaton
type DFA struct {
	states          []state
	state0          state
	alpha           string
	transitiontable map[transpair]state
}

// Build a valid DFA and return a pointer to it
func NewDFA(states []state, state0 state, alpha string, tt map[transpair]state) (*DFA, error) {

	if states == nil {
		return nil, &DFAError{DFAMissingParams, "states can't be empty"}
	}
	if state0.name == "" {
		return nil, &DFAError{DFAMissingParams, "must pass in state0"}
	}
	if alpha == "" {
		return nil, &DFAError{DFAMissingParams, "alphabet can't be empty"}
	}
	if tt == nil {
		return nil, &DFAError{DFAMissingParams, "transition table can't be empty"}
	}

	//Create a 'set' of the input states, and alphabet
	states_set := make(map[string]struct{})
	var finalfound bool
	for _, st := range states {
		states_set[st.name] = *new(struct{})
		finalfound = finalfound || st.final
	}

	//Verify that at least one final state exists
	if !finalfound {
		return nil, &DFAError{DFAInvalidParams, "no final state"}
	}

	//Verify that state0 is in the set of states
	if _, ok := states_set[state0.name]; !ok {
		return nil, &DFAError{DFAInvalidParams, "state0 not in states"}
	}

	//Verify that each state has a transition for each letter in alphabet
	for _, st := range states {
		for _, a := range alpha {
			if _, ok := tt[transpair{st, string(a)}]; !ok {
				return nil, &DFAError{DFAInvalidParams,
					fmt.Sprintf("no transition from state %v on %s", st, a)}
			}
		}
	}

	newdfa := new(DFA)
	newdfa.states = states
	newdfa.state0 = state0
	newdfa.alpha = alpha
	newdfa.transitiontable = tt
	return newdfa, nil
}

/* DFA.DeltaFunc returns the result of the extended transition function for
 * the calling DFA.
 *
 */
func (dfavar *DFA) DeltaFunc(st state, w string, tr *trace) (final state, ok bool) {

	//If empty string, end recursion
	if w == "" {
		final = st
		ok = true
		return
	}

	//Start of computation, begin at state0
	if st.name == "" {
		st = dfavar.state0
	}

	//Split character off of w; w = au
	a := string(w[0])
	u := string(w[1:])

	//Perform transition
	nextSt, ok := dfavar.transitiontable[transpair{st, a}]
	if !ok {
		final = nextSt
		return
	}

	//Add resulting state, remaining str to trace
	tr.addComputation(nextSt, u)

	return dfavar.DeltaFunc(nextSt, u, tr)
}
