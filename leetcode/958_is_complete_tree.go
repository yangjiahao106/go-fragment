package main

import "fmt"

func TestIsCompleteTree() {

	coder := Constructor()
	s := coder.serialize(nil)
	fmt.Printf("*%s*", s)
	root := coder.deserializeV3("1,2,3,X,X,X,X")
	fmt.Printf("%+v%+v%+v\n", root, root.Left, root.Right)

	fmt.Println(isCompleteTreeV2(root))

}

// dfs
func isCompleteTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var hasNull = false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			hasNull = true
		}
		if node != nil && hasNull {
			return false
		}
		if node != nil {
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}
	return true
}

func isCompleteTreeV2(root *TreeNode) bool {

	type nodeIdx struct {
		Node *TreeNode
		Idx  int
	}

	queue := make([]nodeIdx, 0)
	queue = append(queue, nodeIdx{Node: root, Idx: 1})

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if node.Node != nil {
			queue = append(queue, nodeIdx{Node: node.Node.Left, Idx: node.Idx * 2})
			queue = append(queue, nodeIdx{Node: node.Node.Right, Idx: node.Idx*2 + 1})
		}
	}

	return queue[len(queue)-1].Idx == len(queue)
}

/*
class Solution(object):
def isCompleteTree(self, root):
nodes = [(root, 1)]
i = 0
while i < len(nodes):
node, v = nodes[i]
i += 1
if node:
nodes.append((node.left, 2*v))
nodes.append((node.right, 2*v+1))

return nodes[-1][1] == len(nodes)

*/
