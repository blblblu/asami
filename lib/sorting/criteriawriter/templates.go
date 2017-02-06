package criteriawriter

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	set,
}

var set = &typewriter.Template{
	Name: "Criteria",
	Text: `
  func (p {{.Name}}) Len() int      { return len(p) / 4 }
  func (p {{.Name}}) Swap(i, j int) {
	for k := 0; k < 4; k++ {
		p[i*4+k], p[j*4+k] = p[j*4+k], p[i*4+k]
	}
}
  `,
	TypeConstraint: typewriter.Constraint{Comparable: false},
}
