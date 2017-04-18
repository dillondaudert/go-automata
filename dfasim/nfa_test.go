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
