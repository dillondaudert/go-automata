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
    ParseDFA(scanner)
}

func ParseDFA(scanner *bufio.Scanner) {
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
    tr_tuples := strings.Split(scanner.Text(), "), (")
    tr_tuples[0] = strings.TrimLeft(tr_tuples[0], "(")
    tr_tuples[len(tr_tuples)-1] = strings.TrimRight(tr_tuples[len(tr_tuples)-1], ")")

    //Get the final states
    scanner.Scan()
    final_names := strings.Split(scanner.Text(), ",")

    states_map := make(map[dfasim.State]struct{})

    //Add final states to map first
    for _, st := range final_names {
        if st == start_name {
            state0 = dfasim.State{st, true}
        }
        fmt.Printf("Adding final state %v\n", st)
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
            fmt.Printf("Adding state %v\n", st)
            states_map[dfasim.State{st, false}] = *new(struct{})
        }
    }
    fmt.Printf("%v, %v, %v", state0, alpha, tr_tuples)
}
