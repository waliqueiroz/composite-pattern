package element

import "fmt"

type Title struct {
	text string
}

func NewTitle(text string) *Title {
	return &Title{text: text}
}

func (t *Title) Render() string {
	return fmt.Sprintf("<title>%s</title>", t.text)
}
