package datastructures

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type HashEntry struct {
	key        int
	value      int
	isOccupied bool
	isDeleted  bool
}

type HashTableOpen struct {
	table    []HashEntry
	size     int
	capacity int
}

func NewHashTableOpen(initCap int) *HashTableOpen {
	if initCap <= 0 {
		initCap = 8
	}
	table := make([]HashEntry, initCap)
	for i := 0; i < initCap; i++ {
		table[i] = HashEntry{isOccupied: false, isDeleted: false}
	}
	return &HashTableOpen{
		table:    table,
		size:     0,
		capacity: initCap,
	}
}

func (h *HashTableOpen) hash(key int) int {
	return int(math.Abs(float64(key))) % h.capacity
}

func (h *HashTableOpen) resize() {
	oldCapacity := h.capacity
	oldTable := h.table

	h.capacity *= 2
	h.table = make([]HashEntry, h.capacity)
	for i := 0; i < h.capacity; i++ {
		h.table[i] = HashEntry{isOccupied: false, isDeleted: false}
	}

	h.size = 0
	for i := 0; i < oldCapacity; i++ {
		if oldTable[i].isOccupied && !oldTable[i].isDeleted {
			h.Insert(oldTable[i].key, oldTable[i].value)
		}
	}
}

func (h *HashTableOpen) Insert(key, value int) {
	if float64(h.size) >= float64(h.capacity)*0.7 {
		h.resize()
	}

	idx := h.hash(key)
	startIdx := idx

	for h.table[idx].isOccupied && !h.table[idx].isDeleted && h.table[idx].key != key {
		idx = (idx + 1) % h.capacity
		if idx == startIdx {
			return
		}
	}

	h.table[idx].key = key
	h.table[idx].value = value
	h.table[idx].isOccupied = true
	h.table[idx].isDeleted = false
	h.size++
}

func (h *HashTableOpen) Get(key int) (int, bool) {
	idx := h.hash(key)
	startIdx := idx

	for h.table[idx].isOccupied {
		if !h.table[idx].isDeleted && h.table[idx].key == key {
			return h.table[idx].value, true
		}
		idx = (idx + 1) % h.capacity
		if idx == startIdx {
			break
		}
	}
	return 0, false
}

func (h *HashTableOpen) Remove(key int) {
	idx := h.hash(key)
	startIdx := idx

	for h.table[idx].isOccupied {
		if !h.table[idx].isDeleted && h.table[idx].key == key {
			h.table[idx].isDeleted = true
			h.size--
			return
		}
		idx = (idx + 1) % h.capacity
		if idx == startIdx {
			break
		}
	}
}

func (h *HashTableOpen) Print() {
	fmt.Println("HashTableOpen:")
	for i := 0; i < h.capacity; i++ {
		if h.table[i].isOccupied && !h.table[i].isDeleted {
			fmt.Printf("[%d]: %d -> %d\n", i, h.table[i].key, h.table[i].value)
		}
	}
}

// Binary Serialization
func (h *HashTableOpen) Serialize(filename string) error {
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
		if err := binary.Write(file, binary.LittleEndian, int32(h.table[i].key)); err != nil {
			return err
		}
		if err := binary.Write(file, binary.LittleEndian, int32(h.table[i].value)); err != nil {
			return err
		}
		if err := binary.Write(file, binary.LittleEndian, h.table[i].isOccupied); err != nil {
			return err
		}
		if err := binary.Write(file, binary.LittleEndian, h.table[i].isDeleted); err != nil {
			return err
		}
	}
	return nil
}

func (h *HashTableOpen) Deserialize(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	var size, capacity uint64
	if err := binary.Read(file, binary.LittleEndian, &size); err != nil {
		return err
	}
	if err := binary.Read(file, binary.LittleEndian, &capacity); err != nil {
		return err
	}

	h.size = int(size)
	h.capacity = int(capacity)
	h.table = make([]HashEntry, h.capacity)

	for i := 0; i < h.capacity; i++ {
		var key, value int32
		if err := binary.Read(file, binary.LittleEndian, &key); err != nil {
			return err
		}
		if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
			return err
		}
		if err := binary.Read(file, binary.LittleEndian, &h.table[i].isOccupied); err != nil {
			return err
		}
		if err := binary.Read(file, binary.LittleEndian, &h.table[i].isDeleted); err != nil {
			return err
		}
		h.table[i].key = int(key)
		h.table[i].value = int(value)
	}
	return nil
}

// JSON Serialization
type openEntry struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

type hashTableOpenJSON struct {
	Entries []openEntry `json:"entries"`
}

func (h *HashTableOpen) SerializeJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for writing: %w", err)
	}
	defer file.Close()

	entries := make([]openEntry, 0, h.size)
	for i := 0; i < h.capacity; i++ {
		if h.table[i].isOccupied && !h.table[i].isDeleted {
			entries = append(entries, openEntry{Key: h.table[i].key, Value: h.table[i].value})
		}
	}

	data := hashTableOpenJSON{Entries: entries}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (h *HashTableOpen) DeserializeJSON(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("cannot open file for reading: %w", err)
	}
	defer file.Close()

	// Clear existing data
	h.table = make([]HashEntry, h.capacity)
	for i := 0; i < h.capacity; i++ {
		h.table[i] = HashEntry{isOccupied: false, isDeleted: false}
	}
	h.size = 0

	var data hashTableOpenJSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for _, entry := range data.Entries {
		h.Insert(entry.Key, entry.Value)
	}
	return nil
}
