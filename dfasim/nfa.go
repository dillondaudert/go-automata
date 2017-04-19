/* Nondeterministic Finite Automata
 * Simulate the behavior of NFAs and NFA-lambdas
 */

package dfasim

import (
    "strings"
    "fmt"
    "bytes"
)

// Automaton interface for all machines 
type Automaton interface {
    DeltaFunc(string, *Trace) (bool, bool)
}

// A struct for a particular nondeterministic finite automaton with lambda transitions
type NFA_l struct {
    States []State
    State0 State
    Alpha string
    TTable map[TransPair]EquivSet
}

// Perform the extended transition function for an NFA-l
func (nfavar NFA_l) DeltaFunc(w string, tr *Trace) (bool, bool) {
    //A set of states, demonstrating the current state of the automaton
    var tmp EquivSet
    curr_state := make(EquivSet)
    next_state := make(EquivSet)

    curr_state.AddMember(nfavar.State0)
    curr_state = nfavar.Lclosure(&curr_state)

    //Loop through the string
    for i := 0; i < len(w); i++ {
        var a = string(w[i])

        //Verify that the letter is in the alphabet
        if !strings.Contains(nfavar.Alpha, a) {
            return false, false
        }

        //For each state in curr_state, perform transition on next character
        for _, state := range curr_state.Members() {
            if res, ok := nfavar.TTable[TransPair{state, a}]; !ok {
                continue
            } else {
                //Add results of computation to next_state, remove that state from curr_state
                for _, res_state := range res.Members() {
                    next_state.AddMember(res_state)
                }
            }
            curr_state.DelMember(state)
        }
        /*
        //If there were no valid transitions for this input symbol, this string isn't in the alphabet
        if len(next_state.Members()) == 0 {
            return false, true
        }*/

        next_state = nfavar.Lclosure(&next_state)
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

//Compute the lambda closure for a particular set of states in an NFA-l
func (nfavar NFA_l) Lclosure(stset *EquivSet) (EquivSet) {
    //Make a copy of the set of states
    lclose := *stset
    iterate := true
    for iterate {
        iterate = false
        //For each state in the set, perform lambda transitions
        for _,state := range lclose.Members() {
            //if transition exists
            if rset, ok := nfavar.TTable[TransPair{state, "lambda"}]; ok {
                //if states not already in lclosure
                if !IsSubset(rset, lclose) {
                    lclose = Union(&lclose, &rset)
                    iterate = true
                }
            }
            
        }
    }
    return lclose
}

// A struct for a particular nondeterministic finite automaton
type NFA struct {
    States  []State
    State0  State
    Alpha   string
    TTable  map[TransPair]EquivSet
}

/* Perform the extended transition function of a Nondeterministic Finite Automaton
 */
func (nfavar NFA) DeltaFunc(w string, tr *Trace) (bool, bool) {
    //A set of states, demonstrating the current state of the automaton
    curr_state := make(EquivSet)
    next_state := make(EquivSet)
    var tmp EquivSet
    curr_state.AddMember(nfavar.State0)

    //Loop through the string
    for i := 0; i < len(w); i++ {
        var a = string(w[i])

        //Verify that the letter is in the alphabet
        if !strings.Contains(nfavar.Alpha, a) {
            return false, false
        }

        //For each state in curr_state, perform transition on next character
        for _, state := range curr_state.Members() {
            if res, ok := nfavar.TTable[TransPair{state, a}]; !ok {
                continue
            } else {
                //Add results of computation to next_state, remove that state from curr_state
                for _, res_state := range res.Members() {
                    next_state.AddMember(res_state)
                }
            }
            curr_state.DelMember(state)
        }
        //If there were no valid transitions for this input symbol, this string isn't in the alphabet
        if len(next_state.Members()) == 0 {
            return false, true
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

/* NFA_l implements the Stringer interface
 * Return a nicely formatted transition table
 */
func (nfal *NFA_l) String() string {
	var tt_string bytes.Buffer
	tt_string.WriteString(fmt.Sprintf("    |"))
	for _, a := range nfal.Alpha {
		tt_string.WriteString(fmt.Sprintf(" %3s |", string(a)))
	}
    tt_string.WriteString(fmt.Sprintf(" %6s ", "lambda"))
	tt_string.WriteString("\n")
	for _, state := range nfal.States {
		tt_string.WriteString(fmt.Sprintf("%3s |", state.Name))
		for _, a := range nfal.Alpha {
			out_state := nfal.TTable[TransPair{state, string(a)}]
			tt_string.WriteString(fmt.Sprintf(" %s |", out_state))
		}
        out_states := nfal.TTable[TransPair{state, "lambda"}]
        tt_string.WriteString(fmt.Sprintf(" %s ", out_states))
		if state.Final == true {
			tt_string.WriteString(" *")
		}
		if state == nfal.State0 {
			tt_string.WriteString(" <-")
		}
		tt_string.WriteString("\n")
	}
	return tt_string.String()
}

