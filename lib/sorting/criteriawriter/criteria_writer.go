package criteriawriter

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewSetWriter())
	if err != nil {
		panic(err)
	}
}

type SetWriter struct{}

func NewSetWriter() *SetWriter {
	return &SetWriter{}
}

func (sw *SetWriter) Name() string {
	return "criteria"
}

func (sw *SetWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	// none
	return result
}

func (sw *SetWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(sw)

	if !found {
		// nothing to be done
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}

	return nil
}
