package main

import "fmt"

func main() {

	file1 := File{name: "some1.txt", size: 20}
	file2 := File{name: "other.png", size: 20}

	folder1 := Folder{name: "mydocs"}
	folder1.AddItem(&file1)
	folder1.AddItem(&file2)

	file1.PrintStructure(" ")
	file1.GetSize()

	folder1.PrintStructure(" ")
folder1.Delete()
}

// filesytstem

type FileSystem interface {
	GetSize() int
	PrintStructure(indent string)
	Delete()
}

// leaf node

type File struct {
	name string
	size int
}

func (f *File) GetSize() int {
	return f.size
}
func (f *File) PrintStructure(indent string) {

	fmt.Println(indent,  f.name, f.size)
}

func (f *File) Delete() {
	fmt.Println("deleting file", f.name)
}

// composite node  Foloder

type Folder struct {
	name     string
	children []FileSystem
}

func (f *Folder) GetSize() int {
	totalSize := 0
	for _, file := range f.children {
		totalSize += file.GetSize()
	}
	return totalSize
}

func (f *Folder) Delete() {
	fmt.Println("deleting folder", f.name)
}

func (f *Folder) PrintStructure(indent string) {
	fmt.Println(indent + " +" + f.name + "/")
	for _, item := range f.children {
		item.PrintStructure(indent + "    ")
	}
}

func (f *Folder) AddItem(item FileSystem) {
	f.children = append(f.children, item)
}

func (f *Folder) RemoveItem(item FileSystem) {

	for i, child := range f.children {
		if child == item {
			f.children = append(f.children[:i], f.children[i+1:]...)
			return
		}
	}
}
