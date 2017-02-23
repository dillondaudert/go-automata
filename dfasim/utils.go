//Utility structures and their functions for package, no exported names
package dfasim

import "fmt"

type state struct {
	name  string
	final bool
}

type transpair struct {
	state  state
	symbol string
}

func (tr transpair) String() string {
    return fmt.Sprintf("%v, \"%v\"", tr.state.name, tr.symbol)
}

type trace []transpair

/* AddComputation adds a State, string pair to the array slice of
 * transitions.
 */
func (t *trace) addComputation(st state, w string) {
	*t = append(*t, transpair{st, w})
}

func (t trace) String() string {
    var transitions string
    for i, pair := range t {
        if i == 0 {
            transitions += fmt.Sprintf("%v\n", pair)
        } else {
            transitions += fmt.Sprintf("\t-> %v\n", pair)
        }
    }
    return fmt.Sprintf("%v", transitions)
}


/* Error Types -------------------------------------------------------------*/

const (
    DFAMissingParams = "Missing parameters"
    DFAInvalidParams = "Invalid parameters"
)

type DFAError struct {
    Type string
    Message string
}

func (e *DFAError) Error() string {
    return fmt.Sprintf("DFA error encountered: %s: %s\n", e.Type, e.Message)
}
