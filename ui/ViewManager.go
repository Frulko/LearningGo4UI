package ui

type ViewManager struct {
	nodes []*View
}

func NewViewManager(rootNode *View) *ViewManager {
	v := ViewManager{}
	v.nodes = append(v.nodes, rootNode)
	return &v
}

func (v *ViewManager) GetNodes() []*View {
	return v.nodes
}

func (v *ViewManager) DestroyAndFreeNodes() {
	for i := 0; i < len(v.nodes); i++ {
		v.nodes[i].DestroyAndFree()
	}
	v.nodes = nil
}
