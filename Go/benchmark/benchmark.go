package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	ds "benchmark/datastructures"
)

// BenchmarkResult holds the results of a single benchmark
type BenchmarkResult struct {
	Operation     string
	DataStructure string
	NumElements   int
	Duration      time.Duration
	OpsPerSecond  float64
	MemoryUsed    uint64
}

// BenchmarkSuite manages all benchmarks
type BenchmarkSuite struct {
	results []BenchmarkResult
	sizes   []int
}

func NewBenchmarkSuite() *BenchmarkSuite {
	return &BenchmarkSuite{
		results: make([]BenchmarkResult, 0),
		sizes:   []int{1000, 10000, 100000},
	}
}

// getMemoryUsage returns current memory allocation in bytes
func getMemoryUsage() uint64 {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// calcMemoryDiff safely calculates memory difference
func calcMemoryDiff(before, after uint64) uint64 {
	if after > before {
		return after - before
	}
	return 0
}

// formatDuration formats duration in a human-readable way
func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2fµs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.2fms", float64(d.Nanoseconds())/1000000)
	}
	return fmt.Sprintf("%.2fs", d.Seconds())
}

// formatMemory formats memory in human-readable way
func formatMemory(bytes uint64) string {
	if bytes < 1024 {
		return fmt.Sprintf("%dB", bytes)
	} else if bytes < 1024*1024 {
		return fmt.Sprintf("%.2fKB", float64(bytes)/1024)
	}
	return fmt.Sprintf("%.2fMB", float64(bytes)/(1024*1024))
}

// formatOps formats operations per second
func formatOps(ops float64) string {
	if ops >= 1000000 {
		return fmt.Sprintf("%.2fM ops/s", ops/1000000)
	} else if ops >= 1000 {
		return fmt.Sprintf("%.2fK ops/s", ops/1000)
	}
	return fmt.Sprintf("%.2f ops/s", ops)
}

// generateRandomData generates a slice of random integers
func generateRandomData(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(n * 10)
	}
	return data
}

// generateSequentialData generates sequential integers
func generateSequentialData(n int) []int {
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = i
	}
	return data
}

// ============================================================================
// ARRAY BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkArrayInsertEnd(n int) BenchmarkResult {
	data := generateRandomData(n)
	arr := ds.NewMyArray()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		arr.AddToEnd(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Insert End",
		DataStructure: "MyArray",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkArrayInsertMiddle(n int) BenchmarkResult {
	arr := ds.NewMyArray()
	// Pre-populate with some elements
	for i := 0; i < n/10; i++ {
		arr.AddToEnd(i)
	}

	insertCount := n / 10
	memBefore := getMemoryUsage()
	start := time.Now()

	for i := 0; i < insertCount; i++ {
		idx := arr.GetLength() / 2
		arr.AddAtIndex(idx, i)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Insert Middle",
		DataStructure: "MyArray",
		NumElements:   insertCount,
		Duration:      duration,
		OpsPerSecond:  float64(insertCount) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkArrayRandomAccess(n int) BenchmarkResult {
	arr := ds.NewMyArray()
	for i := 0; i < n; i++ {
		arr.AddToEnd(i)
	}

	accessCount := n
	indices := generateRandomData(accessCount)
	for i := range indices {
		indices[i] = indices[i] % n
	}

	start := time.Now()

	for _, idx := range indices {
		arr.GetAtIndex(idx)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Random Access",
		DataStructure: "MyArray",
		NumElements:   accessCount,
		Duration:      duration,
		OpsPerSecond:  float64(accessCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkArrayRemove(n int) BenchmarkResult {
	arr := ds.NewMyArray()
	for i := 0; i < n; i++ {
		arr.AddToEnd(i)
	}

	removeCount := n / 2
	start := time.Now()

	for i := 0; i < removeCount; i++ {
		arr.RemoveAtIndex(0)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Remove Front",
		DataStructure: "MyArray",
		NumElements:   removeCount,
		Duration:      duration,
		OpsPerSecond:  float64(removeCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// SINGLY LINKED LIST BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkSLLPushFront(n int) BenchmarkResult {
	data := generateRandomData(n)
	list := ds.NewSinglyLinkedList()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		list.PushFront(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Push Front",
		DataStructure: "SinglyLinkedList",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkSLLPushBack(n int) BenchmarkResult {
	data := generateRandomData(n)
	list := ds.NewSinglyLinkedList()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		list.PushBack(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Push Back",
		DataStructure: "SinglyLinkedList",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkSLLFind(n int) BenchmarkResult {
	list := ds.NewSinglyLinkedList()
	for i := 0; i < n; i++ {
		list.PushBack(i)
	}

	searchCount := 1000
	targets := generateRandomData(searchCount)
	for i := range targets {
		targets[i] = targets[i] % (n * 2) // Some will be found, some won't
	}

	start := time.Now()

	for _, target := range targets {
		list.Find(target)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Find",
		DataStructure: "SinglyLinkedList",
		NumElements:   searchCount,
		Duration:      duration,
		OpsPerSecond:  float64(searchCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkSLLPopFront(n int) BenchmarkResult {
	list := ds.NewSinglyLinkedList()
	for i := 0; i < n; i++ {
		list.PushBack(i)
	}

	start := time.Now()

	for i := 0; i < n; i++ {
		list.PopFront()
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Pop Front",
		DataStructure: "SinglyLinkedList",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// DOUBLY LINKED LIST BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkDLLPushFront(n int) BenchmarkResult {
	data := generateRandomData(n)
	list := ds.NewDoublyLinkedList()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		list.PushFront(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Push Front",
		DataStructure: "DoublyLinkedList",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkDLLPushBack(n int) BenchmarkResult {
	data := generateRandomData(n)
	list := ds.NewDoublyLinkedList()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		list.PushBack(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Push Back",
		DataStructure: "DoublyLinkedList",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkDLLPopBack(n int) BenchmarkResult {
	list := ds.NewDoublyLinkedList()
	for i := 0; i < n; i++ {
		list.PushBack(i)
	}

	start := time.Now()

	for i := 0; i < n; i++ {
		list.PopBack()
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Pop Back",
		DataStructure: "DoublyLinkedList",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkDLLFind(n int) BenchmarkResult {
	list := ds.NewDoublyLinkedList()
	for i := 0; i < n; i++ {
		list.PushBack(i)
	}

	searchCount := 1000
	targets := generateRandomData(searchCount)
	for i := range targets {
		targets[i] = targets[i] % (n * 2)
	}

	start := time.Now()

	for _, target := range targets {
		list.Find(target)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Find",
		DataStructure: "DoublyLinkedList",
		NumElements:   searchCount,
		Duration:      duration,
		OpsPerSecond:  float64(searchCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// STACK BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkStackPush(n int) BenchmarkResult {
	data := generateRandomData(n)
	stack := ds.NewMyStack()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		stack.Push(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Push",
		DataStructure: "Stack",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkStackPop(n int) BenchmarkResult {
	stack := ds.NewMyStack()
	for i := 0; i < n; i++ {
		stack.Push(i)
	}

	start := time.Now()

	for i := 0; i < n; i++ {
		stack.Pop()
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Pop",
		DataStructure: "Stack",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkStackPeek(n int) BenchmarkResult {
	stack := ds.NewMyStack()
	for i := 0; i < n; i++ {
		stack.Push(i)
	}

	peekCount := n
	start := time.Now()

	for i := 0; i < peekCount; i++ {
		stack.Peek()
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Peek",
		DataStructure: "Stack",
		NumElements:   peekCount,
		Duration:      duration,
		OpsPerSecond:  float64(peekCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// QUEUE BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkQueuePush(n int) BenchmarkResult {
	data := generateRandomData(n)
	queue := ds.NewMyQueue()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		queue.Push(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Push",
		DataStructure: "Queue",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkQueuePop(n int) BenchmarkResult {
	queue := ds.NewMyQueue()
	for i := 0; i < n; i++ {
		queue.Push(i)
	}

	start := time.Now()

	for i := 0; i < n; i++ {
		queue.Pop()
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Pop",
		DataStructure: "Queue",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkQueuePeek(n int) BenchmarkResult {
	queue := ds.NewMyQueue()
	for i := 0; i < n; i++ {
		queue.Push(i)
	}

	peekCount := n
	start := time.Now()

	for i := 0; i < peekCount; i++ {
		queue.Peek()
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Peek",
		DataStructure: "Queue",
		NumElements:   peekCount,
		Duration:      duration,
		OpsPerSecond:  float64(peekCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// HASH TABLE (CHAINING) BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkHashChainInsert(n int) BenchmarkResult {
	data := generateRandomData(n)
	ht := ds.NewHashTableChain(n / 4)

	memBefore := getMemoryUsage()
	start := time.Now()

	for i, v := range data {
		ht.Insert(i, v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Insert",
		DataStructure: "HashTableChain",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkHashChainGet(n int) BenchmarkResult {
	ht := ds.NewHashTableChain(n / 4)
	for i := 0; i < n; i++ {
		ht.Insert(i, i*2)
	}

	getCount := n
	keys := generateRandomData(getCount)
	for i := range keys {
		keys[i] = keys[i] % (n * 2)
	}

	start := time.Now()

	for _, key := range keys {
		ht.Get(key)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Get",
		DataStructure: "HashTableChain",
		NumElements:   getCount,
		Duration:      duration,
		OpsPerSecond:  float64(getCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkHashChainRemove(n int) BenchmarkResult {
	ht := ds.NewHashTableChain(n / 4)
	for i := 0; i < n; i++ {
		ht.Insert(i, i*2)
	}

	removeCount := n / 2
	start := time.Now()

	for i := 0; i < removeCount; i++ {
		ht.Remove(i)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Remove",
		DataStructure: "HashTableChain",
		NumElements:   removeCount,
		Duration:      duration,
		OpsPerSecond:  float64(removeCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// HASH TABLE (OPEN ADDRESSING) BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkHashOpenInsert(n int) BenchmarkResult {
	data := generateRandomData(n)
	ht := ds.NewHashTableOpen(n / 4)

	memBefore := getMemoryUsage()
	start := time.Now()

	for i, v := range data {
		ht.Insert(i, v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Insert",
		DataStructure: "HashTableOpen",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkHashOpenGet(n int) BenchmarkResult {
	ht := ds.NewHashTableOpen(n / 4)
	for i := 0; i < n; i++ {
		ht.Insert(i, i*2)
	}

	getCount := n
	keys := generateRandomData(getCount)
	for i := range keys {
		keys[i] = keys[i] % (n * 2)
	}

	start := time.Now()

	for _, key := range keys {
		ht.Get(key)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Get",
		DataStructure: "HashTableOpen",
		NumElements:   getCount,
		Duration:      duration,
		OpsPerSecond:  float64(getCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkHashOpenRemove(n int) BenchmarkResult {
	ht := ds.NewHashTableOpen(n / 4)
	for i := 0; i < n; i++ {
		ht.Insert(i, i*2)
	}

	removeCount := n / 2
	start := time.Now()

	for i := 0; i < removeCount; i++ {
		ht.Remove(i)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Remove",
		DataStructure: "HashTableOpen",
		NumElements:   removeCount,
		Duration:      duration,
		OpsPerSecond:  float64(removeCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// AVL TREE BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkAVLInsertRandom(n int) BenchmarkResult {
	data := generateRandomData(n)
	tree := ds.NewAVLTree()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		tree.Insert(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Insert Random",
		DataStructure: "AVLTree",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkAVLInsertSequential(n int) BenchmarkResult {
	data := generateSequentialData(n)
	tree := ds.NewAVLTree()

	memBefore := getMemoryUsage()
	start := time.Now()

	for _, v := range data {
		tree.Insert(v)
	}

	duration := time.Since(start)
	memAfter := getMemoryUsage()

	return BenchmarkResult{
		Operation:     "Insert Sequential",
		DataStructure: "AVLTree",
		NumElements:   n,
		Duration:      duration,
		OpsPerSecond:  float64(n) / duration.Seconds(),
		MemoryUsed:    calcMemoryDiff(memBefore, memAfter),
	}
}

func (bs *BenchmarkSuite) benchmarkAVLFind(n int) BenchmarkResult {
	tree := ds.NewAVLTree()
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}

	searchCount := n
	targets := generateRandomData(searchCount)
	for i := range targets {
		targets[i] = targets[i] % (n * 2)
	}

	start := time.Now()

	for _, target := range targets {
		tree.Find(target)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Find",
		DataStructure: "AVLTree",
		NumElements:   searchCount,
		Duration:      duration,
		OpsPerSecond:  float64(searchCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

func (bs *BenchmarkSuite) benchmarkAVLRemove(n int) BenchmarkResult {
	tree := ds.NewAVLTree()
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}

	removeCount := n / 2
	start := time.Now()

	for i := 0; i < removeCount; i++ {
		tree.Remove(i)
	}

	duration := time.Since(start)

	return BenchmarkResult{
		Operation:     "Remove",
		DataStructure: "AVLTree",
		NumElements:   removeCount,
		Duration:      duration,
		OpsPerSecond:  float64(removeCount) / duration.Seconds(),
		MemoryUsed:    0,
	}
}

// ============================================================================
// SERIALIZATION BENCHMARKS
// ============================================================================

func (bs *BenchmarkSuite) benchmarkSerializationBinary(n int) []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	tmpDir := os.TempDir()

	// Array
	arr := ds.NewMyArray()
	for i := 0; i < n; i++ {
		arr.AddToEnd(i)
	}
	filename := tmpDir + "/array_bench.bin"
	start := time.Now()
	arr.Serialize(filename)
	serDur := time.Since(start)
	start = time.Now()
	arr.Deserialize(filename)
	desDur := time.Since(start)
	os.Remove(filename)
	results = append(results, BenchmarkResult{
		Operation:     "Binary Serialize",
		DataStructure: "MyArray",
		NumElements:   n,
		Duration:      serDur,
		OpsPerSecond:  float64(n) / serDur.Seconds(),
	})
	results = append(results, BenchmarkResult{
		Operation:     "Binary Deserialize",
		DataStructure: "MyArray",
		NumElements:   n,
		Duration:      desDur,
		OpsPerSecond:  float64(n) / desDur.Seconds(),
	})

	// AVL Tree
	tree := ds.NewAVLTree()
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}
	filename = tmpDir + "/avl_bench.bin"
	start = time.Now()
	tree.Serialize(filename)
	serDur = time.Since(start)
	start = time.Now()
	tree.Deserialize(filename)
	desDur = time.Since(start)
	os.Remove(filename)
	results = append(results, BenchmarkResult{
		Operation:     "Binary Serialize",
		DataStructure: "AVLTree",
		NumElements:   n,
		Duration:      serDur,
		OpsPerSecond:  float64(n) / serDur.Seconds(),
	})
	results = append(results, BenchmarkResult{
		Operation:     "Binary Deserialize",
		DataStructure: "AVLTree",
		NumElements:   n,
		Duration:      desDur,
		OpsPerSecond:  float64(n) / desDur.Seconds(),
	})

	// Hash Table Chain
	htc := ds.NewHashTableChain(n / 4)
	for i := 0; i < n; i++ {
		htc.Insert(i, i*2)
	}
	filename = tmpDir + "/htc_bench.bin"
	start = time.Now()
	htc.Serialize(filename)
	serDur = time.Since(start)
	start = time.Now()
	htc.Deserialize(filename)
	desDur = time.Since(start)
	os.Remove(filename)
	results = append(results, BenchmarkResult{
		Operation:     "Binary Serialize",
		DataStructure: "HashTableChain",
		NumElements:   n,
		Duration:      serDur,
		OpsPerSecond:  float64(n) / serDur.Seconds(),
	})
	results = append(results, BenchmarkResult{
		Operation:     "Binary Deserialize",
		DataStructure: "HashTableChain",
		NumElements:   n,
		Duration:      desDur,
		OpsPerSecond:  float64(n) / desDur.Seconds(),
	})

	return results
}

func (bs *BenchmarkSuite) benchmarkSerializationJSON(n int) []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	tmpDir := os.TempDir()

	// Array
	arr := ds.NewMyArray()
	for i := 0; i < n; i++ {
		arr.AddToEnd(i)
	}
	filename := tmpDir + "/array_bench.json"
	start := time.Now()
	arr.SerializeJSON(filename)
	serDur := time.Since(start)
	start = time.Now()
	arr.DeserializeJSON(filename)
	desDur := time.Since(start)
	os.Remove(filename)
	results = append(results, BenchmarkResult{
		Operation:     "JSON Serialize",
		DataStructure: "MyArray",
		NumElements:   n,
		Duration:      serDur,
		OpsPerSecond:  float64(n) / serDur.Seconds(),
	})
	results = append(results, BenchmarkResult{
		Operation:     "JSON Deserialize",
		DataStructure: "MyArray",
		NumElements:   n,
		Duration:      desDur,
		OpsPerSecond:  float64(n) / desDur.Seconds(),
	})

	// AVL Tree
	tree := ds.NewAVLTree()
	for i := 0; i < n; i++ {
		tree.Insert(i)
	}
	filename = tmpDir + "/avl_bench.json"
	start = time.Now()
	tree.SerializeJSON(filename)
	serDur = time.Since(start)
	start = time.Now()
	tree.DeserializeJSON(filename)
	desDur = time.Since(start)
	os.Remove(filename)
	results = append(results, BenchmarkResult{
		Operation:     "JSON Serialize",
		DataStructure: "AVLTree",
		NumElements:   n,
		Duration:      serDur,
		OpsPerSecond:  float64(n) / serDur.Seconds(),
	})
	results = append(results, BenchmarkResult{
		Operation:     "JSON Deserialize",
		DataStructure: "AVLTree",
		NumElements:   n,
		Duration:      desDur,
		OpsPerSecond:  float64(n) / desDur.Seconds(),
	})

	return results
}

// ============================================================================
// PRINTING AND DISPLAY
// ============================================================================

func printHeader() {
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║          DATA STRUCTURES BENCHMARK SUITE - Interactive Edition               ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printMenu() {
	fmt.Println("┌──────────────────────────────────────────────────────────────────────────────┐")
	fmt.Println("│                              BENCHMARK MENU                                  │")
	fmt.Println("├──────────────────────────────────────────────────────────────────────────────┤")
	fmt.Println("│  1.  MyArray                      5.  Queue                                  │")
	fmt.Println("│  2.  Singly Linked List           6.  Hash Table (Chaining)                  │")
	fmt.Println("│  3.  Doubly Linked List           7.  Hash Table (Open Addressing)           │")
	fmt.Println("│  4.  Stack                        8.  AVL Tree                               │")
	fmt.Println("├──────────────────────────────────────────────────────────────────────────────┤")
	fmt.Println("│  9.  Run ALL Benchmarks          10.  Serialization Comparison               │")
	fmt.Println("│ 11.  Compare Similar Operations  12.  Custom Size Benchmark                  │")
	fmt.Println("├──────────────────────────────────────────────────────────────────────────────┤")
	fmt.Println("│  0.  Exit                                                                    │")
	fmt.Println("└──────────────────────────────────────────────────────────────────────────────┘")
	fmt.Print("\nEnter choice: ")
}

func printResults(results []BenchmarkResult) {
	if len(results) == 0 {
		fmt.Println("No results to display.")
		return
	}

	fmt.Println()
	fmt.Println("┌────────────────────┬──────────────────┬────────────┬──────────────┬──────────────┬────────────┐")
	fmt.Println("│ Data Structure     │ Operation        │ Elements   │ Duration     │ Ops/Second   │ Memory     │")
	fmt.Println("├────────────────────┼──────────────────┼────────────┼──────────────┼──────────────┼────────────┤")

	for _, r := range results {
		memStr := "-"
		if r.MemoryUsed > 0 {
			memStr = formatMemory(r.MemoryUsed)
		}
		fmt.Printf("│ %-18s │ %-16s │ %10d │ %12s │ %12s │ %10s │\n",
			truncateString(r.DataStructure, 18),
			truncateString(r.Operation, 16),
			r.NumElements,
			formatDuration(r.Duration),
			formatOps(r.OpsPerSecond),
			memStr,
		)
	}

	fmt.Println("└────────────────────┴──────────────────┴────────────┴──────────────┴──────────────┴────────────┘")
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-2] + ".."
}

func printComparisonTable(results []BenchmarkResult, operation string) {
	fmt.Printf("\n╔══════════════════════════════════════════════════════════════════════════════╗\n")
	fmt.Printf("║  COMPARISON: %-64s ║\n", operation)
	fmt.Println("╠══════════════════════════════════════════════════════════════════════════════╣")

	// Group by size
	sizeMap := make(map[int][]BenchmarkResult)
	for _, r := range results {
		if strings.Contains(strings.ToLower(r.Operation), strings.ToLower(operation)) ||
			strings.ToLower(r.Operation) == strings.ToLower(operation) {
			sizeMap[r.NumElements] = append(sizeMap[r.NumElements], r)
		}
	}

	for size, sizeResults := range sizeMap {
		fmt.Printf("║  Size: %-71d ║\n", size)
		fmt.Println("╟──────────────────────────────────────────────────────────────────────────────╢")
		for _, r := range sizeResults {
			fmt.Printf("║  %-20s │ %12s │ %15s                    ║\n",
				r.DataStructure, formatDuration(r.Duration), formatOps(r.OpsPerSecond))
		}
		fmt.Println("╟──────────────────────────────────────────────────────────────────────────────╢")
	}

	fmt.Println("╚══════════════════════════════════════════════════════════════════════════════╝")
}

// ============================================================================
// MAIN BENCHMARK RUNNERS
// ============================================================================

func (bs *BenchmarkSuite) runArrayBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running MyArray benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkArrayInsertEnd(size))
		results = append(results, bs.benchmarkArrayRandomAccess(size))
	}

	// Middle insert and remove are expensive, use smaller size
	smallSize := bs.sizes[0]
	fmt.Printf("   Testing insert/remove operations with %d elements...\n", smallSize)
	results = append(results, bs.benchmarkArrayInsertMiddle(smallSize))
	results = append(results, bs.benchmarkArrayRemove(smallSize))

	return results
}

func (bs *BenchmarkSuite) runSLLBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Singly Linked List benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkSLLPushFront(size))
		results = append(results, bs.benchmarkSLLPushBack(size))
		results = append(results, bs.benchmarkSLLPopFront(size))
	}

	// Find is O(n), use smaller size
	smallSize := bs.sizes[0]
	fmt.Printf("   Testing find with %d elements...\n", smallSize)
	results = append(results, bs.benchmarkSLLFind(smallSize))

	return results
}

func (bs *BenchmarkSuite) runDLLBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Doubly Linked List benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkDLLPushFront(size))
		results = append(results, bs.benchmarkDLLPushBack(size))
		results = append(results, bs.benchmarkDLLPopBack(size))
	}

	smallSize := bs.sizes[0]
	fmt.Printf("   Testing find with %d elements...\n", smallSize)
	results = append(results, bs.benchmarkDLLFind(smallSize))

	return results
}

func (bs *BenchmarkSuite) runStackBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Stack benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkStackPush(size))
		results = append(results, bs.benchmarkStackPop(size))
		results = append(results, bs.benchmarkStackPeek(size))
	}

	return results
}

func (bs *BenchmarkSuite) runQueueBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Queue benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkQueuePush(size))
		results = append(results, bs.benchmarkQueuePop(size))
		results = append(results, bs.benchmarkQueuePeek(size))
	}

	return results
}

func (bs *BenchmarkSuite) runHashChainBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Hash Table (Chaining) benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkHashChainInsert(size))
		results = append(results, bs.benchmarkHashChainGet(size))
		results = append(results, bs.benchmarkHashChainRemove(size))
	}

	return results
}

func (bs *BenchmarkSuite) runHashOpenBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Hash Table (Open Addressing) benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkHashOpenInsert(size))
		results = append(results, bs.benchmarkHashOpenGet(size))
		results = append(results, bs.benchmarkHashOpenRemove(size))
	}

	return results
}

func (bs *BenchmarkSuite) runAVLBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running AVL Tree benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkAVLInsertRandom(size))
		results = append(results, bs.benchmarkAVLInsertSequential(size))
		results = append(results, bs.benchmarkAVLFind(size))
		results = append(results, bs.benchmarkAVLRemove(size))
	}

	return results
}

func (bs *BenchmarkSuite) runAllBenchmarks() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	results = append(results, bs.runArrayBenchmarks()...)
	results = append(results, bs.runSLLBenchmarks()...)
	results = append(results, bs.runDLLBenchmarks()...)
	results = append(results, bs.runStackBenchmarks()...)
	results = append(results, bs.runQueueBenchmarks()...)
	results = append(results, bs.runHashChainBenchmarks()...)
	results = append(results, bs.runHashOpenBenchmarks()...)
	results = append(results, bs.runAVLBenchmarks()...)
	return results
}

func (bs *BenchmarkSuite) runSerializationComparison() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	fmt.Println("\n🔄 Running Serialization benchmarks...")

	for _, size := range bs.sizes {
		fmt.Printf("   Testing with %d elements...\n", size)
		results = append(results, bs.benchmarkSerializationBinary(size)...)
		results = append(results, bs.benchmarkSerializationJSON(size)...)
	}

	return results
}

func (bs *BenchmarkSuite) runSimilarOperationsComparison() []BenchmarkResult {
	results := make([]BenchmarkResult, 0)
	size := 10000

	fmt.Println("\n🔄 Comparing similar operations across data structures...")
	fmt.Printf("   Using %d elements for fair comparison...\n", size)

	// Push/Insert operations
	fmt.Println("   Testing push/insert operations...")
	results = append(results, bs.benchmarkArrayInsertEnd(size))
	results = append(results, bs.benchmarkSLLPushBack(size))
	results = append(results, bs.benchmarkDLLPushBack(size))
	results = append(results, bs.benchmarkStackPush(size))
	results = append(results, bs.benchmarkQueuePush(size))
	results = append(results, bs.benchmarkHashChainInsert(size))
	results = append(results, bs.benchmarkHashOpenInsert(size))
	results = append(results, bs.benchmarkAVLInsertRandom(size))

	return results
}

func (bs *BenchmarkSuite) runCustomSizeBenchmark(size int) []BenchmarkResult {
	oldSizes := bs.sizes
	bs.sizes = []int{size}
	results := bs.runAllBenchmarks()
	bs.sizes = oldSizes
	return results
}

// ============================================================================
// MAIN
// ============================================================================

func main() {
	rand.Seed(time.Now().UnixNano())
	bs := NewBenchmarkSuite()
	reader := bufio.NewReader(os.Stdin)

	printHeader()

	for {
		printMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		var results []BenchmarkResult

		switch choice {
		case 0:
			fmt.Println("\n👋 Goodbye! Thanks for benchmarking!")
			return
		case 1:
			results = bs.runArrayBenchmarks()
		case 2:
			results = bs.runSLLBenchmarks()
		case 3:
			results = bs.runDLLBenchmarks()
		case 4:
			results = bs.runStackBenchmarks()
		case 5:
			results = bs.runQueueBenchmarks()
		case 6:
			results = bs.runHashChainBenchmarks()
		case 7:
			results = bs.runHashOpenBenchmarks()
		case 8:
			results = bs.runAVLBenchmarks()
		case 9:
			results = bs.runAllBenchmarks()
		case 10:
			results = bs.runSerializationComparison()
		case 11:
			results = bs.runSimilarOperationsComparison()
			printResults(results)
			fmt.Println("\n📊 Performance Summary:")
			printComparisonTable(results, "push")
			printComparisonTable(results, "insert")
			continue
		case 12:
			fmt.Print("Enter custom size (e.g., 50000): ")
			sizeInput, _ := reader.ReadString('\n')
			sizeInput = strings.TrimSpace(sizeInput)
			customSize, err := strconv.Atoi(sizeInput)
			if err != nil || customSize <= 0 {
				fmt.Println("Invalid size. Using default 10000.")
				customSize = 10000
			}
			results = bs.runCustomSizeBenchmark(customSize)
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		if len(results) > 0 {
			printResults(results)
			bs.results = append(bs.results, results...)
		}

		fmt.Println("\nPress Enter to continue...")
		reader.ReadString('\n')
	}
}
