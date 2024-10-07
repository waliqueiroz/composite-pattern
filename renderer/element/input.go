package element

import "fmt"

type Input struct {
	name      string
	inputType string
}

func NewInput(name, inputType string) *Input {
	return &Input{
		name:      name,
		inputType: inputType,
	}
}

func (i *Input) Render() string {
	return fmt.Sprintf(`<input type="%s" name="%s" />`, i.inputType, i.name)
}
