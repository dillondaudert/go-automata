//Test utilities and structures
package dfasim

type testpair struct {
	String string
	Accept bool
	Valid  bool
}

type nfacase struct {
    TestNFA NFA
    Pairs []testpair
}

/* get_test_nfas()
 * Return a slice of (NFA, []testpair) structures.
 * Each element represents a machine to test and the array of
 * test cases against which to test it.
 */
func get_test_nfas() []nfacase {
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
    nfa1 := NFA{sts, st1, alpha, trtable}
    pairs1 := []testpair{
		{"ad", true, true},
		{"abd", true, true},
		{"acd", true, true},
		{"aad", true, true},
		{"aabd", true, true},
		{"aacd", true, true},
		{"aada", false, true},
		{"abda", false, true},
		{"acda", false, true},
		{"cda", false, true},
		{"bda", false, true},
		{"aabdx", false, false},
		{"dz", false, true},
    }
    case1 := nfacase{nfa1, pairs1}

    return []nfacase{case1,}
}
