package permutations

type PermutationObject struct {
	Count int
}

func (p *PermutationObject) permu(a []int) {
	if len(a) > 5 {
		p.Count = p.Count + 1
	} else {
		p.permu(append(a, 0))
		p.permu(append(a, 1))
	}
}
