package element

type Form struct {
	ContainerElement
}

func NewForm() *Form {
	return &Form{
		ContainerElement: ContainerElement{tagName: "form"},
	}
}
