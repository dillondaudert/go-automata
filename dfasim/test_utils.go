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

type nfalcase struct {
    TestNFAl NFA_l
    Pairs []testpair
    Closures map[State] StateSet
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
        {"", false, true},
    }
    case1 := nfacase{nfa1, pairs1}

    return []nfacase{case1,}
}

func get_test_nfals() []nfalcase {
	var (
		s = State{"S", false}
		l_par = State{"(", false}
        s1 = State{"S1", false}
        a = State{"A", false}
        b = State{"B", false}
        f1 = State{"F1", false}
        r_par = State{")", true}

		sts = []State{s, l_par, s1, a, b, f1, r_par}

		trSl = TransPair{s, "lambda"}
		trLparl = TransPair{l_par, "lambda"}
        trS1a = TransPair{s1, "a"}
        trS1b = TransPair{s1, "b"}
        trA = TransPair{a, "lambda"}
        trB = TransPair{b, "lambda"}
        trF1l = TransPair{f1, "lambda"}

		alpha = "ab"

		Sl = EquivSet{
			l_par: *new(struct{}),
		}
        Lparl = EquivSet{
            s1: *new(struct{}),
            r_par: *new(struct{}),
        }
        S1a = EquivSet{
            a: *new(struct{}),
        }
        S1b = EquivSet{
            b: *new(struct{}),
        }
        ABl = EquivSet{
            f1: *new(struct{}),
        }
        F1l = EquivSet{
            s1: *new(struct{}),
            r_par: *new(struct{}),
        }

        trtable_l = map[TransPair]StateSet{
            trSl: Sl,
            trLparl: Lparl,
            trS1a: S1a,
            trS1b: S1b,
            trA: ABl,
            trB: ABl,
            trF1l: F1l,
        }

	)

    nfa_l := NFA_l{sts, s, alpha, trtable_l}
    pairs := []testpair{
        {"", true, true},
        {"a", true, true},
        {"b", true, true},
        {"aa", true, true},
        {"aba", true, true},
        {"baba", true, true},
        {"bbaa", true, true},
        {"aabbb", true, true},
		{"aabdx", false, false},
		{"dz", false, false},
        {"", true, true},
        {"aaaaa", true, true},
    }

    sclose := make(EquivSet)
    sclose.AddMember(s)
    sclose.AddMember(l_par)
    sclose.AddMember(s1)
    sclose.AddMember(r_par)

    lparclose := make(EquivSet)
    lparclose.AddMember(l_par)
    lparclose.AddMember(s1)
    lparclose.AddMember(r_par)

    s1close := make(EquivSet)
    s1close.AddMember(s1)

    aclose := make(EquivSet)
    aclose.AddMember(a)
    aclose.AddMember(f1)
    aclose.AddMember(s1)
    aclose.AddMember(r_par)

    bclose := make(EquivSet)
    bclose.AddMember(b)
    bclose.AddMember(f1)
    bclose.AddMember(s1)
    bclose.AddMember(r_par)

    f1close := make(EquivSet)
    f1close.AddMember(f1)
    f1close.AddMember(s1)
    f1close.AddMember(r_par)

    rparclose := make(EquivSet)
    rparclose.AddMember(r_par)



    closures := map[State]StateSet {
        s: sclose,
        l_par: lparclose,
        s1: s1close,
        a: aclose,
        b: bclose,
        f1: f1close,
        r_par: rparclose,
    }

    case1 := nfalcase{nfa_l, pairs, closures}

    return []nfalcase{case1,}

}
