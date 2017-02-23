//Tests for the dfa.go package
package dfasim

import "testing"

type testpair struct {
    String string
    Accept bool
}

var validpairs = []testpair {
    {"yyx", true},
    {"yyyyyyyx", true},
    {"x", true},
    {"xyxyxy", true},
    {"yyyyy", false},
    {"xxxxy", true},
    {"", false},
}

func TestDeltaFunc(t *testing.T) {
    st1 := State{"A", false}
    st2 := State{"B", true}
    trAx := TransPair{st1, "x"}
    trAy := TransPair{st1, "y"}
    trBx := TransPair{st2, "x"}
    trBy := TransPair{st2, "y"}
    trTable := map[TransPair]State {
        trAx: st2,
        trAy: st1,
        trBx: st2,
        trBy: st2,
    }

    sts := []State{st1, st2}
    alpha := "xy"
    mydfa := DFA{sts, st1, alpha, trTable}

    for _, pair := range validpairs {
        tr := new(Trace)
        finalst, ok := mydfa.DeltaFunc(State{}, pair.String, tr)
        if !ok {
            t.Error("Valid string ", pair.String, "rejected.")
        } else if finalst.Final != pair.Accept {
            t.Error("String ", pair.String, "resulted in ", finalst,
                    "when expected: ", pair.Accept)
            t.Error(tr)
        }
    }
}
