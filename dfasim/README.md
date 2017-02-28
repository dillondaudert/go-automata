##Description of the DFA

####

##Description of the DFA Minimization function

####Table-Filling Algorithm to find equivalent states
*BASIS*: If *p* is an accepting state and *q* is nonaccepting, then the pair {*p, q*} is distinguishable

*INDUCTION*: Let p and q be states such that for some input symbol a, r = delta(p, a) and s = delta(q, a) are a pair of states known to be distinguishable. Then {p, q} is a pair of distinguishable states.

Implementation using 
# dfasim
--
    import "."

Deterministic Finite Automaton *Simulate the behavior of a DFA

Utility structures and their functions for package, no exported names

## Usage

```go
const (
	DFAMissingParams = "Missing parameters"
	DFAInvalidParams = "Invalid parameters"
)
```

#### type DFA

```go
type DFA struct {
}
```

A struct that defines a particular deterministic finite automaton

#### func  NewDFA

```go
func NewDFA(states []state, state0 state, alpha string, tt map[transpair]state) (*DFA, error)
```
Build a valid DFA and return a pointer to it

#### func (*DFA) DeltaFunc

```go
func (dfavar *DFA) DeltaFunc(st state, w string, tr *trace) (final state, ok bool)
```
DFA.DeltaFunc returns the result of the extended transition function for * the
calling DFA. *

#### type DFAError

```go
type DFAError struct {
	Type    string
	Message string
}
```


#### func (*DFAError) Error

```go
func (e *DFAError) Error() string
```
