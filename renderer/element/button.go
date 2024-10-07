package element

import "fmt"

type Button struct {
	label string
}

func NewButton(label string) *Button {
	return &Button{
		label: label,
	}
}

func (b *Button) Render() string {
	return fmt.Sprintf(`<button>%s</button>`, b.label)
}
