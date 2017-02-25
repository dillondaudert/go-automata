##Description of the DFA

####

##Description of the DFA Minimization function

####Table-Filling Algorithm to find equivalent states
*BASIS*: If *p* is an accepting state and *q* is nonaccepting, then the pair {*p, q*} is distinguishable
*INDUCTION*: Let p and q be states such that for some input symbol a, r = delta(p, a) and s = delta(q, a) are a pair of states known to be distinguishable. Then {p, q} is a pair of distinguishable states.

Implementation using 
