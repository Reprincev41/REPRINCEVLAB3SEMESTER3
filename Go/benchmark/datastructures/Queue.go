package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type QueueNode struct {
	data int
	next *QueueNode
}

type MyQueue struct {
	frontNode *QueueNode
	rearNode  *QueueNode
}

func NewMyQueue() *MyQueue {
	return &MyQueue{frontNode: nil, rearNode: nil}
}

func (q *MyQueue) Push(value int) {
	newNode := &QueueNode{data: value, next: nil}
	if q.rearNode == nil {
		q.frontNode = newNode
		q.rearNode = newNode
		return
	}
	q.rearNode.next = newNode
	q.rearNode = newNode
}

func (q *MyQueue) Pop() {
	if q.frontNode == nil {
		return
	}
	q.frontNode = q.frontNode.next
	if q.frontNode == nil {
		q.rearNode = nil
	}
}

func (q *MyQueue) Peek() (int, error) {
	if q.frontNode == nil {
		return 0, errors.New("queue empty")
	}
	return q.frontNode.data, nil
}

func (q *MyQueue) Print() {
	fmt.Print("Queue [")
	curr := q.frontNode
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
func (q *MyQueue) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	count := uint64(0)
	curr := q.frontNode
	for curr != nil {
		count++
		curr = curr.next
	}

	if err := binary.Write(file, binary.LittleEndian, count); err != nil {
		return err
	}

	curr = q.frontNode
	for curr != nil {
		if err := binary.Write(file, binary.LittleEndian, int32(curr.data)); err != nil {
			return err
		}
		curr = curr.next
	}
	return nil
}

func (q *MyQueue) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing queue
	for q.frontNode != nil {
		q.Pop()
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
		q.Push(int(value))
	}
	return nil
}

// JSON Serialization
type queueJSON struct {
	Data []int `json:"data"`
}

func (q *MyQueue) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	data := make([]int, 0)
	curr := q.frontNode
	for curr != nil {
		data = append(data, curr.data)
		curr = curr.next
	}

	queueData := queueJSON{Data: data}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(queueData)
}

func (q *MyQueue) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing queue
	for q.frontNode != nil {
		q.Pop()
	}

	var queueData queueJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&queueData); err != nil {
		return err
	}

	for _, value := range queueData.Data {
		q.Push(value)
	}
	return nil
}
