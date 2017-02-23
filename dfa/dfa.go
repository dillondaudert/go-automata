/*Deterministic Finite Automaton
*Simulate the behavior of a DFA
 */
package dfasim

type State struct {
	Name  string
	Final bool
}

type TransPair struct {
	State  State
	Symbol string
}

/* -------------------------------------------------------------------*/
type Trace []TransPair

/* AddComputation adds a State, string pair to the array slice of
 * transitions.
 */
func (t *Trace) AddComputation(st State, w string) {
	*t = append(*t, TransPair{st, w})
}

/* TO BE COMPLETED
func (t Trace) String() string {

    return fmt.Sprintf(
}
*/
/* -------------------------------------------------------------------*/

// A struct that defines a particular deterministic finite automaton
type DFA struct {
	States          []State
	State0          State
	Alpha           string
	TransitionTable map[TransPair]State
}

//
func (dfa *DFA) DeltaFunc(st State, w string, tr *Trace) (State, bool) {

	//First, verify that the DFA struct is properly initialized
	if dfa.States == nil {
		panic("PANIC: DFA.DeltaFunc called with DFA.States uninitialized!")
	}
	if dfa.State0.Name == "" {
		panic("PANIC: DFA.DeltaFunc called with DFA.State0 uninitialized!")
	}
	if dfa.Alpha == "" {
		panic("PANIC: DFA.DeltaFunc called with DFA.Alpha uninitialized!")
	}
	if dfa.TransitionTable == nil {
		panic("PANIC: DFA.DeltaFunc called with DFA.TransitionTable uninitialized!")
	}

	//If empty string, end recursion
	if w == "" {
		return st, true
	}

	//Start of computation, begin at State0
	if st.Name == "" {
		st = dfa.State0
	}

	//Split character off of w; w = au
	a := string(w[0])
	u := string(w[1:])

	//Perform transition
    nextSt, ok := dfa.TransitionTable[TransPair{st, a}]
    if !ok {
        return State{}, false
    }

	//Add resulting state, remaining str to trace
	tr.AddComputation(nextSt, u)

	return dfa.DeltaFunc(nextSt, u, tr)
}
/*
func RunDFA() {
	//Create some states
	st1 := State{"A", false}
	st2 := State{"B", true}
	fmt.Println(st1, st2)

	//Create some TransPairs
	trPair1 := TransPair{st1, "x"}
	trPair2 := TransPair{st1, "y"}
	trPair3 := TransPair{st2, "x"}
	trPair4 := TransPair{st2, "y"}
	fmt.Println(trPair1, trPair2)

	trTable := map[TransPair]State{
		trPair1: st2,
		trPair2: st1,
		trPair3: st2,
		trPair4: st2,
	}
	fmt.Println(trTable)

}
*/
