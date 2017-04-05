/* Nondeterministic Finite Automata
 * Simulate the behavior of NFAs and NFA-lambdas
 */

package dfasim

import (
    "fmt"
)

// A struct for a particular nondeterministic finite automaton
type NFA struct {
    States  []State
    State0  State
    Alpha   string
    TTable  map[TransPair]StateSet
}

/* Perform the extended transition function of a Nondeterministic Finite Automaton
 */
func (nfavar *NFA) DeltaFunc(w string, tr *Trace) (final bool, ok bool) {
    //A set of states, demonstrating the current state of the automaton
    curr_state := make(EquivSet)
    next_state := make(EquivSet)
    var tmp EquivSet
    curr_state.AddMember(nfavar.State0)

    //Loop through the string
    for i := 0; i < len(w); i++ {
        var a = string(w[i])
        //For each state in curr_state, perform transition on next character
        for _, state := range curr_state.Members() {
            if res, ok := nfavar.TTable[TransPair{state, a}]; !ok {
                continue
            } else {
                //Add results of computation to next_state, remove that state from curr_state
                fmt.Printf("Transition (%s, %s) -> {%v}\n", state.Name, a, res.Members())
                for _, res_state := range res.Members() {
                    next_state.AddMember(res_state)
                }
            }
            curr_state.DelMember(state)
        }
        //If there were no valid transitions for this input symbol, this string isn't in the alphabet
        if len(next_state.Members()) == 0 {
            return false, false
        }
        //  Change next_state to curr_state
        tmp = next_state
        next_state = curr_state
        curr_state = tmp
    }

    //Finish iterating over string, check curr_state for final state
    for _, state := range curr_state.Members() {
        if state.Final {
            return true, true
        }
    }

    return false, true

}
