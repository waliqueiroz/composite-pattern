package main

import (
	"fmt"

	element "example.com/renderer/element"
)

func main() {
	// Criando a estrutura da página
	html := element.NewHTML()
	head := element.NewHead()
	body := element.NewBody()

	// Adiciona título no <head>
	head.AddElement(element.NewTitle("Minha Página HTML"))

	// Cria um formulário
	form := element.NewForm()
	form.AddElement(element.NewLabel("Insira o nome de usuário", "username"))
	form.AddElement(element.NewInput("username", "text"))
	form.AddElement(element.NewLabel("Insira a senha", "password"))
	form.AddElement(element.NewInput("password", "password"))
	form.AddElement(element.NewButton("Entrar"))

	// Cria uma div dentro do body e adiciona ao form
	div := element.NewDiv()
	div.AddElement(element.NewLabel("Cadastre-se", "email"))
	div.AddElement(element.NewInput("email", "email"))
	div.AddElement(element.NewButton("Cadastrar"))
	form.AddElement(div)

	// Adiciona o form no <body>
	body.AddElement(form)

	// Monta a página HTML
	html.AddElement(head)
	html.AddElement(body)

	// Renderiza a página completa
	fmt.Println(html.Render())
}
