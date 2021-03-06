package ui

type View struct {
	rect     rect
	name     string
	children []*View
}

func NewView(name string) *View {
	r := Rect(150, 100, 0, 0)
	v := View{rect: *r, name: name}
	return &v
}

func (v *View) SetSize(w int, h int) {
	v.rect.setSize(w, h)
}

func (v *View) setPosition(x int, y int) {
	v.rect.setPosition(x, y)
}

func (v *View) AppendChild(childView *View) {
	v.children = append(v.children, childView)
}

func (v *View) GetChildrenLength() int {
	return len(v.children)
}

func (v *View) GetDeepChildrenLength() int {
	count := 0
	nbChildren := len(v.children)
	for i := 0; i < nbChildren; i++ {
		child := v.children[i]
		length := child.GetDeepChildrenLength()
		count += length
	}

	return count + nbChildren
}

func (v *View) Destroy() {
	v = &View{}
}

func (v *View) DestroyAndFree() {
	for i := 0; i < len(v.children); i++ {
		v.children[i].DestroyAndFree()
	}
	v.children = nil
}

func (v *View) GetChildByIndex(index int) *View {
	child := v.children[index]
	return child
}

func (v *View) GetName() string {
	return v.name
}

func (v *View) GetRect() rect {
	return v.rect
}

func (v *View) GetSize() []int {
	return v.rect.getSize()
}
