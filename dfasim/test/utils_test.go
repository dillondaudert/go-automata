package dfasim

//Test functions for the utils.go utilities

import (
	"fmt"
	"testing"
)

type testCase struct {
	st     State
	member bool
}

func TestTraceString(t *testing.T) {
	st1 := State{"name1", false}
	st2 := State{"name2", true}
	st3 := State{"name3", false}
	symb1 := "ab"
	symb2 := "b"
	symb3 := ""
	tr := new(Trace)

	tr.addComputation(st1, symb1)
	tr.addComputation(st2, symb2)
	strout := fmt.Sprintf("%v", tr)
	expected := fmt.Sprintf("name1, \"ab\"\n\t-> name2, \"b\"\n")
	if strout != expected {
		t.Error("Trace print function error, expected:\n", expected,
			"\nreceived:\n", strout)
	}

	tr.addComputation(st3, symb3)
	strout = fmt.Sprintf("%v", tr)
	expected += fmt.Sprintf("\t-> name3, \"\"\n")
	if strout != expected {
		t.Error("Trace print function error, expected:\n", expected,
			"\nreceived:\n", strout)
	}

}

var testpairs_2 = []testCase{
	{st1, true},
	{st2, true},
	{st3, false},
}

func TestEquivSet(t *testing.T) {
	es := make(EquivSet)
	es.AddMember(st1)
	es.AddMember(st2)

	for _, tp := range testpairs_2 {
		exp := tp.member
		out := es.IsMember(tp.st)
		if out != exp {
			t.Error("EquivSet IsMember for %v returned %v when expected %v\n", st1, tp.st, exp)
		}
	}

	exp := []State{st1, st2}
	out := es.Members()
	for i, _ := range out {
		if exp[i] != out[i] {
			t.Error("EquivSet Members returned %v, expected %v\n", out, exp)
		}
	}
}
