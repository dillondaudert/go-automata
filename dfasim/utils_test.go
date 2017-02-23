//Test functions for the utils.go utilities
package dfasim

import (
    "fmt"
    "testing"
)

func TestTraceString(t *testing.T) {
    tr := new(trace)
    st1 := state{"name1", false}
    symb1 := "ab"
    st2 := state{"name2", true}
    symb2 := "b"

    tr.addComputation(st1, symb1)
    tr.addComputation(st2, symb2)
    strout := fmt.Sprintf("%v", tr)
    expected := fmt.Sprintf("name1, \"ab\"\n\t-> name2, \"b\"\n")
    if strout != expected {
        t.Error("Trace print function error, expected:\n", expected,
                "\nreceived:\n", strout)
    }

    st3 := state{"name3", false}
    symb3 := ""
    tr.addComputation(st3, symb3)
    strout = fmt.Sprintf("%v", tr)
    expected += fmt.Sprintf("\t-> name3, \"\"\n")
    if strout != expected {
        t.Error("Trace print function error, expected:\n", expected,
                "\nreceived:\n", strout)
    }

}
