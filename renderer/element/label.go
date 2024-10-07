package element

import "fmt"

type Label struct {
	labelFor string
	text     string
}

func NewLabel(text, labelFor string) *Label {
	return &Label{
		text:     text,
		labelFor: labelFor,
	}
}

func (l *Label) Render() string {
	return fmt.Sprintf(`<label for="%s">%s</label>`, l.labelFor, l.text)
}
