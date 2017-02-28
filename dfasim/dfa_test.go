//Tests for the dfa.go package
package dfasim

import (
	"fmt"
	"testing"
)

type testpair struct {
	String string
	Accept bool
	Valid  bool
}

var testpairs = []testpair{
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
	st1     = State{"A", false}
	st2     = State{"B", true}
	st3     = State{"C", false}
	st4     = State{"D", false}
	trAx    = TransPair{st1, "x"}
	trAy    = TransPair{st1, "y"}
	trBx    = TransPair{st2, "x"}
	trBy    = TransPair{st2, "y"}
	trCx    = TransPair{st3, "x"}
	trCy    = TransPair{st3, "y"}
	trDx    = TransPair{st4, "x"}
	trDy    = TransPair{st4, "y"}
	trtable = map[TransPair]State{
		trAx: st2,
		trAy: st1,
		trBx: st2,
		trBy: st2,
	}
	trtable2 = map[TransPair]State{
		trAx: st3,
		trAy: st4,
		trCx: st4,
		trCy: st2,
		trDx: st3,
		trDy: st2,
		trBx: st2,
		trBy: st2,
	}

	sts   = []State{st1, st2}
	sts2  = []State{st1, st2, st3, st4}
	alpha = "xy"
)

func TestDFA(t *testing.T) {
	//Run the NewDFA tests
	testcases := []struct {
		states []State
		state0 State
		alpha  string
		tt     map[TransPair]State
	}{
		{sts, st1, alpha, trtable},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("/DFA%d", i), func(t *testing.T) {
			_, err := NewDFA(tc.states, tc.state0, tc.alpha, tc.tt)
			if err != nil {
				t.Fatalf("NewDFA failed unexpectedly\n%v", err)
			}
		})

		for j, pair := range testpairs {
			tc := tc
			pair := pair
			mydfa, _ := NewDFA(tc.states, tc.state0, tc.alpha, tc.tt)
			t.Run(fmt.Sprintf("DF%d", j), func(t *testing.T) {
				t.Parallel()
				tr := new(Trace)
				finalst, ok := mydfa.DeltaFunc(State{}, pair.String, tr)
				if ok != pair.Valid {
					t.Error("Invalid string ", pair.String, " resulted in", ok,
						"when expected: ", pair.Valid)
				} else if finalst.Final != pair.Accept {
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
