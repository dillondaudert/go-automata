/*Deterministic Finite Automaton
*Simulate the behavior of a DFA
 */
package dfasim

import "fmt"

// A struct that defines a particular deterministic finite automaton
type DFA struct {
	States          []State
	State0          State
	Alpha           string
	TransitionTable map[TransPair]State
}

// Build a valid DFA and return a pointer to it
func NewDFA(states []State, state0 State, alpha string, tt map[TransPair]State) (*DFA, error) {

	if states == nil {
		return nil, &DFAError{DFAMissingParams, "states can't be empty"}
	}
	if state0.Name == "" {
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
		states_set[st.Name] = *new(struct{})
		finalfound = finalfound || st.Final
	}

	//Verify that at least one final state exists
	if !finalfound {
		return nil, &DFAError{DFAInvalidParams, "no final state"}
	}

	//Verify that state0 is in the set of states
	if _, ok := states_set[state0.Name]; !ok {
		return nil, &DFAError{DFAInvalidParams, "state0 not in states"}
	}

	//Verify that each state has a transition for each letter in alphabet
	for _, st := range states {
		for _, a := range alpha {
			if _, ok := tt[TransPair{st, string(a)}]; !ok {
				return nil, &DFAError{DFAInvalidParams,
					fmt.Sprintf("no transition from state %v on %s", st, a)}
			}
		}
	}

	newdfa := new(DFA)
	newdfa.States = states
	newdfa.State0 = state0
	newdfa.Alpha = alpha
	newdfa.TransitionTable = tt
	return newdfa, nil
}

/*
 * DFA.DeltaFunc returns the result of the extended transition function for
 * the calling DFA.
 */
func (dfavar *DFA) DeltaFunc(st State, w string, tr *Trace) (final State, ok bool) {

	//If empty string, end recursion
	if w == "" {
		final = st
		ok = true
		return
	}

	//Start of computation, begin at State0
	if st.Name == "" {
		st = dfavar.State0
	}

	//Split character off of w; w = au
	a := string(w[0])
	u := string(w[1:])

	//Perform transition
	nextSt, ok := dfavar.TransitionTable[TransPair{st, a}]
	if !ok {
		final = nextSt
		return
	}

	//Add resulting state, remaining str to Trace
	tr.addComputation(nextSt, u)

	return dfavar.DeltaFunc(nextSt, u, tr)
}

/*
 * DFA.Minim turns the calling DFA into an equivalent DFA with the minimum
 * number of states.
 */
func (dfavar *DFA) Minim() {
	//1. Build the table of distinguishable states
	et := MakeET(len(dfavar.States))
	//Create a map from a state to its index to make this simpler
	state_index := make(map[State]int)
	for i, st := range dfavar.States {
		state_index[st] = i
	}
	disting_found := false
	// For every pair of states, if one is final and one is not, distinguish
	for i, st1 := range dfavar.States {
		for j, st2 := range dfavar.States {
			if st1.Final != st2.Final {
				et.SetDistinguished(i, j)
				disting_found = true
			}
		}
	}
	// If a distinguishable state was found last round:
	for disting_found == true {
		disting_found = false
		// Loop:
		//      For every pair of states (p, q)
		for i, st1 := range dfavar.States {
			for j, st2 := range dfavar.States {
				if et.Distinguished(i, j) == true {
					continue
				}
				//          For every symbol a in the alphabet,
				for _, a := range dfavar.Alpha {
					//              compare p' = delta(p, a) to q' = delta(q, a)
					st1_prime := dfavar.TransitionTable[TransPair{st1, string(a)}]
					st2_prime := dfavar.TransitionTable[TransPair{st2, string(a)}]

					//do not compare
					if st1_prime == st2_prime {
						continue
					}

					i_prime := state_index[st1_prime]
					j_prime := state_index[st2_prime]
					//              if (p', q') are distinguished, then (p, q) is distinguished
					if et.Distinguished(i_prime, j_prime) == true {
						et.SetDistinguished(i, j)
						disting_found = true
					}
				}
			}
		}
	}

	//2. Coalesce the equivalent states and build a new transition table
    sets := make([]EquivSet, 0, len(dfavar.States))
	for i, st1 := range dfavar.States {
		for j, st2 := range dfavar.States {
            new_set := true
			if !et.Distinguished(i, j) {
                //For all the current equiv sets, check if either are members
                for k, set := range sets {
                    if set.IsMember(st1) || set.IsMember(st2) {
                        set.AddMember(st1)
                        set.AddMember(st2)
//                        fmt.Printf("%v and %v added to equiv set %d\n", st1, st2, k)
                        new_set = false
                    } 
                }
                if new_set == true {
                    add_set := make(EquivSet)
                    add_set.AddMember(st1)
                    add_set.AddMember(st2)
                    sets = append(sets, add_set)
//                    fmt.Printf("%v and %v added to NEW equiv set\n", st1, st2)                    
                }
			}
		}
	}
    for i, set := range sets {
        fmt.Printf("Set %d: %v\n--\n", i, set)
    }
}
