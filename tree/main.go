package main

import (
	"fmt"
	"math"
	"sort"
)

var OrgNodeData = make(map[int]*OrgNode)

// 组织节点结构体
type OrgNode struct {
	ID        int
	Name      string
	ParentID  int
	Latitude  float64
	Longitude float64
	Children  []*OrgNode
}

// 更新节点
func UpdateNode(node *OrgNode, name string) {
	node.Name = name
}

// 删除节点
func DeleteNode(node *OrgNode) {
	parent := FindParentNode(node)
	if parent == nil {
		return
	}

	children := parent.Children
	for i, child := range children {
		if child == node {
			parent.Children = append(children[:i], children[i+1:]...)
			break
		}
	}
}

// 更换上级绑定
func ChangeParent(node *OrgNode, newParent *OrgNode) {
	parent := FindParentNode(node)
	if parent != nil {
		children := parent.Children
		for i, child := range children {
			if child == node {
				parent.Children = append(children[:i], children[i+1:]...)
				break
			}
		}
	}

	if newParent != nil {
		newParent.Children = append(newParent.Children, node)
		node.ParentID = newParent.ID
	}
}

// 查找上一级节点
func FindParentNode(node *OrgNode) *OrgNode {
	return FindNodeByID(node.ParentID)
}

// 查找下一级节点
func FindChildNodes(node *OrgNode) []*OrgNode {
	return node.Children
}

// 查找某个节点的所有上级节点
func FindAllAncestorNodes(node *OrgNode) []*OrgNode {
	ancestors := []*OrgNode{}
	parent := FindParentNode(node)
	for parent != nil {
		ancestors = append(ancestors, parent)
		parent = FindParentNode(parent)
	}
	return ancestors
}

// 查找某个节点的所有下级节点
func FindAllDescendantNodes(node *OrgNode) []*OrgNode {
	descendants := []*OrgNode{}
	collectDescendants(node, &descendants)
	return descendants
}

// 递归收集节点的所有下级节点
func collectDescendants(node *OrgNode, descendants *[]*OrgNode) {
	for _, child := range node.Children {
		*descendants = append(*descendants, child)
		collectDescendants(child, descendants)
	}
}

// 根据经纬度计算两点之间的距离
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	radius := 6371.0 // 地球半径（单位：公里）

	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	dlon := lon2Rad - lon1Rad
	dlat := lat2Rad - lat1Rad

	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := radius * c

	return distance
}

// 根据经纬度对节点进行距离排序
func SortNodesByDistance(lat, lon float64, nodes []*OrgNode) {
	sort.Slice(nodes, func(i, j int) bool {
		distance1 := CalculateDistance(lat, lon, nodes[i].Latitude, nodes[i].Longitude)
		distance2 := CalculateDistance(lat, lon, nodes[j].Latitude, nodes[j].Longitude)
		return distance1 < distance2
	})
}

// 根据节点ID查找节点
func FindNodeByID(id int) *OrgNode {
	node, ok := OrgNodeData[id]
	if ok {
		return node
	}
	return nil
}

// 存储所有数据
func OrgNodeDataAllData(nodes ...*OrgNode) {
	for _, node := range nodes {
		if node != nil {
			OrgNodeData[node.ID] = node
		}
	}

}

func main() {
	// 创建组织架构树
	root := &OrgNode{
		ID:        1,
		Name:      "Root",
		Latitude:  0.0,
		Longitude: 0.0,
	}

	node1 := &OrgNode{
		ID:        2,
		Name:      "Node 1",
		ParentID:  1,
		Latitude:  10.0,
		Longitude: 10.0,
	}
	node2 := &OrgNode{
		ID:        3,
		Name:      "Node 2",
		ParentID:  1,
		Latitude:  20.0,
		Longitude: 20.0,
	}
	node3 := &OrgNode{
		ID:        4,
		Name:      "Node 3",
		ParentID:  2,
		Latitude:  15.0,
		Longitude: 15.0,
	}
	root.Children = append(root.Children, node1, node2)
	node1.Children = append(node1.Children, node3)
	OrgNodeDataAllData(root, node1, node2, node3)

	//// 更新节点
	//UpdateNode(node3, "Updated Node")
	//fmt.Println("Updated Node Name:", node3.Name)
	//
	//// 删除节点
	//DeleteNode(node2)
	//fmt.Println("Deleted Node:", FindNodeByID(3))

	//更换上级绑定
	newParent := FindNodeByID(1)
	ChangeParent(node3, newParent)
	fmt.Println("New Parent of Node 3:", FindParentNode(node3).ID)

	// 查找上一级节点
	parent := FindParentNode(node3)
	fmt.Println("Parent of Node 3:", parent)

	// 查找下一级节点
	children := FindChildNodes(root)
	fmt.Println("Children of Root:")
	for _, child := range children {
		fmt.Printf("ID: %d, Name: %s\n", child.ID, child.Name)
	}

	// 查找某个节点的所有上级节点
	ancestors := FindAllAncestorNodes(node3)
	fmt.Println("Ancestors of Node 3:")
	for _, ancestor := range ancestors {
		fmt.Printf("ID: %d, Name: %s\n", ancestor.ID, ancestor.Name)
	}

	// 查找某个节点的所有下级节点
	descendants := FindAllDescendantNodes(root)
	fmt.Println("Descendants of Root:")
	for _, descendant := range descendants {
		fmt.Printf("ID: %d, Name: %s\n", descendant.ID, descendant.Name)
	}

	// 根据经纬度实现距离排序
	SortNodesByDistance(0.0, 0.0, root.Children)
	fmt.Println("Nodes sorted by distance:")
	for _, node := range root.Children {
		fmt.Printf("ID: %d, Name: %s, Distance: %f\n", node.ID, node.Name, CalculateDistance(0.0, 0.0, node.Latitude, node.Longitude))
	}
}
