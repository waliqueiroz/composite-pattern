package element

import (
	"fmt"
	"strings"
)

type ContainerElement struct {
	tagName string
	content []HTMLElement
}

// AddElement adiciona um elemento filho a um container.
func (e *ContainerElement) AddElement(element HTMLElement) {
	e.content = append(e.content, element)
}

// Render renderiza o elemento container e seus elementos filhos.
func (e *ContainerElement) Render() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("<%s>", e.tagName))
	for _, element := range e.content {
		builder.WriteString("\n")
		builder.WriteString(element.Render())
	}
	builder.WriteString(fmt.Sprintf("\n</%s>", e.tagName))
	return builder.String()
}
