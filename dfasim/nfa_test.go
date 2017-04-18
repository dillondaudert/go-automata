package dfasim

import (
    "fmt"
    "testing"
)

func TestNFA(t *testing.T) {

    //Test the NFA struct
    //L(M) = a+(b* U c*)d
    nfa_cases := get_test_nfas()
    for i, nfa_case := range nfa_cases {
        for j, pair := range nfa_case.Pairs {
            tr := new(Trace)
            t.Run(fmt.Sprintf("NFA/%d:%d", i,j), func(t *testing.T) {
                accept, ok := nfa_case.TestNFA.DeltaFunc(pair.String, tr)
                if accept != pair.Accept || ok != pair.Valid {
                    t.Error("NFA Test Error: Expected (", pair.Accept,
                        ",", pair.Valid, "), Got (", accept, ", ", ok,
                        ")\n")
                    t.Error(tr)
                }
            })
        }
    }
}

func TestLClosure(t *testing.T) {
    //Test the lambda closure function 
    nfal_cases := get_test_nfals()
    //Get a single NFA_l
    nfavar := nfal_cases[0].TestNFAl

    for _, state := range nfavar.States {
        t.Run(fmt.Sprintf("NFA-l Closure/%s", state.Name), func(t *testing.T) {
            //Put state into a set of just itself
            var as_set StateSet
            as_set = EquivSet{state: *new(struct{}),}
            //Calculate lambda closure for that state
            closeset := nfavar.Lclosure(&as_set)
            if !SetEqual(closeset, nfal_cases[0].Closures[state]) {
                t.Error("NFA-l Lambda Closure Error: Expected %v, Got %v", 
                         nfal_cases[0].Closures[state], closeset)
                
            }
        })
    }

}
