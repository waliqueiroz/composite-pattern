package element

type HTML struct {
	ContainerElement
}

func NewHTML() *HTML {
	return &HTML{
		ContainerElement: ContainerElement{tagName: "html"},
	}
}
