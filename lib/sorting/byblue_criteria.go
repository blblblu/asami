// Generated by: main
// TypeWriter: criteria
// Directive: +gen on byBlue

package sorting

func (p byBlue) Len() int { return len(p) / 4 }
func (p byBlue) Swap(i, j int) {
	for k := 0; k < 4; k++ {
		p[i*4+k], p[j*4+k] = p[j*4+k], p[i*4+k]
	}
}
