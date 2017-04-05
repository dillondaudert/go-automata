package dfasim

import (
    "fmt"
    "testing"
)

var (
    st1 = State{"A", false}
    st2 = State{"B", false}
    st3 = State{"C", false}
    st4 = State{"D", true}

    sts = []State{st1, st2, st3, st4}

    trAa = TransPair{st1, "a"}
    trBb = TransPair{st2, "b"}
    trBd = TransPair{st2, "d"}
    trCc = TransPair{st3, "c"}
    trCd = TransPair{st3, "d"}

    alpha = "abcd"
    
    Aa_out = EquivSet{
        st1: *new(struct{}),
        st2: *new(struct{}),
        st3: *new(struct{}),
    }
    Bb_out = EquivSet{
        st2: *new(struct{}),
    }
    Cc_out = EquivSet{
        st3: *new(struct{}),
    }
    d_out = EquivSet{
        st4: *new(struct{}),
    }

    trtable = map[TransPair]StateSet{
        trAa: Aa_out,
        trBb: Bb_out,
        trCc: Cc_out,
        trBd: d_out,
        trCd: d_out,
    }
)

func TestNFA(t *testing.T) {
    //Test the NFA struct
    nfa := NFA{sts, st1, alpha, trtable}
    tr := new(Trace)
    //fmt.Printf("An NFA: %v\n", nfa)
    t.Run(fmt.Sprintf("NFA/"), func(t *testing.T) {
        final, ok := nfa.DeltaFunc("ab", tr)
        if final != false || ok == false {
            t.Error("NFA Delta Func Error")
        }
    })
}
