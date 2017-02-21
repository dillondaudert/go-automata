/*Deterministic Finite Automaton 
*Simulate the behavior of a DFA
*/
package main

import(
    "fmt"
)

type State struct {
    Name string
    Final bool
}

type TransPair struct {
    State State
    Symbol string
}

type DFA struct {
    States, State0 State
    Alpha string
    TransitionTable map[TransPair]State
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

//
func (dfa *DFA) DeltaFunc(st State, w string, tr *Trace) (bool) {
    //If empty string, end recursion
    if w == "" {
        tr.AddComputation(st, "lambda")
        return st.Final
    }

    //Start of computation, begin at State0
    if st.Name == ""{
        st = dfa.State0
    }

    //Split character off of w; w = au 
    a := string(w[0])
    u := string(w[1:])

    //Perform transition
    nextSt := dfa.TransitionTable[TransPair{st, a}]

    //Add resulting state, remaining str to trace
    tr.AddComputation(nextSt, u)

    return dfa.DeltaFunc(nextSt, u, tr)
}

func main() {
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

    var tr Trace

    fmt.Println("Trace: ", tr)
    tr.AddComputation(st1, "ab")
    fmt.Println("Trace: ", tr)
    tr.AddComputation(st2, "b")
    fmt.Println("Trace: ", tr)

}

//type TransitionTable map[TransPair] State


