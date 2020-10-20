package main

import (
	"fmt"
	"learning-go/ui"
)

func main() {
	// r := ui.Rect(150, 100, 0, 0)

	rootView := ui.NewView("root")
	manager := ui.NewViewManager(rootView)

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

	childView2.DeepChildrenLength()
	// fmt.Println(allChildren)

	fmt.Println(rootView.GetChildByIndex(1).GetName(), childView2.GetName())
	fmt.Println(rootView)

	manager.DestroyAndFreeNodes()
	// fmt.Println(morestrings.ReverseRunes("Hello"))
}
