//Tests for the dfa.go package
package dfasim

import (
    "fmt"
    "testing"
)

type testpair struct {
    String string
    Accept bool
    Valid bool
}

var testpairs = []testpair {
    {"yyx", true, true},
    {"yyyyyyyx", true, true},
    {"x", true, true},
    {"xyxyxy", true, true},
    {"yyyyy", false, true},
    {"xxxxy", true, true},
    {"", false, true},
    {"z", false, false},
    {"xyxyz", false, false},
    {"zxxxx", false, false},
    {"zyyyy", false, false},
    {"xxxxz", false, false},
    {"yyyyz", false, false},
    {"aaaaa", false, false},
}

var (
    st1 = state{"A", false}
    st2 = state{"B", true}
    trAx = transpair{st1, "x"}
    trAy = transpair{st1, "y"}
    trBx = transpair{st2, "x"}
    trBy = transpair{st2, "y"}
    trtable = map[transpair]state {
        trAx: st2,
        trAy: st1,
        trBx: st2,
        trBy: st2,
    }

    sts = []state{st1, st2}
    alpha = "xy"
)

func TestDFA(t *testing.T) {
    //Run the NewDFA tests
    testcases := []struct {
        states []state
        state0 state
        alpha string
        tt map[transpair]state
    }{
        {sts, st1, alpha, trtable},
    }
    for i, tc := range testcases {
        t.Run(fmt.Sprintf("/DFA%d", i), func(t *testing.T){
            _, err := NewDFA(tc.states, tc.state0, tc.alpha, tc.tt)
            if err != nil {
                t.Fatalf("NewDFA failed unexpectedly\n%v", err)
            }
        })

        for j, pair := range testpairs {
            tc := tc
            pair := pair
            mydfa, _ := NewDFA(tc.states, tc.state0, tc.alpha, tc.tt)
            t.Run(fmt.Sprintf("DF%d", j), func(t *testing.T){
                t.Parallel()
                tr := new(trace)
                finalst, ok := mydfa.DeltaFunc(state{}, pair.String, tr)
                if ok != pair.Valid {
                    t.Error("Invalid string ", pair.String, " resulted in", ok,
                            "when expected: ", pair.Valid)
                } else if finalst.final != pair.Accept {
                    t.Error("String ", pair.String, "resulted in ", finalst,
                            "when expected: ", pair.Accept)
                    t.Error(tr)
                }
                })
        }
    }
}

func parallelDeltaTest(t *testing.T) {
}
