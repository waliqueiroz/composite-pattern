package element

type Body struct {
	ContainerElement
}

func NewBody() *Body {
	return &Body{
		ContainerElement: ContainerElement{tagName: "body"},
	}
}
