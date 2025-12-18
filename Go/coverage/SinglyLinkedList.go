package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type SNode struct {
	data int
	next *SNode
}

type SinglyLinkedList struct {
	head *SNode
	tail *SNode
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (s *SinglyLinkedList) PushFront(value int) {
	newNode := &SNode{data: value, next: s.head}
	s.head = newNode
	if s.tail == nil {
		s.tail = s.head
	}
	s.size++
}

func (s *SinglyLinkedList) PushBack(value int) {
	newNode := &SNode{data: value, next: nil}
	if s.head == nil {
		s.head = newNode
		s.tail = newNode
	} else {
		s.tail.next = newNode
		s.tail = newNode
	}
	s.size++
}

func (s *SinglyLinkedList) InsertAfter(index int, value int) error {
	if index >= s.size {
		return errors.New("index out of bounds")
	}
	curr := s.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	newNode := &SNode{data: value, next: curr.next}
	curr.next = newNode
	if curr == s.tail {
		s.tail = newNode
	}
	s.size++
	return nil
}

func (s *SinglyLinkedList) InsertBefore(index int, value int) error {
	if index == 0 {
		s.PushFront(value)
		return nil
	}
	if index > s.size {
		return errors.New("index out of bounds")
	}
	return s.InsertAfter(index-1, value)
}

func (s *SinglyLinkedList) PopFront() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
	if s.head == nil {
		s.tail = nil
	}
	s.size--
}

func (s *SinglyLinkedList) PopBack() {
	if s.head == nil {
		return
	}
	if s.head == s.tail {
		s.head = nil
		s.tail = nil
		s.size = 0
		return
	}
	curr := s.head
	for curr.next != s.tail {
		curr = curr.next
	}
	s.tail = curr
	s.tail.next = nil
	s.size--
}

func (s *SinglyLinkedList) RemoveAt(index int) error {
	if index >= s.size {
		return errors.New("index out of bounds")
	}
	if index == 0 {
		s.PopFront()
		return nil
	}
	curr := s.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}
	toDel := curr.next
	curr.next = toDel.next
	if toDel == s.tail {
		s.tail = curr
	}
	toDel = nil
	s.size--
	return nil
}

func (s *SinglyLinkedList) RemoveByValue(value int) {
	if s.head == nil {
		return
	}
	if s.head.data == value {
		s.PopFront()
		return
	}
	curr := s.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}
	if curr.next != nil {
		toDel := curr.next
		curr.next = toDel.next
		if toDel == s.tail {
			s.tail = curr
		}
		toDel = nil
		s.size--
	}
}

func (s *SinglyLinkedList) Find(value int) bool {
	curr := s.head
	for curr != nil {
		if curr.data == value {
			return true
		}
		curr = curr.next
	}
	return false
}

func (s *SinglyLinkedList) GetSize() int {
	return s.size
}

func (s *SinglyLinkedList) Print() {
	fmt.Print("SinglyLinkedList [")
	curr := s.head
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
func (s *SinglyLinkedList) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, uint64(s.size)); err != nil {
		return err
	}
	curr := s.head
	for curr != nil {
		if err := binary.Write(file, binary.LittleEndian, int32(curr.data)); err != nil {
			return err
		}
		curr = curr.next
	}
	return nil
}

func (s *SinglyLinkedList) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing list
	s.head = nil
	s.tail = nil
	s.size = 0

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
		s.PushBack(int(value))
	}
	return nil
}

// JSON Serialization
type singlyListJSON struct {
	Data []int `json:"data"`
}

func (s *SinglyLinkedList) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	data := make([]int, 0, s.size)
	curr := s.head
	for curr != nil {
		data = append(data, curr.data)
		curr = curr.next
	}

	listData := singlyListJSON{Data: data}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(listData)
}

func (s *SinglyLinkedList) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing list
	s.head = nil
	s.tail = nil
	s.size = 0

	var listData singlyListJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&listData); err != nil {
		return err
	}

	if len(listData.Data) > 1000000 {
		return errors.New("suspiciously large size in file")
	}

	for _, value := range listData.Data {
		s.PushBack(value)
	}
	return nil
}
