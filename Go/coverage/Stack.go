package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type StackNode struct {
	data int
	next *StackNode
}

type MyStack struct {
	topNode *StackNode
}

func NewMyStack() *MyStack {
	return &MyStack{topNode: nil}
}

func (s *MyStack) Push(value int) {
	newNode := &StackNode{data: value, next: s.topNode}
	s.topNode = newNode
}

func (s *MyStack) Pop() {
	if s.topNode != nil {
		s.topNode = s.topNode.next
	}
}

func (s *MyStack) Peek() (int, error) {
	if s.topNode == nil {
		return 0, errors.New("stack empty")
	}
	return s.topNode.data, nil
}

func (s *MyStack) Print() {
	fmt.Print("Stack [")
	curr := s.topNode
	for curr != nil {
		fmt.Print(curr.data)
		if curr.next != nil {
			fmt.Print(" -> ")
		}
		curr = curr.next
	}
	fmt.Println("]")
}

// Binary Serialization
func (s *MyStack) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	count := uint64(0)
	curr := s.topNode
	for curr != nil {
		count++
		curr = curr.next
	}

	if err := binary.Write(file, binary.LittleEndian, count); err != nil {
		return err
	}

	curr = s.topNode
	for curr != nil {
		if err := binary.Write(file, binary.LittleEndian, int32(curr.data)); err != nil {
			return err
		}
		curr = curr.next
	}
	return nil
}

func (s *MyStack) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing stack
	for s.topNode != nil {
		s.Pop()
	}

	var count uint64
	if err := binary.Read(file, binary.LittleEndian, &count); err != nil {
		return err
	}

	for i := uint64(0); i < count; i++ {
		var value int32
		if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
			return err
		}
		s.Push(int(value))
	}
	return nil
}

// JSON Serialization
type stackJSON struct {
	Data []int `json:"data"`
}

func (s *MyStack) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	data := make([]int, 0)
	curr := s.topNode
	for curr != nil {
		data = append(data, curr.data)
		curr = curr.next
	}

	stackData := stackJSON{Data: data}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(stackData)
}

func (s *MyStack) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing stack
	for s.topNode != nil {
		s.Pop()
	}

	var stackData stackJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&stackData); err != nil {
		return err
	}

	for _, value := range stackData.Data {
		s.Push(value)
	}
	return nil
}
