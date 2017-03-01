# dfasim
--
    import "."

Deterministic Finite Automaton *Simulate the behavior of a DFA


Error types and functions for the dfasim package

Utility structures and their functions for package

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
	States          []State
	State0          State
	Alpha           string
	TransitionTable map[TransPair]State
}
```

A struct that defines a particular deterministic finite automaton

#### func  NewDFA

```go
func NewDFA(states []State, state0 State, alpha string, tt map[TransPair]State) (*DFA, error)
```
Build a valid DFA and return a pointer to it

#### func (*DFA) DeltaFunc

```go
func (dfavar *DFA) DeltaFunc(st State, w string, tr *Trace) (final State, ok bool)
```
* DFA.DeltaFunc returns the result of the extended transition function for * the
calling DFA.

#### func (*DFA) Minim

```go
func (dfavar *DFA) Minim() (EquivTable, *DFA)
```
* DFA.Minim turns the calling DFA into an equivalent DFA with the minimum *
number of states

#### func (*DFA) String

```go
func (dfavar *DFA) String() string
```
DFA implements the Stringer interface * Return a nicely formatted transition
table

#### type DFAError

```go
type DFAError struct {
	Type    string
	Message string
}
```

package structs ------------------------------------------------------------

#### func (*DFAError) Error

```go
func (e *DFAError) Error() string
```

#### type EquivSet

```go
type EquivSet map[State]struct{}
```


#### func (*EquivSet) AddMember

```go
func (es *EquivSet) AddMember(st State)
```

#### func (EquivSet) IsMember

```go
func (es EquivSet) IsMember(st State) bool
```

#### func (EquivSet) Members

```go
func (es EquivSet) Members() []State
```

#### func (EquivSet) RandomMember

```go
func (es EquivSet) RandomMember() State
```

#### type EquivTable

```go
type EquivTable [][]int
```


#### func  MakeET

```go
func MakeET(numStates int) EquivTable
```
Return a new Equivalence Table with size numStates x numStates

#### func (EquivTable) Distinguished

```go
func (et EquivTable) Distinguished(p int, q int) bool
```
Get whether a pair of states are distinguished in the Equivalence Table

#### func (EquivTable) FormatTable

```go
func (et EquivTable) FormatTable(states []State) string
```

#### func (EquivTable) SetDistinguished

```go
func (et EquivTable) SetDistinguished(p int, q int, round int)
```
Set a pair of states as distinguished in the Equivalence Table

#### type State

```go
type State struct {
	Name  string
	Final bool
}
```


#### type StatePair

```go
type StatePair struct {
	State1 State
	State2 State
}
```


#### type Trace

```go
type Trace []TransPair
```


#### func (Trace) String

```go
func (t Trace) String() string
```

#### type TransPair

```go
type TransPair struct {
	State  State
	Symbol string
}
```


#### func (TransPair) String

```go
func (tr TransPair) String() string
```
