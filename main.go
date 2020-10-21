package main

import (
	"fmt"
	"learning-go/ui"
)

func App() *ui.View {
	rootView := ui.NewView("root")
	rootView.SetSize(800, 600)

	fmt.Println(rootView.GetSize())

	childView := ui.NewView("child 1")
	childView2 := ui.NewView("child 2")

	childView3 := ui.NewView("subchild 1")
	childView4 := ui.NewView("subchild 2")
	childView5 := ui.NewView("subsubchild 1")

	childView.AppendChild(childView5)
	childView2.AppendChild(childView3)
	childView2.AppendChild(childView4)
	rootView.AppendChild(childView)
	rootView.AppendChild(childView2)

	fmt.Println(rootView.GetDeepChildrenLength())

	return rootView
}

func main() {
	manager := ui.NewViewManager(App())
	manager.DestroyAndFreeNodes()
}
