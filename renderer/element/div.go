package element

type Div struct {
	ContainerElement
}

func NewDiv() *Div {
	return &Div{
		ContainerElement: ContainerElement{tagName: "div"},
	}
}
