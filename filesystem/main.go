package main

import "fmt"

// FileSystemNode: um componente que pode ser um arquivo ou um diretório
type FileSystemNode interface {
	GetSize() int
}

// File: um nó folha no sistema de arquivos
type File struct {
	size int
}

func NewFile(size int) *File {
	return &File{
		size: size,
	}
}

// GetSize: retorna o tamanho do arquivo
func (f *File) GetSize() int {
	return f.size
}

// Directory: um nó composto no sistema de arquivos
type Directory struct {
	children []FileSystemNode
}

func NewDirectory() *Directory {
	return &Directory{}
}

// GetSize: retorna o tamanho total do diretório (incluindo todos os filhos)
func (d *Directory) GetSize() int {
	size := 0
	for _, child := range d.children {
		size += child.GetSize()
	}
	return size
}

// AddChild: adiciona um novo nó filho ao diretório
func (d *Directory) AddChild(child FileSystemNode) {
	d.children = append(d.children, child)
}

func main() {
	file1 := NewFile(10)
	file2 := NewFile(10)
	file3 := NewFile(5)

	directory1 := NewDirectory()
	directory2 := NewDirectory()

	directory2.AddChild(file3)

	directory1.AddChild(file1)
	directory1.AddChild(file2)
	directory1.AddChild(directory2)

	fmt.Println("Tamanho do diretório 2: ", directory2.GetSize())
	fmt.Println("Tamanho do diretório 1: ", directory1.GetSize())
}
