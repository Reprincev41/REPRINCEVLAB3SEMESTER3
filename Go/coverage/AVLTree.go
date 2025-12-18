package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type AVLNode struct {
	key    int
	left   *AVLNode
	right  *AVLNode
	height int
}

type AVLTree struct {
	root *AVLNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{root: nil}
}

func (t *AVLTree) height(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return n.height
}

func (t *AVLTree) max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t *AVLTree) rightRotate(y *AVLNode) *AVLNode {
	x := y.left
	T2 := x.right

	x.right = y
	y.left = T2

	y.height = t.max(t.height(y.left), t.height(y.right)) + 1
	x.height = t.max(t.height(x.left), t.height(x.right)) + 1

	return x
}

func (t *AVLTree) leftRotate(x *AVLNode) *AVLNode {
	y := x.right
	T2 := y.left

	y.left = x
	x.right = T2

	x.height = t.max(t.height(x.left), t.height(x.right)) + 1
	y.height = t.max(t.height(y.left), t.height(y.right)) + 1

	return y
}

func (t *AVLTree) getBalance(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return t.height(n.left) - t.height(n.right)
}

func (t *AVLTree) insertNode(node *AVLNode, key int) *AVLNode {
	if node == nil {
		return &AVLNode{key: key, left: nil, right: nil, height: 1}
	}

	if key < node.key {
		node.left = t.insertNode(node.left, key)
	} else if key > node.key {
		node.right = t.insertNode(node.right, key)
	} else {
		return node
	}

	node.height = 1 + t.max(t.height(node.left), t.height(node.right))
	balance := t.getBalance(node)

	// LL Case
	if balance > 1 && key < node.left.key {
		return t.rightRotate(node)
	}

	// RR Case
	if balance < -1 && key > node.right.key {
		return t.leftRotate(node)
	}

	// LR Case
	if balance > 1 && key > node.left.key {
		node.left = t.leftRotate(node.left)
		return t.rightRotate(node)
	}

	// RL Case
	if balance < -1 && key < node.right.key {
		node.right = t.rightRotate(node.right)
		return t.leftRotate(node)
	}

	return node
}

func (t *AVLTree) minValueNode(node *AVLNode) *AVLNode {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (t *AVLTree) deleteNode(root *AVLNode, key int) *AVLNode {
	if root == nil {
		return root
	}

	if key < root.key {
		root.left = t.deleteNode(root.left, key)
	} else if key > root.key {
		root.right = t.deleteNode(root.right, key)
	} else {
		if root.left == nil || root.right == nil {
			var temp *AVLNode
			if root.left != nil {
				temp = root.left
			} else {
				temp = root.right
			}

			if temp == nil {
				temp = root
				root = nil
			} else {
				*root = *temp
			}
			temp = nil
		} else {
			temp := t.minValueNode(root.right)
			root.key = temp.key
			root.right = t.deleteNode(root.right, temp.key)
		}
	}

	if root == nil {
		return root
	}

	root.height = 1 + t.max(t.height(root.left), t.height(root.right))
	balance := t.getBalance(root)

	if balance > 1 && t.getBalance(root.left) >= 0 {
		return t.rightRotate(root)
	}

	if balance > 1 && t.getBalance(root.left) < 0 {
		root.left = t.leftRotate(root.left)
		return t.rightRotate(root)
	}

	if balance < -1 && t.getBalance(root.right) <= 0 {
		return t.leftRotate(root)
	}

	if balance < -1 && t.getBalance(root.right) > 0 {
		root.right = t.rightRotate(root.right)
		return t.leftRotate(root)
	}

	return root
}

func (t *AVLTree) inOrder(root *AVLNode) {
	if root != nil {
		t.inOrder(root.left)
		fmt.Print(root.key, " ")
		t.inOrder(root.right)
	}
}

func (t *AVLTree) destroyTree(node *AVLNode) {
	if node != nil {
		t.destroyTree(node.left)
		t.destroyTree(node.right)
		node = nil
	}
}

func (t *AVLTree) Insert(key int) {
	t.root = t.insertNode(t.root, key)
}

func (t *AVLTree) Remove(key int) {
	t.root = t.deleteNode(t.root, key)
}

func (t *AVLTree) Find(key int) bool {
	curr := t.root
	for curr != nil {
		if key == curr.key {
			return true
		}
		if key < curr.key {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (t *AVLTree) Print() {
	fmt.Print("AVLTree (In-order): ")
	t.inOrder(t.root)
	fmt.Println()
}

// Binary Serialization
func (t *AVLTree) serializeHelper(node *AVLNode, file *os.File) error {
	if node == nil {
		nullMarker := int32(-1)
		return binary.Write(file, binary.LittleEndian, nullMarker)
	}

	if err := binary.Write(file, binary.LittleEndian, int32(node.key)); err != nil {
		return err
	}
	if err := t.serializeHelper(node.left, file); err != nil {
		return err
	}
	return t.serializeHelper(node.right, file)
}

func (t *AVLTree) deserializeHelper(file *os.File) (*AVLNode, error) {
	var key int32
	if err := binary.Read(file, binary.LittleEndian, &key); err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, err
	}

	if key == -1 {
		return nil, nil
	}

	node := &AVLNode{key: int(key), left: nil, right: nil, height: 1}

	left, err := t.deserializeHelper(file)
	if err != nil && err != io.EOF {
		return nil, err
	}
	node.left = left

	right, err := t.deserializeHelper(file)
	if err != nil && err != io.EOF {
		return nil, err
	}
	node.right = right

	node.height = 1 + t.max(t.height(node.left), t.height(node.right))
	return node, nil
}

func (t *AVLTree) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	return t.serializeHelper(t.root, file)
}

func (t *AVLTree) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	t.destroyTree(t.root)
	root, err := t.deserializeHelper(file)
	if err != nil && err != io.EOF {
		return err
	}
	t.root = root
	return nil
}

// JSON Serialization
type avlTreeJSON struct {
	Keys []int `json:"keys"`
}

func (t *AVLTree) collectKeys(node *AVLNode, keys *[]int) {
	if node != nil {
		t.collectKeys(node.left, keys)
		*keys = append(*keys, node.key)
		t.collectKeys(node.right, keys)
	}
}

func (t *AVLTree) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	keys := make([]int, 0)
	t.collectKeys(t.root, &keys)

	treeData := avlTreeJSON{Keys: keys}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(treeData)
}

func (t *AVLTree) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	t.destroyTree(t.root)
	t.root = nil

	var treeData avlTreeJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&treeData); err != nil {
		return err
	}

	for _, key := range treeData.Keys {
		t.Insert(key)
	}
	return nil
}
