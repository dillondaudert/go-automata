//Main executor package
package main

import (
    "./dfasim"
	"fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    file, err := os.Open("examples/ex4_8.txt")
    if err != nil {
        panic(1)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    mydfa1 := ParseDFA(scanner)
    file2, err := os.Open("examples/ex4_14.txt")
    if err != nil {
	panic(1)
    }
    defer file2.Close()
    scanner2 := bufio.NewScanner(file2)
    mydfa2 := ParseDFA(scanner2)
    file3, err := os.Open("examples/ex4_15.txt")
    if err != nil {
	panic(1)
    }
    defer file3.Close()
    scanner3 := bufio.NewScanner(file3)
    mydfa3 := ParseDFA(scanner3)

    et1, minim1 := mydfa1.Minim()
    et2, minim2 := mydfa2.Minim()
    et3, minim3 := mydfa3.Minim()

    var tr *dfasim.Trace

//EXAMPLE 4.8
    fmt.Printf("Example 4.8 DFA:\n%v\n", mydfa1)
    tr = new(dfasim.Trace)
    finalst, _ := mydfa1.DeltaFunc(dfasim.State{}, "0101", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n", finalst.Final)

    tr = new(dfasim.Trace)
    finalst, _ = mydfa1.DeltaFunc(dfasim.State{}, "0111", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n\n", finalst.Final)


    fmt.Printf("Example 4.8 Minimized DFA:\n%v\n", minim1)
    fmt.Printf("Distinguishable States Table (number is round, -1 means equivalent):\n%s\n", et1.FormatTable(mydfa1.States))
    tr = new(dfasim.Trace)
    finalst, _ = minim1.DeltaFunc(dfasim.State{}, "0101", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n", finalst.Final)

    tr = new(dfasim.Trace)
    finalst, _ = minim1.DeltaFunc(dfasim.State{}, "0111", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n\n", finalst.Final)


//EXAMPLE 4.14
    fmt.Printf("Example 4.14 DFA:\n%v\n", mydfa2)
    tr = new(dfasim.Trace)
    finalst, _ = mydfa2.DeltaFunc(dfasim.State{}, "0100", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n", finalst.Final)

    tr = new(dfasim.Trace)
    finalst, _ = mydfa2.DeltaFunc(dfasim.State{}, "0111", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n\n", finalst.Final)


    fmt.Printf("Example 4.14 Minimized DFA:\n%v\n", minim2)
    fmt.Printf("Distinguishable States Table (number is round, -1 means equivalent):\n%s\n", et2.FormatTable(mydfa2.States))
    tr = new(dfasim.Trace)
    finalst, _ = minim2.DeltaFunc(dfasim.State{}, "0100", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n", finalst.Final)

    tr = new(dfasim.Trace)
    finalst, _ = minim2.DeltaFunc(dfasim.State{}, "0111", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n\n", finalst.Final)


//EXAMPLE 4.15
    fmt.Printf("Example 4.15 DFA:\n%v\n", mydfa3)
    tr = new(dfasim.Trace)
    finalst, _ = mydfa3.DeltaFunc(dfasim.State{}, "11", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n", finalst.Final)

    tr = new(dfasim.Trace)
    finalst, _ = mydfa3.DeltaFunc(dfasim.State{}, "01111", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n\n", finalst.Final)
    fmt.Printf("Example 4.15 Minimized DFA:\n%v\n", minim3)
    fmt.Printf("Distinguishable States Table (number is round, -1 means equivalent):\n%s\n", et3.FormatTable(mydfa3.States))
    tr = new(dfasim.Trace)
    finalst, _ = minim3.DeltaFunc(dfasim.State{}, "11", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n", finalst.Final)

    tr = new(dfasim.Trace)
    finalst, _ = minim3.DeltaFunc(dfasim.State{}, "01111", tr)
    fmt.Printf("Trace: \n%v", tr)
    fmt.Printf("Accepted: %v\n\n", finalst.Final)

}

func ParseDFA(scanner *bufio.Scanner) (*dfasim.DFA) {
    //Get the list of states
    scanner.Scan()
    state_names := strings.Split(scanner.Text(), ",")

    //Get the start state
    scanner.Scan()
    start_name := scanner.Text()
    var state0 dfasim.State

    //Get the alphabet
    scanner.Scan()
    alpha := scanner.Text()

    //Get the transition function
    scanner.Scan()
    tr_tuples := strings.Split(scanner.Text(), "),(")
    tr_tuples[0] = strings.TrimLeft(tr_tuples[0], "(")
    tr_tuples[len(tr_tuples)-1] = strings.TrimRight(tr_tuples[len(tr_tuples)-1], ")")
    tr_table := make(map[dfasim.TransPair]dfasim.State)


    //Get the final states
    scanner.Scan()
    final_names := strings.Split(scanner.Text(), ",")

    states_map := make(map[dfasim.State]struct{})

    //Add final states to map first
    for _, st := range final_names {
        if st == start_name {
            state0 = dfasim.State{st, true}
        }
        states_map[dfasim.State{st, true}] = *new(struct{})
    }

    //Add the rest of the states
    for _, st := range state_names {
        //if already in set, skip
        if _, ok := states_map[dfasim.State{st, true}]; ok {
            continue
        } else {
            if start_name == st {
                state0 = dfasim.State{st, false}
            }
            states_map[dfasim.State{st, false}] = *new(struct{})
        }
    }

    //Build the new transition function
    for _, tuple := range tr_tuples {
        var in_state, out_state dfasim.State
        triple := strings.Split(tuple, ",")
        //Find state in states_map; may be final or nonfinal
        if _, ok := states_map[dfasim.State{triple[0], false}]; ok {
            in_state = dfasim.State{triple[0], false}
        } else {
            in_state = dfasim.State{triple[0], true}
        }
        if _, ok := states_map[dfasim.State{triple[2], false}]; ok {
            out_state = dfasim.State{triple[2], false}
        } else {
            out_state = dfasim.State{triple[2], true}
        }

        //Build transition
        tr_table[dfasim.TransPair{in_state, triple[1]}] = out_state

    }

    states := make([]dfasim.State, 0, len(states_map))
    for k, _ := range states_map {
        states = append(states, k)
    }
    mydfa, _ := dfasim.NewDFA(states, state0, alpha, tr_table)
    return mydfa

}
