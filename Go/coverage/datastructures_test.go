package datastructures

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ==================== MyArray Tests ====================

func TestNewMyArray(t *testing.T) {
	arr := NewMyArray()
	assert.NotNil(t, arr)
	assert.Equal(t, 0, arr.GetLength())
	assert.Equal(t, 2, arr.capacity)
}

func TestMyArray_AddToEnd(t *testing.T) {
	arr := NewMyArray()

	// Test adding single element
	arr.AddToEnd(10)
	assert.Equal(t, 1, arr.GetLength())
	val, err := arr.GetAtIndex(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	// Test adding multiple elements
	arr.AddToEnd(20)
	arr.AddToEnd(30)
	assert.Equal(t, 3, arr.GetLength())

	// Test resize trigger (initial capacity is 2)
	arr.AddToEnd(40)
	arr.AddToEnd(50)
	assert.Equal(t, 5, arr.GetLength())
	assert.True(t, arr.capacity >= 5)
}

func TestMyArray_AddAtIndex(t *testing.T) {
	arr := NewMyArray()
	arr.AddToEnd(10)
	arr.AddToEnd(30)

	// Insert in middle
	err := arr.AddAtIndex(1, 20)
	assert.NoError(t, err)
	val, _ := arr.GetAtIndex(1)
	assert.Equal(t, 20, val)

	// Insert at beginning
	err = arr.AddAtIndex(0, 5)
	assert.NoError(t, err)
	val, _ = arr.GetAtIndex(0)
	assert.Equal(t, 5, val)

	// Insert at end
	err = arr.AddAtIndex(arr.GetLength(), 100)
	assert.NoError(t, err)

	// Test out of bounds
	err = arr.AddAtIndex(100, 999)
	assert.Error(t, err)
	assert.Equal(t, "index out of bounds", err.Error())
}

func TestMyArray_GetAtIndex(t *testing.T) {
	arr := NewMyArray()

	// Test empty array
	_, err := arr.GetAtIndex(0)
	assert.Error(t, err)

	arr.AddToEnd(10)
	arr.AddToEnd(20)

	// Valid index
	val, err := arr.GetAtIndex(0)
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	val, err = arr.GetAtIndex(1)
	assert.NoError(t, err)
	assert.Equal(t, 20, val)

	// Out of bounds
	_, err = arr.GetAtIndex(2)
	assert.Error(t, err)

	_, err = arr.GetAtIndex(-1)
	assert.Error(t, err)
}

func TestMyArray_RemoveAtIndex(t *testing.T) {
	arr := NewMyArray()
	arr.AddToEnd(10)
	arr.AddToEnd(20)
	arr.AddToEnd(30)

	// Remove middle
	err := arr.RemoveAtIndex(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, arr.GetLength())
	val, _ := arr.GetAtIndex(1)
	assert.Equal(t, 30, val)

	// Remove first
	err = arr.RemoveAtIndex(0)
	assert.NoError(t, err)
	val, _ = arr.GetAtIndex(0)
	assert.Equal(t, 30, val)

	// Out of bounds
	err = arr.RemoveAtIndex(10)
	assert.Error(t, err)
}

func TestMyArray_ReplaceAtIndex(t *testing.T) {
	arr := NewMyArray()
	arr.AddToEnd(10)
	arr.AddToEnd(20)

	// Valid replace
	err := arr.ReplaceAtIndex(0, 100)
	assert.NoError(t, err)
	val, _ := arr.GetAtIndex(0)
	assert.Equal(t, 100, val)

	// Out of bounds
	err = arr.ReplaceAtIndex(10, 999)
	assert.Error(t, err)
}

func TestMyArray_Print(t *testing.T) {
	arr := NewMyArray()
	arr.Print() // Empty array

	arr.AddToEnd(1)
	arr.AddToEnd(2)
	arr.AddToEnd(3)
	arr.Print() // Array with elements
}

func TestMyArray_Serialize(t *testing.T) {
	arr := NewMyArray()
	arr.AddToEnd(10)
	arr.AddToEnd(20)
	arr.AddToEnd(30)

	filename := "test_array.bin"
	defer os.Remove(filename)

	err := arr.Serialize(filename)
	assert.NoError(t, err)

	// Deserialize into new array
	arr2 := NewMyArray()
	err = arr2.Deserialize(filename)
	assert.NoError(t, err)

	assert.Equal(t, arr.GetLength(), arr2.GetLength())
	for i := 0; i < arr.GetLength(); i++ {
		v1, _ := arr.GetAtIndex(i)
		v2, _ := arr2.GetAtIndex(i)
		assert.Equal(t, v1, v2)
	}
}

func TestMyArray_SerializeErrors(t *testing.T) {
	arr := NewMyArray()

	// Invalid path
	err := arr.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = arr.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)
}

func TestMyArray_SerializeJSON(t *testing.T) {
	arr := NewMyArray()
	arr.AddToEnd(100)
	arr.AddToEnd(200)
	arr.AddToEnd(300)

	filename := "test_array.json"
	defer os.Remove(filename)

	err := arr.SerializeJSON(filename)
	assert.NoError(t, err)

	arr2 := NewMyArray()
	err = arr2.DeserializeJSON(filename)
	assert.NoError(t, err)

	assert.Equal(t, arr.GetLength(), arr2.GetLength())
}

func TestMyArray_DeserializeJSONErrors(t *testing.T) {
	arr := NewMyArray()

	err := arr.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}

func TestMyArray_Resize(t *testing.T) {
	arr := NewMyArray()

	// Force multiple resizes
	for i := 0; i < 100; i++ {
		arr.AddToEnd(i)
	}
	assert.Equal(t, 100, arr.GetLength())
	assert.True(t, arr.capacity >= 100)
}

// ==================== Stack Tests ====================

func TestNewMyStack(t *testing.T) {
	stack := NewMyStack()
	assert.NotNil(t, stack)
	assert.Nil(t, stack.topNode)
}

func TestMyStack_Push(t *testing.T) {
	stack := NewMyStack()
	stack.Push(10)
	stack.Push(20)

	val, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 20, val)
}

func TestMyStack_Pop(t *testing.T) {
	stack := NewMyStack()

	// Pop empty stack (should not panic)
	stack.Pop()

	stack.Push(10)
	stack.Push(20)
	stack.Pop()

	val, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	// Pop all elements
	stack.Pop()
	_, err = stack.Peek()
	assert.Error(t, err)
}

func TestMyStack_Peek(t *testing.T) {
	stack := NewMyStack()

	// Empty stack
	_, err := stack.Peek()
	assert.Error(t, err)
	assert.Equal(t, "stack empty", err.Error())

	stack.Push(42)
	val, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestMyStack_Print(t *testing.T) {
	stack := NewMyStack()
	stack.Print() // Empty

	stack.Push(1)
	stack.Push(2)
	stack.Print() // With elements
}

func TestMyStack_Serialize(t *testing.T) {
	stack := NewMyStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	filename := "test_stack.bin"
	defer os.Remove(filename)

	err := stack.Serialize(filename)
	assert.NoError(t, err)

	stack2 := NewMyStack()
	err = stack2.Deserialize(filename)
	assert.NoError(t, err)
}

func TestMyStack_SerializeJSON(t *testing.T) {
	stack := NewMyStack()
	stack.Push(100)
	stack.Push(200)

	filename := "test_stack.json"
	defer os.Remove(filename)

	err := stack.SerializeJSON(filename)
	assert.NoError(t, err)

	stack2 := NewMyStack()
	err = stack2.DeserializeJSON(filename)
	assert.NoError(t, err)
}

func TestMyStack_SerializeErrors(t *testing.T) {
	stack := NewMyStack()

	err := stack.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = stack.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = stack.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = stack.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}

// ==================== Queue Tests ====================

func TestNewMyQueue(t *testing.T) {
	queue := NewMyQueue()
	assert.NotNil(t, queue)
	assert.Nil(t, queue.frontNode)
	assert.Nil(t, queue.rearNode)
}

func TestMyQueue_Push(t *testing.T) {
	queue := NewMyQueue()

	// Push to empty queue
	queue.Push(10)
	val, err := queue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 10, val)

	// Push more elements
	queue.Push(20)
	queue.Push(30)
	val, _ = queue.Peek()
	assert.Equal(t, 10, val) // FIFO - first should still be 10
}

func TestMyQueue_Pop(t *testing.T) {
	queue := NewMyQueue()

	// Pop empty queue
	queue.Pop()

	queue.Push(10)
	queue.Push(20)

	queue.Pop()
	val, _ := queue.Peek()
	assert.Equal(t, 20, val)

	// Pop last element
	queue.Pop()
	_, err := queue.Peek()
	assert.Error(t, err)

	// Pop empty again
	queue.Pop()
}

func TestMyQueue_Peek(t *testing.T) {
	queue := NewMyQueue()

	// Empty queue
	_, err := queue.Peek()
	assert.Error(t, err)
	assert.Equal(t, "queue empty", err.Error())

	queue.Push(42)
	val, err := queue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 42, val)
}

func TestMyQueue_Print(t *testing.T) {
	queue := NewMyQueue()
	queue.Print()

	queue.Push(1)
	queue.Push(2)
	queue.Print()
}

func TestMyQueue_Serialize(t *testing.T) {
	queue := NewMyQueue()
	queue.Push(10)
	queue.Push(20)
	queue.Push(30)

	filename := "test_queue.bin"
	defer os.Remove(filename)

	err := queue.Serialize(filename)
	assert.NoError(t, err)

	queue2 := NewMyQueue()
	err = queue2.Deserialize(filename)
	assert.NoError(t, err)

	val, _ := queue2.Peek()
	assert.Equal(t, 10, val)
}

func TestMyQueue_SerializeJSON(t *testing.T) {
	queue := NewMyQueue()
	queue.Push(100)
	queue.Push(200)

	filename := "test_queue.json"
	defer os.Remove(filename)

	err := queue.SerializeJSON(filename)
	assert.NoError(t, err)

	queue2 := NewMyQueue()
	err = queue2.DeserializeJSON(filename)
	assert.NoError(t, err)
}

func TestMyQueue_SerializeErrors(t *testing.T) {
	queue := NewMyQueue()

	err := queue.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = queue.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = queue.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = queue.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}

// ==================== SinglyLinkedList Tests ====================

func TestNewSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.GetSize())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestSinglyLinkedList_PushFront(t *testing.T) {
	list := NewSinglyLinkedList()

	list.PushFront(10)
	assert.Equal(t, 1, list.GetSize())
	assert.True(t, list.Find(10))

	list.PushFront(20)
	assert.Equal(t, 2, list.GetSize())
	// 20 should be at front
	assert.Equal(t, 20, list.head.data)
}

func TestSinglyLinkedList_PushBack(t *testing.T) {
	list := NewSinglyLinkedList()

	// Push to empty list
	list.PushBack(10)
	assert.Equal(t, 1, list.GetSize())
	assert.Equal(t, list.head, list.tail)

	list.PushBack(20)
	assert.Equal(t, 2, list.GetSize())
	assert.Equal(t, 20, list.tail.data)
}

func TestSinglyLinkedList_InsertAfter(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack(10)
	list.PushBack(30)

	// Insert after first element
	err := list.InsertAfter(0, 20)
	assert.NoError(t, err)
	assert.Equal(t, 3, list.GetSize())

	// Insert after tail
	err = list.InsertAfter(2, 40)
	assert.NoError(t, err)
	assert.Equal(t, 40, list.tail.data)

	// Out of bounds
	err = list.InsertAfter(100, 999)
	assert.Error(t, err)
}

func TestSinglyLinkedList_InsertBefore(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack(20)
	list.PushBack(30)

	// Insert before first (becomes PushFront)
	err := list.InsertBefore(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 10, list.head.data)

	// Insert before middle
	err = list.InsertBefore(2, 25)
	assert.NoError(t, err)

	// Out of bounds
	err = list.InsertBefore(100, 999)
	assert.Error(t, err)
}

func TestSinglyLinkedList_PopFront(t *testing.T) {
	list := NewSinglyLinkedList()

	// Pop empty list
	list.PopFront()

	list.PushBack(10)
	list.PushBack(20)

	list.PopFront()
	assert.Equal(t, 1, list.GetSize())
	assert.Equal(t, 20, list.head.data)

	// Pop last element
	list.PopFront()
	assert.Equal(t, 0, list.GetSize())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestSinglyLinkedList_PopBack(t *testing.T) {
	list := NewSinglyLinkedList()

	// Pop empty list
	list.PopBack()

	list.PushBack(10)

	// Pop single element
	list.PopBack()
	assert.Equal(t, 0, list.GetSize())

	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	list.PopBack()
	assert.Equal(t, 2, list.GetSize())
	assert.Equal(t, 20, list.tail.data)
}

func TestSinglyLinkedList_RemoveAt(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	// Remove middle
	err := list.RemoveAt(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, list.GetSize())

	// Remove first
	err = list.RemoveAt(0)
	assert.NoError(t, err)

	// Remove last (which is now at index 0)
	list.PushBack(40)
	err = list.RemoveAt(1)
	assert.NoError(t, err)
	assert.Equal(t, 30, list.tail.data)

	// Out of bounds
	err = list.RemoveAt(100)
	assert.Error(t, err)
}

func TestSinglyLinkedList_RemoveByValue(t *testing.T) {
	list := NewSinglyLinkedList()

	// Remove from empty list
	list.RemoveByValue(10)

	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	// Remove head
	list.RemoveByValue(10)
	assert.Equal(t, 2, list.GetSize())

	// Remove middle
	list.PushFront(10)
	list.RemoveByValue(20)
	assert.Equal(t, 2, list.GetSize())

	// Remove tail
	list.RemoveByValue(30)
	assert.Equal(t, 1, list.GetSize())

	// Remove non-existent
	list.RemoveByValue(999)
	assert.Equal(t, 1, list.GetSize())
}

func TestSinglyLinkedList_Find(t *testing.T) {
	list := NewSinglyLinkedList()

	assert.False(t, list.Find(10))

	list.PushBack(10)
	list.PushBack(20)

	assert.True(t, list.Find(10))
	assert.True(t, list.Find(20))
	assert.False(t, list.Find(30))
}

func TestSinglyLinkedList_Print(t *testing.T) {
	list := NewSinglyLinkedList()
	list.Print()

	list.PushBack(1)
	list.PushBack(2)
	list.Print()
}

func TestSinglyLinkedList_Serialize(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	filename := "test_singly.bin"
	defer os.Remove(filename)

	err := list.Serialize(filename)
	assert.NoError(t, err)

	list2 := NewSinglyLinkedList()
	err = list2.Deserialize(filename)
	assert.NoError(t, err)

	assert.Equal(t, list.GetSize(), list2.GetSize())
}

func TestSinglyLinkedList_SerializeJSON(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack(100)
	list.PushBack(200)

	filename := "test_singly.json"
	defer os.Remove(filename)

	err := list.SerializeJSON(filename)
	assert.NoError(t, err)

	list2 := NewSinglyLinkedList()
	err = list2.DeserializeJSON(filename)
	assert.NoError(t, err)

	assert.Equal(t, list.GetSize(), list2.GetSize())
}

func TestSinglyLinkedList_SerializeErrors(t *testing.T) {
	list := NewSinglyLinkedList()

	err := list.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = list.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = list.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = list.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}


// ==================== DoublyLinkedList Tests ====================

func TestNewDoublyLinkedList(t *testing.T) {
	list := NewDoublyLinkedList()
	assert.NotNil(t, list)
	assert.Equal(t, 0, list.GetSize())
}

func TestDoublyLinkedList_PushFront(t *testing.T) {
	list := NewDoublyLinkedList()

	// Push to empty list
	list.PushFront(10)
	assert.Equal(t, 1, list.GetSize())
	assert.Equal(t, list.head, list.tail)

	list.PushFront(20)
	assert.Equal(t, 2, list.GetSize())
	assert.Equal(t, 20, list.head.data)
	assert.Equal(t, 10, list.tail.data)
	assert.Equal(t, list.head, list.tail.prev)
}

func TestDoublyLinkedList_PushBack(t *testing.T) {
	list := NewDoublyLinkedList()

	// Push to empty list
	list.PushBack(10)
	assert.Equal(t, 1, list.GetSize())

	list.PushBack(20)
	assert.Equal(t, 2, list.GetSize())
	assert.Equal(t, 20, list.tail.data)
	assert.Equal(t, list.tail, list.head.next)
}

func TestDoublyLinkedList_InsertAfter(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(10)
	list.PushBack(30)

	// Insert in middle
	err := list.InsertAfter(0, 20)
	assert.NoError(t, err)
	assert.Equal(t, 3, list.GetSize())

	// Insert after last (becomes PushBack)
	err = list.InsertAfter(2, 40)
	assert.NoError(t, err)
	assert.Equal(t, 40, list.tail.data)

	// Out of bounds
	err = list.InsertAfter(100, 999)
	assert.Error(t, err)
}

func TestDoublyLinkedList_InsertBefore(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(20)
	list.PushBack(30)

	// Insert before first (becomes PushFront)
	err := list.InsertBefore(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 10, list.head.data)

	// Insert in middle
	err = list.InsertBefore(2, 25)
	assert.NoError(t, err)

	// Out of bounds
	err = list.InsertBefore(100, 999)
	assert.Error(t, err)
}

func TestDoublyLinkedList_PopFront(t *testing.T) {
	list := NewDoublyLinkedList()

	// Pop empty
	list.PopFront()

	list.PushBack(10)
	list.PushBack(20)

	list.PopFront()
	assert.Equal(t, 1, list.GetSize())
	assert.Nil(t, list.head.prev)

	// Pop last element
	list.PopFront()
	assert.Equal(t, 0, list.GetSize())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestDoublyLinkedList_PopBack(t *testing.T) {
	list := NewDoublyLinkedList()

	// Pop empty
	list.PopBack()

	list.PushBack(10)
	list.PushBack(20)

	list.PopBack()
	assert.Equal(t, 1, list.GetSize())
	assert.Nil(t, list.tail.next)

	// Pop last element
	list.PopBack()
	assert.Equal(t, 0, list.GetSize())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestDoublyLinkedList_RemoveAt(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	// Remove middle
	err := list.RemoveAt(1)
	assert.NoError(t, err)
	assert.Equal(t, 2, list.GetSize())

	// Remove first
	err = list.RemoveAt(0)
	assert.NoError(t, err)

	// Remove last
	list.PushBack(40)
	err = list.RemoveAt(1)
	assert.NoError(t, err)

	// Out of bounds
	err = list.RemoveAt(100)
	assert.Error(t, err)
}

func TestDoublyLinkedList_RemoveByValue(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	// Remove head
	list.RemoveByValue(10)
	assert.Equal(t, 2, list.GetSize())

	// Remove tail
	list.RemoveByValue(30)
	assert.Equal(t, 1, list.GetSize())

	// Remove middle
	list.PushFront(10)
	list.PushBack(30)
	list.RemoveByValue(20)
	assert.Equal(t, 2, list.GetSize())

	// Remove non-existent
	list.RemoveByValue(999)
	assert.Equal(t, 2, list.GetSize())
}

func TestDoublyLinkedList_Find(t *testing.T) {
	list := NewDoublyLinkedList()

	assert.False(t, list.Find(10))

	list.PushBack(10)
	list.PushBack(20)

	assert.True(t, list.Find(10))
	assert.True(t, list.Find(20))
	assert.False(t, list.Find(30))
}

func TestDoublyLinkedList_Print(t *testing.T) {
	list := NewDoublyLinkedList()
	list.Print()

	list.PushBack(1)
	list.PushBack(2)
	list.Print()
}

func TestDoublyLinkedList_Serialize(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	filename := "test_doubly.bin"
	defer os.Remove(filename)

	err := list.Serialize(filename)
	assert.NoError(t, err)

	list2 := NewDoublyLinkedList()
	err = list2.Deserialize(filename)
	assert.NoError(t, err)

	assert.Equal(t, list.GetSize(), list2.GetSize())
}

func TestDoublyLinkedList_SerializeJSON(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(100)
	list.PushBack(200)

	filename := "test_doubly.json"
	defer os.Remove(filename)

	err := list.SerializeJSON(filename)
	assert.NoError(t, err)

	list2 := NewDoublyLinkedList()
	err = list2.DeserializeJSON(filename)
	assert.NoError(t, err)

	assert.Equal(t, list.GetSize(), list2.GetSize())
}

func TestDoublyLinkedList_SerializeErrors(t *testing.T) {
	list := NewDoublyLinkedList()

	err := list.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = list.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = list.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = list.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}



// ==================== HashTableChain Tests ====================

func TestNewHashTableChain(t *testing.T) {
	// Default capacity
	ht := NewHashTableChain(0)
	assert.NotNil(t, ht)
	assert.Equal(t, 8, ht.capacity)

	// Custom capacity
	ht = NewHashTableChain(16)
	assert.Equal(t, 16, ht.capacity)

	// Negative capacity
	ht = NewHashTableChain(-5)
	assert.Equal(t, 8, ht.capacity)
}

func TestHashTableChain_Insert(t *testing.T) {
	ht := NewHashTableChain(8)

	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Insert(9, 900) // Same bucket as 1 (collision)

	val, found := ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)

	val, found = ht.Get(9)
	assert.True(t, found)
	assert.Equal(t, 900, val)
}

func TestHashTableChain_Get(t *testing.T) {
	ht := NewHashTableChain(8)

	// Get from empty
	_, found := ht.Get(1)
	assert.False(t, found)

	ht.Insert(1, 100)
	ht.Insert(2, 200)

	val, found := ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)

	// Get non-existent
	_, found = ht.Get(999)
	assert.False(t, found)
}

func TestHashTableChain_Remove(t *testing.T) {
	ht := NewHashTableChain(8)

	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Insert(9, 900) // Collision with 1

	// Remove existing
	ht.Remove(1)
	_, found := ht.Get(1)
	assert.False(t, found)

	// 9 should still exist
	val, found := ht.Get(9)
	assert.True(t, found)
	assert.Equal(t, 900, val)

	// Remove from chain
	ht.Insert(1, 100)
	ht.Remove(9)
	_, found = ht.Get(9)
	assert.False(t, found)

	// Remove non-existent
	ht.Remove(999)
}

func TestHashTableChain_RemoveHead(t *testing.T) {
	ht := NewHashTableChain(8)

	// Test removing head of chain
	ht.Insert(1, 100)
	ht.Insert(9, 900) // Same bucket
	ht.Insert(17, 1700) // Same bucket

	// Remove head
	ht.Remove(17) // 17 should be head of chain
	_, found := ht.Get(17)
	assert.False(t, found)

	// Others should remain
	_, found = ht.Get(1)
	assert.True(t, found)
}

func TestHashTableChain_NegativeKey(t *testing.T) {
	ht := NewHashTableChain(8)

	ht.Insert(-5, 500)
	val, found := ht.Get(-5)
	assert.True(t, found)
	assert.Equal(t, 500, val)
}

func TestHashTableChain_Print(t *testing.T) {
	ht := NewHashTableChain(8)
	ht.Print()

	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Insert(9, 900)
	ht.Print()
}

func TestHashTableChain_Serialize(t *testing.T) {
	ht := NewHashTableChain(8)
	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Insert(3, 300)

	filename := "test_htchain.bin"
	defer os.Remove(filename)

	err := ht.Serialize(filename)
	assert.NoError(t, err)

	ht2 := NewHashTableChain(8)
	err = ht2.Deserialize(filename)
	assert.NoError(t, err)

	val, found := ht2.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)
}

func TestHashTableChain_SerializeJSON(t *testing.T) {
	ht := NewHashTableChain(8)
	ht.Insert(1, 100)
	ht.Insert(2, 200)

	filename := "test_htchain.json"
	defer os.Remove(filename)

	err := ht.SerializeJSON(filename)
	assert.NoError(t, err)

	ht2 := NewHashTableChain(8)
	err = ht2.DeserializeJSON(filename)
	assert.NoError(t, err)
}

func TestHashTableChain_SerializeErrors(t *testing.T) {
	ht := NewHashTableChain(8)

	err := ht.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = ht.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = ht.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = ht.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}

// ==================== HashTableOpen Tests ====================

func TestNewHashTableOpen(t *testing.T) {
	// Default capacity
	ht := NewHashTableOpen(0)
	assert.NotNil(t, ht)
	assert.Equal(t, 8, ht.capacity)

	// Custom capacity
	ht = NewHashTableOpen(16)
	assert.Equal(t, 16, ht.capacity)

	// Negative capacity
	ht = NewHashTableOpen(-5)
	assert.Equal(t, 8, ht.capacity)
}

func TestHashTableOpen_Insert(t *testing.T) {
	ht := NewHashTableOpen(8)

	ht.Insert(1, 100)
	ht.Insert(2, 200)

	val, found := ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)

	// Update existing key
	ht.Insert(1, 150)
	val, found = ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 150, val)
}

func TestHashTableOpen_InsertCollision(t *testing.T) {
	ht := NewHashTableOpen(8)

	// Insert keys that will collide
	ht.Insert(1, 100)
	ht.Insert(9, 900) // Same bucket as 1

	val, found := ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)

	val, found = ht.Get(9)
	assert.True(t, found)
	assert.Equal(t, 900, val)
}

func TestHashTableOpen_InsertResize(t *testing.T) {
	ht := NewHashTableOpen(4)

	// Insert enough to trigger resize (>70% load)
	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Insert(3, 300)
	ht.Insert(4, 400)

	assert.True(t, ht.capacity > 4)

	// All values should still be accessible
	for i := 1; i <= 4; i++ {
		val, found := ht.Get(i)
		assert.True(t, found)
		assert.Equal(t, i*100, val)
	}
}

func TestHashTableOpen_Get(t *testing.T) {
	ht := NewHashTableOpen(8)

	// Get from empty
	_, found := ht.Get(1)
	assert.False(t, found)

	ht.Insert(1, 100)

	val, found := ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)

	// Get non-existent
	_, found = ht.Get(999)
	assert.False(t, found)
}

func TestHashTableOpen_Remove(t *testing.T) {
	ht := NewHashTableOpen(8)

	ht.Insert(1, 100)
	ht.Insert(9, 900) // Collision

	ht.Remove(1)
	_, found := ht.Get(1)
	assert.False(t, found)

	// 9 should still be found (linear probing with tombstone)
	val, found := ht.Get(9)
	assert.True(t, found)
	assert.Equal(t, 900, val)

	// Remove non-existent
	ht.Remove(999)
}

func TestHashTableOpen_NegativeKey(t *testing.T) {
	ht := NewHashTableOpen(8)

	ht.Insert(-5, 500)
	val, found := ht.Get(-5)
	assert.True(t, found)
	assert.Equal(t, 500, val)

	ht.Remove(-5)
	_, found = ht.Get(-5)
	assert.False(t, found)
}

func TestHashTableOpen_Print(t *testing.T) {
	ht := NewHashTableOpen(8)
	ht.Print()

	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Print()
}

func TestHashTableOpen_Serialize(t *testing.T) {
	ht := NewHashTableOpen(8)
	ht.Insert(1, 100)
	ht.Insert(2, 200)
	ht.Insert(3, 300)

	filename := "test_htopen.bin"
	defer os.Remove(filename)

	err := ht.Serialize(filename)
	assert.NoError(t, err)

	ht2 := NewHashTableOpen(8)
	err = ht2.Deserialize(filename)
	assert.NoError(t, err)

	val, found := ht2.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)
}

func TestHashTableOpen_SerializeJSON(t *testing.T) {
	ht := NewHashTableOpen(8)
	ht.Insert(1, 100)
	ht.Insert(2, 200)

	filename := "test_htopen.json"
	defer os.Remove(filename)

	err := ht.SerializeJSON(filename)
	assert.NoError(t, err)

	ht2 := NewHashTableOpen(8)
	err = ht2.DeserializeJSON(filename)
	assert.NoError(t, err)

	val, found := ht2.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)
}

func TestHashTableOpen_SerializeErrors(t *testing.T) {
	ht := NewHashTableOpen(8)

	err := ht.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = ht.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = ht.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = ht.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}

func TestHashTableOpen_FullTableProbing(t *testing.T) {
	// Test linear probing when table is nearly full
	ht := NewHashTableOpen(8)

	// Fill table to trigger wraparound probing
	for i := 0; i < 5; i++ {
		ht.Insert(i, i*100)
	}

	// Verify all values
	for i := 0; i < 5; i++ {
		val, found := ht.Get(i)
		assert.True(t, found)
		assert.Equal(t, i*100, val)
	}
}

// ==================== AVLTree Tests ====================

func TestNewAVLTree(t *testing.T) {
	tree := NewAVLTree()
	assert.NotNil(t, tree)
	assert.Nil(t, tree.root)
}

func TestAVLTree_Insert(t *testing.T) {
	tree := NewAVLTree()

	tree.Insert(10)
	assert.True(t, tree.Find(10))

	tree.Insert(20)
	tree.Insert(5)
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(5))

	// Insert duplicate (should not add)
	tree.Insert(10)
}

func TestAVLTree_InsertLLRotation(t *testing.T) {
	tree := NewAVLTree()

	// LL case: insert in descending order
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(10)

	// Tree should be balanced
	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(30))
}

func TestAVLTree_InsertRRRotation(t *testing.T) {
	tree := NewAVLTree()

	// RR case: insert in ascending order
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)

	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(30))
}

func TestAVLTree_InsertLRRotation(t *testing.T) {
	tree := NewAVLTree()

	// LR case
	tree.Insert(30)
	tree.Insert(10)
	tree.Insert(20)

	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(30))
}

func TestAVLTree_InsertRLRotation(t *testing.T) {
	tree := NewAVLTree()

	// RL case
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(20)

	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(30))
}

func TestAVLTree_Remove(t *testing.T) {
	tree := NewAVLTree()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(25)

	// Remove leaf
	tree.Remove(25)
	assert.False(t, tree.Find(25))

	// Remove node with one child
	tree.Remove(20)
	assert.False(t, tree.Find(20))

	// Remove node with two children
	tree.Insert(20)
	tree.Insert(25)
	tree.Remove(10)
	assert.False(t, tree.Find(10))

	// Remove root
	tree.Remove(15)
	assert.False(t, tree.Find(15))

	// Remove from empty tree
	emptyTree := NewAVLTree()
	emptyTree.Remove(10)
}

func TestAVLTree_RemoveRebalance(t *testing.T) {
	tree := NewAVLTree()

	// Build tree that will need rebalancing after delete
	tree.Insert(50)
	tree.Insert(25)
	tree.Insert(75)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(60)
	tree.Insert(80)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(27)
	tree.Insert(55)
	tree.Insert(1)

	// Remove to trigger rebalance
	tree.Remove(80)
	tree.Remove(60)

	// Tree should still be balanced and searchable
	assert.True(t, tree.Find(50))
	assert.True(t, tree.Find(25))
}

func TestAVLTree_Find(t *testing.T) {
	tree := NewAVLTree()

	// Find in empty tree
	assert.False(t, tree.Find(10))

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(5))
	assert.True(t, tree.Find(15))
	assert.False(t, tree.Find(20))
}

func TestAVLTree_Print(t *testing.T) {
	tree := NewAVLTree()
	tree.Print()

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Print()
}

func TestAVLTree_Serialize(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	filename := "test_avl.bin"
	defer os.Remove(filename)

	err := tree.Serialize(filename)
	assert.NoError(t, err)

	tree2 := NewAVLTree()
	err = tree2.Deserialize(filename)
	assert.NoError(t, err)

	assert.True(t, tree2.Find(10))
	assert.True(t, tree2.Find(5))
	assert.True(t, tree2.Find(15))
}

func TestAVLTree_SerializeJSON(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	filename := "test_avl.json"
	defer os.Remove(filename)

	err := tree.SerializeJSON(filename)
	assert.NoError(t, err)

	tree2 := NewAVLTree()
	err = tree2.DeserializeJSON(filename)
	assert.NoError(t, err)

	assert.True(t, tree2.Find(10))
	assert.True(t, tree2.Find(5))
	assert.True(t, tree2.Find(15))
}

func TestAVLTree_SerializeErrors(t *testing.T) {
	tree := NewAVLTree()

	err := tree.Serialize("/invalid/path/file.bin")
	assert.Error(t, err)

	err = tree.Deserialize("/nonexistent/file.bin")
	assert.Error(t, err)

	err = tree.SerializeJSON("/invalid/path/file.json")
	assert.Error(t, err)

	err = tree.DeserializeJSON("/nonexistent/file.json")
	assert.Error(t, err)
}

func TestAVLTree_SerializeEmptyTree(t *testing.T) {
	tree := NewAVLTree()

	filename := "test_avl_empty.bin"
	defer os.Remove(filename)

	err := tree.Serialize(filename)
	assert.NoError(t, err)

	tree2 := NewAVLTree()
	err = tree2.Deserialize(filename)
	assert.NoError(t, err)
	assert.Nil(t, tree2.root)
}

func TestAVLTree_SerializeJSONEmptyTree(t *testing.T) {
	tree := NewAVLTree()

	filename := "test_avl_empty.json"
	defer os.Remove(filename)

	err := tree.SerializeJSON(filename)
	assert.NoError(t, err)

	tree2 := NewAVLTree()
	err = tree2.DeserializeJSON(filename)
	assert.NoError(t, err)
}

func TestAVLTree_ComplexOperations(t *testing.T) {
	tree := NewAVLTree()

	// Insert many values to trigger multiple rotations
	values := []int{50, 25, 75, 10, 30, 60, 80, 5, 15, 27, 55, 1, 100, 90, 85}
	for _, v := range values {
		tree.Insert(v)
	}

	// Verify all inserted
	for _, v := range values {
		assert.True(t, tree.Find(v))
	}

	// Remove some and verify
	tree.Remove(50)
	tree.Remove(25)
	tree.Remove(1)

	assert.False(t, tree.Find(50))
	assert.False(t, tree.Find(25))
	assert.False(t, tree.Find(1))

	// Others should still exist
	assert.True(t, tree.Find(75))
	assert.True(t, tree.Find(100))
}

func TestAVLTree_RemoveNodeWithTwoChildren(t *testing.T) {
	tree := NewAVLTree()

	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(25)
	tree.Insert(35)

	// Remove node (30) with two children
	tree.Remove(30)
	assert.False(t, tree.Find(30))
	assert.True(t, tree.Find(25))
	assert.True(t, tree.Find(35))
}

func TestAVLTree_RemoveRootWithChildren(t *testing.T) {
	tree := NewAVLTree()

	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)

	// Remove root
	tree.Remove(20)
	assert.False(t, tree.Find(20))
	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(30))
}

// ==================== Helper Functions ====================

func writeUint64(file *os.File, val uint64) error {
	bytes := make([]byte, 8)
	bytes[0] = byte(val)
	bytes[1] = byte(val >> 8)
	bytes[2] = byte(val >> 16)
	bytes[3] = byte(val >> 24)
	bytes[4] = byte(val >> 32)
	bytes[5] = byte(val >> 40)
	bytes[6] = byte(val >> 48)
	bytes[7] = byte(val >> 56)
	_, err := file.Write(bytes)
	return err
}

// ==================== Additional Edge Case Tests ====================

func TestMyArray_EdgeCases(t *testing.T) {
	arr := NewMyArray()

	// Test boundary conditions
	arr.AddToEnd(0)
	arr.AddToEnd(-1)
	arr.AddToEnd(2147483647) // Max int32

	assert.Equal(t, 3, arr.GetLength())
}

func TestStack_MultiplePopPush(t *testing.T) {
	stack := NewMyStack()

	// Push and pop multiple times
	for i := 0; i < 100; i++ {
		stack.Push(i)
	}

	for i := 99; i >= 0; i-- {
		val, err := stack.Peek()
		assert.NoError(t, err)
		assert.Equal(t, i, val)
		stack.Pop()
	}

	_, err := stack.Peek()
	assert.Error(t, err)
}

func TestQueue_MultipleOperations(t *testing.T) {
	queue := NewMyQueue()

	// Interleaved push and pop
	queue.Push(1)
	queue.Push(2)
	queue.Pop()
	queue.Push(3)

	val, _ := queue.Peek()
	assert.Equal(t, 2, val)
}

func TestHashTableOpen_InsertAfterDelete(t *testing.T) {
	ht := NewHashTableOpen(8)

	ht.Insert(1, 100)
	ht.Remove(1)

	// Insert same key after delete
	ht.Insert(1, 200)
	val, found := ht.Get(1)
	assert.True(t, found)
	assert.Equal(t, 200, val)
}

func TestHashTableChain_MultipleCollisions(t *testing.T) {
	ht := NewHashTableChain(4)

	// Insert multiple keys that hash to same bucket
	ht.Insert(0, 0)
	ht.Insert(4, 400)
	ht.Insert(8, 800)
	ht.Insert(12, 1200)

	for _, k := range []int{0, 4, 8, 12} {
		val, found := ht.Get(k)
		assert.True(t, found)
		assert.Equal(t, k*100, val)
	}
}

func TestAVLTree_DeserializeExisting(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)

	filename := "test_avl_exist.bin"
	defer os.Remove(filename)

	err := tree.Serialize(filename)
	require.NoError(t, err)

	// Create tree with existing data
	tree2 := NewAVLTree()
	tree2.Insert(100)
	tree2.Insert(200)

	// Deserialize should replace existing data
	err = tree2.Deserialize(filename)
	assert.NoError(t, err)

	assert.True(t, tree2.Find(1))
	assert.True(t, tree2.Find(2))
	assert.True(t, tree2.Find(3))
}

func TestDoublyLinkedList_BidirectionalTraversal(t *testing.T) {
	list := NewDoublyLinkedList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	// Forward traversal
	curr := list.head
	expected := 1
	for curr != nil {
		assert.Equal(t, expected, curr.data)
		expected++
		curr = curr.next
	}

	// Backward traversal
	curr = list.tail
	expected = 3
	for curr != nil {
		assert.Equal(t, expected, curr.data)
		expected--
		curr = curr.prev
	}
}

func TestSinglyLinkedList_RemoveLastElement(t *testing.T) {
	list := NewSinglyLinkedList()
	list.PushBack(10)

	list.RemoveAt(0)
	assert.Equal(t, 0, list.GetSize())
	assert.Nil(t, list.head)
	assert.Nil(t, list.tail)
}

func TestMyArray_DeserializeWithResize(t *testing.T) {
	arr := NewMyArray()
	for i := 0; i < 50; i++ {
		arr.AddToEnd(i)
	}

	filename := "test_array_resize.bin"
	defer os.Remove(filename)

	err := arr.Serialize(filename)
	require.NoError(t, err)

	// Deserialize into small array
	arr2 := NewMyArray()
	err = arr2.Deserialize(filename)
	assert.NoError(t, err)

	assert.Equal(t, 50, arr2.GetLength())
	assert.True(t, arr2.capacity >= 50)
}

func TestMyArray_DeserializeJSONWithResize(t *testing.T) {
	arr := NewMyArray()
	for i := 0; i < 50; i++ {
		arr.AddToEnd(i)
	}

	filename := "test_array_resize.json"
	defer os.Remove(filename)

	err := arr.SerializeJSON(filename)
	require.NoError(t, err)

	arr2 := NewMyArray()
	err = arr2.DeserializeJSON(filename)
	assert.NoError(t, err)

	assert.Equal(t, 50, arr2.GetLength())
}

func TestHashTableChain_DeserializeWithExistingData(t *testing.T) {
	ht := NewHashTableChain(8)
	ht.Insert(1, 100)
	ht.Insert(2, 200)

	filename := "test_htchain_existing.bin"
	defer os.Remove(filename)

	err := ht.Serialize(filename)
	require.NoError(t, err)

	// Create table with existing data
	ht2 := NewHashTableChain(8)
	ht2.Insert(10, 1000)
	ht2.Insert(20, 2000)

	err = ht2.Deserialize(filename)
	assert.NoError(t, err)
}

func TestHashTableChain_DeserializeJSONWithExistingData(t *testing.T) {
	ht := NewHashTableChain(8)
	ht.Insert(1, 100)

	filename := "test_htchain_existing.json"
	defer os.Remove(filename)

	err := ht.SerializeJSON(filename)
	require.NoError(t, err)

	ht2 := NewHashTableChain(8)
	ht2.Insert(10, 1000)

	err = ht2.DeserializeJSON(filename)
	assert.NoError(t, err)
}

func TestHashTableOpen_DeserializeJSONWithExistingData(t *testing.T) {
	ht := NewHashTableOpen(8)
	ht.Insert(1, 100)

	filename := "test_htopen_existing.json"
	defer os.Remove(filename)

	err := ht.SerializeJSON(filename)
	require.NoError(t, err)

	ht2 := NewHashTableOpen(8)
	ht2.Insert(10, 1000)

	err = ht2.DeserializeJSON(filename)
	assert.NoError(t, err)

	val, found := ht2.Get(1)
	assert.True(t, found)
	assert.Equal(t, 100, val)
}

func TestStack_DeserializeOverwritesExisting(t *testing.T) {
	stack := NewMyStack()
	stack.Push(1)
	stack.Push(2)

	filename := "test_stack_overwrite.bin"
	defer os.Remove(filename)

	err := stack.Serialize(filename)
	require.NoError(t, err)

	stack2 := NewMyStack()
	stack2.Push(100)
	stack2.Push(200)

	err = stack2.Deserialize(filename)
	assert.NoError(t, err)
}

func TestQueue_DeserializeOverwritesExisting(t *testing.T) {
	queue := NewMyQueue()
	queue.Push(1)
	queue.Push(2)

	filename := "test_queue_overwrite.bin"
	defer os.Remove(filename)

	err := queue.Serialize(filename)
	require.NoError(t, err)

	queue2 := NewMyQueue()
	queue2.Push(100)

	err = queue2.Deserialize(filename)
	assert.NoError(t, err)
}

func TestAVLTree_RemoveLeafRight(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	// Remove right leaf
	tree.Remove(15)
	assert.False(t, tree.Find(15))
	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(5))
}

func TestAVLTree_RemoveNodeWithLeftChildOnly(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(12)

	// 15 has only left child (12)
	tree.Remove(15)
	assert.False(t, tree.Find(15))
	assert.True(t, tree.Find(12))
}

func TestAVLTree_RemoveNodeWithRightChildOnly(t *testing.T) {
	tree := NewAVLTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(20)

	// 15 has only right child (20)
	tree.Remove(15)
	assert.False(t, tree.Find(15))
	assert.True(t, tree.Find(20))
}

func TestAVLTree_DeleteRebalanceRightHeavy(t *testing.T) {
	tree := NewAVLTree()

	// Create right-heavy tree
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(25)
	tree.Insert(30)

	// Remove left side to trigger rebalance
	tree.Remove(5)

	assert.True(t, tree.Find(10))
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(25))
}

func TestAVLTree_DeleteRebalanceLeftHeavy(t *testing.T) {
	tree := NewAVLTree()

	// Create left-heavy tree
	tree.Insert(30)
	tree.Insert(20)
	tree.Insert(35)
	tree.Insert(10)
	tree.Insert(25)
	tree.Insert(5)

	// Remove right side to trigger rebalance
	tree.Remove(35)

	assert.True(t, tree.Find(30))
	assert.True(t, tree.Find(20))
	assert.True(t, tree.Find(10))
}
