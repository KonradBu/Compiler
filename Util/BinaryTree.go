package util

type BinaryNode[T Numeric] struct{
	value T
	//Smaller
	left *BinaryNode[T]
	//Bigger
	right *BinaryNode[T]
}

type Numeric interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
    ~float32 | ~float64
}

func makeBinaryTree[T Numeric] (values []T) BinaryNode[T]{
	var head = new(BinaryNode[T])
	head.value = values[0]
	for _, val := range values{
		head.append(val)
	}
	return *head
}

func (head BinaryNode[T]) append (val T){
	if val == head.value{
		return
	}
	if val < head.value{
		if head.left != nil{
			head.left.append(val)
		}else{
			newNode := new(BinaryNode[T])
			newNode.value = val
			head.left = newNode
		}
	}else{
		// val is bigger than head
		if head.left != nil{
			head.right.append(val)
		}else{
			newNode := new(BinaryNode[T])
			newNode.value = val
			head.right = newNode
		}
	}
}

func (head *BinaryNode[T]) appendTree (tree BinaryNode[T]){
	head.append(tree.value)
	if(tree.left != nil){
		head.appendTree(*tree.left)
	}
	if(tree.right != nil){
		head.appendTree(*tree.right)
	}
}

func (head *BinaryNode[T]) delete (val T) BinaryNode[T]{
	if val == head.value{
		if(head.left != nil){
			head.left.appendTree(*head.left)
		}
		if(head.right != nil){
			head.left.appendTree(*head.right)
		}
		return *head.left
	}
	if(val < head.value){
		return head.left.delete(val)
	}else{
		return head.right.delete(val)
	}
}