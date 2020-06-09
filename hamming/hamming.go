package hamming

import "errors"

// Distance checks the hammind distance between two or more strings and
//
func Distance(a, b string) (int, error) {
	var count int = 0
	one := a[:]
	two := b[:]

	if len(a) != len(b) {
		return count, errors.New("mismatched genomes")
	}
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			count = count + 1
		}

	}
	return count, nil
}
