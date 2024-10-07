package element

type Head struct {
	ContainerElement
}

func NewHead() *Head {
	return &Head{
		ContainerElement: ContainerElement{tagName: "head"},
	}
}
