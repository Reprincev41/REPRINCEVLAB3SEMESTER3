package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type DNode struct {
	data int
	next *DNode
	prev *DNode
}

type DoublyLinkedList struct {
	head *DNode
	tail *DNode
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (d *DoublyLinkedList) GetSize() int {
	return d.size
}

func (d *DoublyLinkedList) PushFront(value int) {
	newNode := &DNode{data: value, next: nil, prev: nil}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.size++
}

func (d *DoublyLinkedList) PushBack(value int) {
	newNode := &DNode{data: value, next: nil, prev: nil}
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.size++
}

func (d *DoublyLinkedList) InsertAfter(index int, value int) error {
	if index >= d.size {
		return errors.New("index out of bounds")
	}
	if index == d.size-1 {
		d.PushBack(value)
		return nil
	}
	curr := d.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	newNode := &DNode{data: value, next: curr.next, prev: curr}
	curr.next.prev = newNode
	curr.next = newNode
	d.size++
	return nil
}

func (d *DoublyLinkedList) InsertBefore(index int, value int) error {
	if index > d.size {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		d.PushFront(value)
		return nil
	}
	return d.InsertAfter(index-1, value)
}

func (d *DoublyLinkedList) PopFront() {
	if d.head == nil {
		return
	}
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.size--
}

func (d *DoublyLinkedList) PopBack() {
	if d.tail == nil {
		return
	}
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.size--
}

func (d *DoublyLinkedList) RemoveAt(index int) error {
	if index >= d.size {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		d.PopFront()
		return nil
	}
	if index == d.size-1 {
		d.PopBack()
		return nil
	}
	curr := d.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	curr = nil
	d.size--
	return nil
}

func (d *DoublyLinkedList) RemoveByValue(value int) {
	curr := d.head
	for curr != nil {
		if curr.data == value {
			if curr == d.head {
				d.PopFront()
			} else if curr == d.tail {
				d.PopBack()
			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
				curr = nil
				d.size--
			}
			return
		}
		curr = curr.next
	}
}

func (d *DoublyLinkedList) Find(value int) bool {
	curr := d.head
	for curr != nil {
		if curr.data == value {
			return true
		}
		curr = curr.next
	}
	return false
}

func (d *DoublyLinkedList) Print() {
	fmt.Print("DoublyLinkedList [")
	curr := d.head
	for curr != nil {
		fmt.Print(curr.data)
		if curr.next != nil {
			fmt.Print(" <-> ")
		}
		curr = curr.next
	}
	fmt.Println("]")
}

// Binary Serialization
func (d *DoublyLinkedList) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, uint64(d.size)); err != nil {
		return err
	}
	curr := d.head
	for curr != nil {
		if err := binary.Write(file, binary.LittleEndian, int32(curr.data)); err != nil {
			return err
		}
		curr = curr.next
	}
	return nil
}

func (d *DoublyLinkedList) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing list
	d.head = nil
	d.tail = nil
	d.size = 0

	var fileSize uint64
	if err := binary.Read(file, binary.LittleEndian, &fileSize); err != nil {
		return err
	}

	if fileSize > 1000000 {
		return errors.New("suspiciously large size in file")
	}

	for i := uint64(0); i < fileSize; i++ {
		var value int32
		if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
			return err
		}
		d.PushBack(int(value))
	}
	return nil
}

// JSON Serialization
type doublyListJSON struct {
	Data []int `json:"data"`
}

func (d *DoublyLinkedList) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	data := make([]int, 0, d.size)
	curr := d.head
	for curr != nil {
		data = append(data, curr.data)
		curr = curr.next
	}

	listData := doublyListJSON{Data: data}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(listData)
}

func (d *DoublyLinkedList) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing list
	d.head = nil
	d.tail = nil
	d.size = 0

	var listData doublyListJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&listData); err != nil {
		return err
	}

	if len(listData.Data) > 1000000 {
		return errors.New("suspiciously large size in file")
	}

	for _, value := range listData.Data {
		d.PushBack(value)
	}
	return nil
}
