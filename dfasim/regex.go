/* Regular Expression Parser
 * Turn a regular expression in an automaton
 */

package dfasim

import (
    "strconv"
    "reflect"
)

func ParseRegex(w string, dtype string) (Automaton) {

    ttable := make(TransTable)
    sts := make(EquivSet)
    start := State{"S", false}
    alpha := ""

    RegexToNFAl(w, start, 0, &sts, &ttable, &alpha)
    
    return NFA_l{sts.Members(), start, alpha, ttable}

}

func RegexToNFAl(w string, entry State, pos int, sts *EquivSet, ttable *TransTable, alpha *string) (State, int) {
    //Frame state information
    var (
        curr State
        next State
        leftP State
        rightP State
        innerS State
        innerF State
        finals EquivSet
    )
    //Finals keeps track of the exit point for any parse frame
    finals = make(EquivSet)

    //Set the current state to the entry state
    curr = entry
    //Parse loop
    for i := pos; i < len(w); i++ {
        a := string(w[i])
        switch a {
            case " ":
            
            case "(":
                //Create a left parenthesis state
                leftP = State{"(:"+strconv.Itoa(len(sts.Members())), false}
                sts.AddMember(leftP)
                //Create a lambda transition from current 
                ttable.AddTransition(curr, "lambda", leftP)
                //ttable[TransPair{curr, "lambda"}] = EquivSet{leftP:*new(struct{}),}

                //Create a new entry point
                innerS = State{"S:"+strconv.Itoa(len(sts.Members())), false}
                sts.AddMember(innerS)
                //Create a lambda transition from parenth to entry
                ttable.AddTransition(leftP, "lambda", innerS)
                //ttable[TransPair{leftP, "lambda"}] = EquivSet{innerS:*new(struct{}),}

                //Recurse, assign output to innerF
                innerF, i = RegexToNFAl(w, innerS, i+1, sts, ttable, alpha)
                sts.AddMember(innerF)

                //Now we have read up through the right parenthesis
                rightP = State{"):"+strconv.Itoa(len(sts.Members())), false}
                sts.AddMember(rightP)
                //Create transition
                ttable.AddTransition(innerF, "lambda", rightP)
                //ttable[TransPair{innerF, "lambda"}] = EquivSet{rightP:*new(struct{}),}
                
                curr = rightP

            case ")":
                //Exit up one level
                finals.AddMember(curr)

                innerF = State{"F:"+strconv.Itoa(len(sts.Members())), false}
                //For every state in the set of finals
                for _, fst := range finals.Members() {
                    //Create a lambda transition to innerF
                    ttable.AddTransition(fst, "lambda", innerF)
                    //ttable[TransPair{fst, "lambda"}] = EquivSet{innerF:*new(struct{}),}
                    
                }

                //Return the inner final state, and the current index + 1
                return innerF, i

            case "*":
                //ASSUME: Previously exited a set of parentheses
                //Add lambda transitions
                ttable.AddTransition(leftP, "lambda", rightP)
                ttable.AddTransition(innerF, "lambda", innerS)

            case "+":
                //Add current to finals, reset the curr pointer to the entry point
                finals.AddMember(curr)
                curr = entry

            default:
                (*alpha) += a
                //Concatenate, add state and create transition
                next = State{a+":"+string(i), false}
                sts.AddMember(next)
                ttable.AddTransition(curr, a, next)
                curr = next

        }

    }
    finals.AddMember(curr)
    //Mark each state in finals as a final state
    for _,fst := range finals.Members() {
        sts.DelMember(fst)
        sts.AddMember(State{fst.Name, true})
    }
    
    //Update all transitions in the ttable that involve a final state
    for tkey := range (*ttable) {
        rset,_ := (*ttable)[tkey]
        //If the key has a final state, change
        newkey := tkey
        if finals.IsMember(newkey.State) {
            newkey.State.Final = true
        }
        newset := rset

        for _,state := range rset.Members() {
            if finals.IsMember(state) {
                //Replace nonfinal version with final version
                newstate := state
                newstate.Final = true
                newset.DelMember(state)
                newset.AddMember(newstate)
            }
        }
        if newkey != tkey || !reflect.DeepEqual(newset, rset) {
            delete((*ttable), tkey)
            (*ttable)[newkey] = newset
        }

    }
    return State{}, 0
}
