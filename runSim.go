//Main executor package
package main

import (
	"./dfasim"
	"fmt"
)

var (
    alpha_ex4 = "10"
    stA = dfasim.State{"A", false}
    stB = dfasim.State{"B", false}
    stC = dfasim.State{"C", true}
    stD = dfasim.State{"D", false}
    stE = dfasim.State{"E", false}
    stF = dfasim.State{"F", false}
    stG = dfasim.State{"G", false}
    stH = dfasim.State{"H", false}
    sts_ex4 = []dfasim.State{stA, stB, stC, stD, stE, stF, stG, stH}
    tt_ex4 = map[dfasim.TransPair]dfasim.State{
        dfasim.TransPair{stA, "0"}: stB,
        dfasim.TransPair{stA, "1"}: stF,
        dfasim.TransPair{stB, "0"}: stG,
        dfasim.TransPair{stB, "1"}: stC,
        dfasim.TransPair{stC, "0"}: stA,
        dfasim.TransPair{stC, "1"}: stC,
        dfasim.TransPair{stD, "0"}: stC,
        dfasim.TransPair{stD, "1"}: stG,
        dfasim.TransPair{stE, "0"}: stH,
        dfasim.TransPair{stE, "1"}: stF,
        dfasim.TransPair{stF, "0"}: stC,
        dfasim.TransPair{stF, "1"}: stG,
        dfasim.TransPair{stG, "0"}: stG,
        dfasim.TransPair{stG, "1"}: stE,
        dfasim.TransPair{stH, "0"}: stG,
        dfasim.TransPair{stH, "1"}: stC,
    }
)

func main() {
	mydfa, _ := dfasim.NewDFA(sts_ex4, stA, alpha_ex4, tt_ex4)
	mydfa.Minim()
    fmt.Printf("The equivalent states are: (%v, %v), (%v, %v), (%v, %v)\n",
                stE, stA, stH, stB, stF, stD)
}
