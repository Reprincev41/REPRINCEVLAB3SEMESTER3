package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type MyArray struct {
	data     []int
	capacity int
	size     int
}

func NewMyArray() *MyArray {
	return &MyArray{
		data:     make([]int, 2),
		capacity: 2,
		size:     0,
	}
}

func (a *MyArray) resize(newCapacity int) {
	newData := make([]int, newCapacity)
	copy(newData, a.data[:a.size])
	a.data = newData
	a.capacity = newCapacity
}

func (a *MyArray) AddToEnd(value int) {
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}
	a.data[a.size] = value
	a.size++
}

func (a *MyArray) AddAtIndex(index int, value int) error {
	if index < 0 || index > a.size  {
		return errors.New("index out of bounds")
	}
	if a.size == a.capacity {
		a.resize(a.capacity * 2)
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = value
	a.size++
	return nil
}

func (a *MyArray) GetAtIndex(index int) (int, error) {
	if index < 0 || index >= a.size {
		return 0, errors.New("index out of bounds")
	}
	return a.data[index], nil
}

func (a *MyArray) RemoveAtIndex(index int) error {
	if  index < 0 || index >= a.size{
		return errors.New("index out of bounds")
	}
	for i := index; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.size--
	return nil
}

func (a *MyArray) ReplaceAtIndex(index int, value int) error {
	if index < 0 || index >= a.size  {
		return errors.New("index out of bounds")
	}
	a.data[index] = value
	return nil
}

func (a *MyArray) GetLength() int {
	return a.size
}

func (a *MyArray) Print() {
	fmt.Print("Array [")
	for i := 0; i < a.size; i++ {
		fmt.Print(a.data[i])
		if i < a.size-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

// Binary Serialization
func (a *MyArray) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, uint64(a.size)); err != nil {
		return err
	}
	for i := 0; i < a.size; i++ {
		if err := binary.Write(file, binary.LittleEndian, int32(a.data[i])); err != nil {
			return err
		}
	}
	return nil
}

func (a *MyArray) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	var newSize uint64
	if err := binary.Read(file, binary.LittleEndian, &newSize); err != nil {
		return err
	}

	if int(newSize) > a.capacity {
		a.resize(int(newSize))
	}
	a.size = int(newSize)

	for i := 0; i < a.size; i++ {
		var value int32
		if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
			return err
		}
		a.data[i] = int(value)
	}
	return nil
}

// JSON Serialization
type arrayJSON struct {
	Data []int `json:"data"`
}

func (a *MyArray) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	data := arrayJSON{Data: a.data[:a.size]}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (a *MyArray) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	var data arrayJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	newSize := len(data.Data)     
	if newSize > a.capacity {      
    	a.resize(newSize)         
	}
	a.size = newSize              
	copy(a.data, data.Data)
	return nil
}
