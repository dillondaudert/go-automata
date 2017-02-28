//Main executor package
package main

import (
	"./dfasim"
	"fmt"
)

var (
	st1     = dfasim.State{"A", false}
	st2     = dfasim.State{"B", true}
	st3     = dfasim.State{"C", false}
	st4     = dfasim.State{"D", false}
	trAx    = dfasim.TransPair{st1, "x"}
	trAy    = dfasim.TransPair{st1, "y"}
	trBx    = dfasim.TransPair{st2, "x"}
	trBy    = dfasim.TransPair{st2, "y"}
	trCx    = dfasim.TransPair{st3, "x"}
	trCy    = dfasim.TransPair{st3, "y"}
	trDx    = dfasim.TransPair{st4, "x"}
	trDy    = dfasim.TransPair{st4, "y"}
	trtable = map[dfasim.TransPair]dfasim.State{
		trAx: st2,
		trAy: st1,
		trBx: st2,
		trBy: st2,
	}
	trtable2 = map[dfasim.TransPair]dfasim.State{
		trAx: st3,
		trAy: st4,
		trCx: st4,
		trCy: st2,
		trDx: st3,
		trDy: st2,
		trBx: st2,
		trBy: st2,
	}

	sts   = []dfasim.State{st1, st2}
	sts2  = []dfasim.State{st1, st2, st3, st4}
	alpha = "xy"
)

func main() {
    mydfa, _ := dfasim.NewDFA(sts2, st1, alpha, trtable2)
    fmt.Printf("My DFA: %+v\n", mydfa)
    mydfa.Minim()
}
