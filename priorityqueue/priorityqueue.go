package priorityqueue

import "fmt"

type Node struct {
	parent *Node
	childs [2]*Node
	priority int
	value string
}

type PriorityQueue struct {
	root *Node
}

func (queue *PriorityQueue) Add(priority int, value string) {

	node := Node {priority: priority, value: value}

	if queue.root == nil {
		queue.root = &node
		return
	}

	add(queue.root, &node)

	swimHeapify(queue, &node)

}

func (queue PriorityQueue) Peek() (int, string) {
	return queue.root.priority, queue.root.value
}

func (queue PriorityQueue) isEmpty() bool {
	return queue.root == nil
}

func (queue *PriorityQueue) PopMax() (int, string) {

	maxElement := queue.root
	lastElement := getLastElement(queue.root)
	
	if lastElement == nil {
		queue.root = nil
		return maxElement.priority, maxElement.value
	}

	queue.root = lastElement
	maxElement.childs[0].parent = lastElement
	maxElement.childs[1].parent = lastElement

	lastElement.childs = maxElement.childs

	if lastElement.parent.childs[0] == lastElement {
		lastElement.parent.childs[0] = nil
	} else {
		lastElement.parent.childs[1] = nil
	}

	lastElement.parent = nil

	sinkHeapify(queue)

	return maxElement.priority, maxElement.value
}

func getLastElement(node *Node) *Node {

	if node == nil {
		return nil
	}

	if node.childs[0] == nil && node.childs[1] == nil {
		return node
	} 

	rightNode := getLastElement(node.childs[1])
	leftNode := getLastElement(node.childs[0])

	if rightNode == nil || leftNode.priority < rightNode.priority {
		return leftNode
	} else {
		return rightNode
	}
}

func swimHeapify(queue *PriorityQueue, lastAddedNode *Node) {

	if lastAddedNode.parent == nil {
		return
	}

	for lastAddedNode.parent != nil && lastAddedNode.priority > lastAddedNode.parent.priority {
		exchange(lastAddedNode.parent, lastAddedNode)	
	}


	if lastAddedNode.parent == nil {
		queue.root = lastAddedNode
	}
}

func exchange(parent *Node, child *Node) {
	if parent.parent != nil {
		if (parent.parent.childs[0] == parent) {
			parent.parent.childs[0] = child
		} else {
			parent.parent.childs[1] = child
		}
	}	

	child.parent = parent.parent
	parent.parent = child
	lastAddedNodeChilds := child.childs

	if parent.childs[0] == child {
		child.childs[0] = parent
		child.childs[1] = parent.childs[1]
	} else {
		child.childs[0] = parent.childs[0]
		child.childs[1] = parent
	}

	if child.childs[0] != nil {
		child.childs[0].parent = child
	}
	if child.childs[1] != nil {
		child.childs[1].parent = child
	}

	parent.childs = lastAddedNodeChilds

	if parent.childs[0] != nil {
		parent.childs[0].parent = parent 
	}
	if parent.childs[1] != nil {
		parent.childs[1].parent = parent
	}
}

func sinkHeapify(queue *PriorityQueue) {

	currentNode := queue.root

	for (currentNode.childs[0] != nil && currentNode.childs[0].priority > currentNode.priority) || (currentNode.childs[1] != nil && currentNode.childs[1].priority > currentNode.priority) {
		if currentNode.childs[0] != nil && currentNode.childs[0].priority > currentNode.priority {
			if currentNode.parent == nil {
				queue.root = currentNode.childs[0]
			}
			exchange(currentNode, currentNode.childs[0])
		} else {
			if currentNode.parent == nil {
				queue.root = currentNode.childs[1]
			}
			exchange(currentNode, currentNode.childs[1])
		}
	}
}

func add(startNode *Node, newNode *Node) {

	leftNodes := []*Node {startNode}

	for len(leftNodes) > 0 {
		currentNode := leftNodes[0]
		leftNodes = leftNodes[1:]

		if (currentNode.childs[0] == nil) {
			newNode.parent = currentNode
			currentNode.childs[0] = newNode
			break
		} else if (currentNode.childs[1] == nil){
			newNode.parent = currentNode
			currentNode.childs[1] = newNode
			break
		}

		leftNodes = append(leftNodes, currentNode.childs[0], currentNode.childs[1])

	}
}

func PrintQueue(queue PriorityQueue) {
	PrintTree(queue.root, "")
}

func PrintTree(root *Node, indent string) {
    if root == nil {
        return 
    }

    // Print node content
    fmt.Println(indent, "Priority:", root.priority, "Value:", root.value)

    // Prepare indent for child nodes
    newIndent := indent + "  " 

    // Recursively print child nodes (left and then right)
    PrintTree(root.childs[0], newIndent) 
    PrintTree(root.childs[1], newIndent) 
}

