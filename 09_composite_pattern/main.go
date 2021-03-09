package main

func main() {
	directory := NewComponent(CompositeNode, "文件夹A")
	directory.AddChild(NewComponent(LeafNode, "1.go"))
	directory.AddChild(NewComponent(LeafNode, "2.go"))
	directory.AddChild(NewComponent(LeafNode, "3.go"))
	directory2 := NewComponent(CompositeNode, "文件夹B")
	directory2.AddChild(NewComponent(LeafNode, "11.go"))
	directory2.AddChild(NewComponent(LeafNode, "22.go"))
	directory2.AddChild(NewComponent(LeafNode, "33.go"))
	directory.AddChild(directory2)

	directory.Print("文件系统")
}
