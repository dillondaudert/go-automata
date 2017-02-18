/*Deterministic Finite Automaton 
*Simulate the behavior of a DFA
*/
package dfa

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

//
func (dfa *DFA) DeltaFunc(state State, w string, trace []string) (State, []string) {
    //If empty string, end recursion
    if w == "" {
        return state, trace
    }

    //Start of trace, begin at State0
    if state == nil{
        state = DFA.State0
    }
    
    //Split character off of w, 
    a := string(w[0])
}

//type TransitionTable map[TransPair] State


