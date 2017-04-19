package regex

import (
    "fmt"
    "testing"
)

func TestParseRegex(t *testing.T) {

    regex_cases := get_test_regex()
    for i, regex_case := range regex_cases {
        for j, pair := range regex_case.Pairs {
            tr := new(Trace)
            t.Run(fmt.Sprintf("Regex/%d:%d", i,j), func(t *testing.T) {
                nfal := ParseRegex(regex_case.Regex, "")
                accept, ok := nfal.DeltaFunc(pair.String, tr)
                if accept != pair.Accept || ok != pair.Valid {
                    t.Error("Expected (", pair.Accept,
                        ",", pair.Valid, "), Got (", accept, ", ", ok,
                        ")\n")
                    t.Error(tr)
                }
            })
        }
    }

}

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

func TestNFAL(t *testing.T) {
    //Test the NFA L delta function
    nfal_cases := get_test_nfals()

    for i, nfal_case := range nfal_cases {
        for j, pair := range nfal_case.Pairs {
            tr := new(Trace)
            t.Run(fmt.Sprintf("NFA-lambda/%d:%d", i,j), func(t *testing.T) {
                accept, ok := nfal_case.TestNFAl.DeltaFunc(pair.String, tr)
                if accept != pair.Accept || ok != pair.Valid {
                    t.Error("Input:", pair.String, "Expected (", pair.Accept,
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
    var as_set EquivSet

    for _, state := range nfavar.States {
        t.Run(fmt.Sprintf("NFA-l Closure/%s", state.Name), func(t *testing.T) {
            //Put state into a set of just itself
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
