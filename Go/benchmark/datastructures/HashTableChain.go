package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type ChainNode struct {
	key   int
	value int
	next  *ChainNode
}

type HashTableChain struct {
	table    []*ChainNode
	size     int
	capacity int
}

func NewHashTableChain(initCap int) *HashTableChain {
	if initCap <= 0 {
		initCap = 8
	}
	return &HashTableChain{
		table:    make([]*ChainNode, initCap),
		size:     0,
		capacity: initCap,
	}
}

func (h *HashTableChain) hash(key int) int {
	return int(math.Abs(float64(key))) % h.capacity
}

func (h *HashTableChain) Insert(key, value int) {
	idx := h.hash(key)
	newNode := &ChainNode{key: key, value: value, next: h.table[idx]}
	h.table[idx] = newNode
	h.size++
}

func (h *HashTableChain) Get(key int) (int, bool) {
	idx := h.hash(key)
	curr := h.table[idx]
	for curr != nil {
		if curr.key == key {
			return curr.value, true
		}
		curr = curr.next
	}
	return 0, false
}

func (h *HashTableChain) Remove(key int) {
	idx := h.hash(key)
	curr := h.table[idx]
	var prev *ChainNode

	for curr != nil {
		if curr.key == key {
			if prev != nil {
				prev.next = curr.next
			} else {
				h.table[idx] = curr.next
			}
			curr = nil
			h.size--
			return
		}
		prev = curr
		curr = curr.next
	}
}

func (h *HashTableChain) Print() {
	fmt.Println("HashTableChain:")
	for i := 0; i < h.capacity; i++ {
		if h.table[i] != nil {
			fmt.Printf("[%d]: ", i)
			curr := h.table[i]
			for curr != nil {
				fmt.Printf("(%d->%d)", curr.key, curr.value)
				if curr.next != nil {
					fmt.Print(" -> ")
				}
				curr = curr.next
			}
			fmt.Println()
		}
	}
}

// Binary Serialization
func (h *HashTableChain) Serialize(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	if err := binary.Write(file, binary.LittleEndian, uint64(h.size)); err != nil {
		return err
	}
	if err := binary.Write(file, binary.LittleEndian, uint64(h.capacity)); err != nil {
		return err
	}

	for i := 0; i < h.capacity; i++ {
		chainSize := uint64(0)
		temp := h.table[i]
		for temp != nil {
			chainSize++
			temp = temp.next
		}

		if err := binary.Write(file, binary.LittleEndian, chainSize); err != nil {
			return err
		}

		curr := h.table[i]
		for curr != nil {
			if err := binary.Write(file, binary.LittleEndian, int32(curr.key)); err != nil {
				return err
			}
			if err := binary.Write(file, binary.LittleEndian, int32(curr.value)); err != nil {
				return err
			}
			curr = curr.next
		}
	}
	return nil
}

func (h *HashTableChain) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing data
	for i := 0; i < h.capacity; i++ {
		node := h.table[i]
		for node != nil {
			node = node.next
		}
	}

	var size, capacity uint64
	if err := binary.Read(file, binary.LittleEndian, &size); err != nil {
		return err
	}
	if err := binary.Read(file, binary.LittleEndian, &capacity); err != nil {
		return err
	}

	h.size = int(size)
	h.capacity = int(capacity)
	h.table = make([]*ChainNode, h.capacity)

	for i := 0; i < h.capacity; i++ {
		h.table[i] = nil
	}

	for i := 0; i < h.capacity; i++ {
		var chainSize uint64
		if err := binary.Read(file, binary.LittleEndian, &chainSize); err != nil {
			return err
		}

		for j := uint64(0); j < chainSize; j++ {
			var key, value int32
			if err := binary.Read(file, binary.LittleEndian, &key); err != nil {
				return err
			}
			if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
				return err
			}
			h.Insert(int(key), int(value))
		}
	}
	return nil
}

// JSON Serialization
type chainEntry struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

type hashTableChainJSON struct {
	Entries []chainEntry `json:"entries"`
}

func (h *HashTableChain) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	entries := make([]chainEntry, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		curr := h.table[i]
		for curr != nil {
			entries = append(entries, chainEntry{Key: curr.key, Value: curr.value})
			curr = curr.next
		}
	}

	data := hashTableChainJSON{Entries: entries}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (h *HashTableChain) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing data
	for i := 0; i < h.capacity; i++ {
		node := h.table[i]
		for node != nil {
			node = node.next
		}
	}

	h.table = make([]*ChainNode, h.capacity)
	h.size = 0

	var data hashTableChainJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for _, entry := range data.Entries {
		h.Insert(entry.Key, entry.Value)
	}
	return nil
}
