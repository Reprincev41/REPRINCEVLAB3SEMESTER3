#define CATCH_CONFIG_MAIN
#include "catch.hpp"

#include <cstdio>
#include <filesystem>
#include <sstream>
#include <iostream>
#include "AVLTree.h"
#include "Stack.h"
#include "Queue.h"
#include "Array.h"
#include "DoublyList.h"
#include "SinglyList.h"
#include "HashTableOpen.h"
#include "HashTableChain.h"

TEST_CASE("MyArray", "[array]") {
    MyArray arr;
    
    SECTION("Empty array") {
        REQUIRE(arr.get_length() == 0);
        REQUIRE_THROWS_AS(arr.get_at_index(0), std::out_of_range);
    }
    
    SECTION("Add and get elements") {
        arr.add_to_end(10);
        arr.add_to_end(20);
        arr.add_to_end(30);
        
        REQUIRE(arr.get_length() == 3);
        REQUIRE(arr.get_at_index(0) == 10);
        REQUIRE(arr.get_at_index(1) == 20);
        REQUIRE(arr.get_at_index(2) == 30);
        
        arr.add_at_index(1, 15);
        REQUIRE(arr.get_length() == 4);
        REQUIRE(arr.get_at_index(1) == 15);
        REQUIRE(arr.get_at_index(2) == 20);
    }
    
    SECTION("Remove elements") {
        arr.add_to_end(10);
        arr.add_to_end(20);
        arr.add_to_end(30);
        arr.add_to_end(40);
        
        arr.remove_at_index(1);
        REQUIRE(arr.get_length() == 3);
        REQUIRE(arr.get_at_index(0) == 10);
        REQUIRE(arr.get_at_index(1) == 30);
        
        arr.remove_at_index(2);
        REQUIRE(arr.get_length() == 2);
        REQUIRE(arr.get_at_index(1) == 30);
    }
    
    SECTION("Replace elements") {
        arr.add_to_end(10);
        arr.add_to_end(20);
        arr.add_to_end(30);
        
        arr.replace_at_index(1, 25);
        REQUIRE(arr.get_at_index(1) == 25);
        
        REQUIRE_THROWS_AS(arr.replace_at_index(5, 100), std::out_of_range);
    }
    
    SECTION("Boundary conditions") {
        // Test resize
        for (int i = 0; i < 100; i++) {
            arr.add_to_end(i);
        }
        REQUIRE(arr.get_length() == 100);
        REQUIRE(arr.get_at_index(99) == 99);
        
        // Test out of bounds
        REQUIRE_THROWS_AS(arr.get_at_index(100), std::out_of_range);
        REQUIRE_THROWS_AS(arr.add_at_index(101, 0), std::out_of_range);
    }
    
    SECTION("Serialization") {
        // Prepare test data
        for (int i = 0; i < 10; i++) {
            arr.add_to_end(i * 10);
        }
        
        // Serialize
        const std::string filename = "test_array.bin";
        arr.serialize(filename);
        
        // Deserialize to new array
        MyArray arr2;
        arr2.deserialize(filename);
        
        // Verify
        REQUIRE(arr2.get_length() == arr.get_length());
        for (size_t i = 0; i < arr.get_length(); i++) {
            REQUIRE(arr2.get_at_index(i) == arr.get_at_index(i));
        }
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("SinglyLinkedList", "[singlylist]") {
    SinglyLinkedList list;
    
    SECTION("Empty list") {
        REQUIRE(list.get_size() == 0);
        REQUIRE_FALSE(list.find(10));
    }
    
    SECTION("Push and pop operations") {
        list.push_front(30);
        list.push_front(20);
        list.push_front(10);
        
        REQUIRE(list.get_size() == 3);
        
        list.push_back(40);
        REQUIRE(list.get_size() == 4);
        
        list.pop_front();
        REQUIRE(list.get_size() == 3);
        
        list.pop_back();
        REQUIRE(list.get_size() == 2);
    }
    
    SECTION("Insert operations") {
        list.push_back(10);
        list.push_back(30);
        
        list.insert_after(0, 20);
        REQUIRE(list.get_size() == 3);
        
        list.insert_before(0, 5);
        REQUIRE(list.get_size() == 4);
    }
    
    SECTION("Remove operations") {
        list.push_back(10);
        list.push_back(20);
        list.push_back(30);
        list.push_back(20); // Duplicate
        
        list.remove_at(1);
        REQUIRE(list.get_size() == 3);
        
        list.remove_by_value(20);
        REQUIRE(list.get_size() == 2);
        
        // Remove non-existent value
        list.remove_by_value(100);
        REQUIRE(list.get_size() == 2);
    }
    
    SECTION("Find operations") {
        list.push_back(10);
        list.push_back(20);
        list.push_back(30);
        
        REQUIRE(list.find(20));
        REQUIRE_FALSE(list.find(40));
    }
    
    SECTION("Boundary conditions") {
        // Insert/remove at invalid indices
        REQUIRE_THROWS_AS(list.insert_after(0, 10), std::out_of_range);
        REQUIRE_THROWS_AS(list.remove_at(0), std::out_of_range);
        
        // Pop from empty list
        list.pop_front(); // Should not throw
        list.pop_back(); // Should not throw
        
        // Large list
        for (int i = 0; i < 1000; i++) {
            list.push_back(i);
        }
        REQUIRE(list.get_size() == 1000);
    }
    
    SECTION("Serialization") {
        for (int i = 0; i < 10; i++) {
            list.push_back(i * 5);
        }
        
        const std::string filename = "test_singlylist.bin";
        list.serialize(filename);
        
        SinglyLinkedList list2;
        list2.deserialize(filename);
        
        REQUIRE(list2.get_size() == list.get_size());
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("DoublyLinkedList", "[doublylist]") {
    DoublyLinkedList list;
    
    SECTION("Empty list") {
        REQUIRE(list.get_size() == 0);
        REQUIRE_FALSE(list.find(10));
    }
    
    SECTION("Push and pop operations") {
        list.push_front(30);
        list.push_front(20);
        list.push_front(10);
        
        REQUIRE(list.get_size() == 3);
        
        list.push_back(40);
        REQUIRE(list.get_size() == 4);
        
        list.pop_front();
        REQUIRE(list.get_size() == 3);
        
        list.pop_back();
        REQUIRE(list.get_size() == 2);
    }
    
    SECTION("Insert operations") {
        list.push_back(10);
        list.push_back(30);
        
        list.insert_after(0, 20);
        REQUIRE(list.get_size() == 3);
        
        list.insert_before(0, 5);
        REQUIRE(list.get_size() == 4);
        
        list.insert_before(3, 25);
        REQUIRE(list.get_size() == 5);
    }
    
    SECTION("Remove operations") {
        list.push_back(10);
        list.push_back(20);
        list.push_back(30);
        list.push_back(40);
        
        list.remove_at(1);
        REQUIRE(list.get_size() == 3);
        REQUIRE(list.find(20) == false);
        
        list.remove_by_value(30);
        REQUIRE(list.get_size() == 2);
        REQUIRE(list.find(30) == false);
        
        // Remove head
        list.remove_by_value(10);
        REQUIRE(list.get_size() == 1);
        
        // Remove tail
        list.remove_by_value(40);
        REQUIRE(list.get_size() == 0);
    }
    
    SECTION("Find operations") {
        list.push_back(10);
        list.push_back(20);
        list.push_back(30);
        
        REQUIRE(list.find(20));
        REQUIRE_FALSE(list.find(40));
    }
    
    SECTION("Serialization") {
        for (int i = 0; i < 10; i++) {
            list.push_back(i * 3);
        }
        
        const std::string filename = "test_doublylist.bin";
        list.serialize(filename);
        
        DoublyLinkedList list2;
        list2.deserialize(filename);
        
        REQUIRE(list2.get_size() == list.get_size());
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("MyStack", "[stack]") {
    MyStack stack;
    
    SECTION("Empty stack") {
        REQUIRE_THROWS_AS(stack.peek(), std::runtime_error);
    }
    
    SECTION("Push and pop") {
        stack.push(10);
        stack.push(20);
        stack.push(30);
        
        REQUIRE(stack.peek() == 30);
        
        stack.pop();
        REQUIRE(stack.peek() == 20);
        
        stack.pop();
        REQUIRE(stack.peek() == 10);
        
        stack.pop();
        REQUIRE_THROWS_AS(stack.peek(), std::runtime_error);
    }
    
    SECTION("Multiple operations") {
        for (int i = 0; i < 100; i++) {
            stack.push(i);
        }
        
        for (int i = 99; i >= 0; i--) {
            REQUIRE(stack.peek() == i);
            stack.pop();
        }
        
        REQUIRE_THROWS_AS(stack.peek(), std::runtime_error);
    }
    
    SECTION("Serialization") {
        for (int i = 0; i < 10; i++) {
            stack.push(i);
        }
        
        const std::string filename = "test_stack.bin";
        stack.serialize(filename);
        
        MyStack stack2;
        stack2.deserialize(filename);
        
        // Compare stacks by popping all elements
        for (int i = 9; i >= 0; i--) {
            REQUIRE(stack2.peek() == i);
            stack2.pop();
        }
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("MyQueue", "[queue]") {
    MyQueue queue;
    
    SECTION("Empty queue") {
        REQUIRE_THROWS_AS(queue.peek(), std::runtime_error);
    }
    
    SECTION("Push and pop") {
        queue.push(10);
        queue.push(20);
        queue.push(30);
        
        REQUIRE(queue.peek() == 10);
        
        queue.pop();
        REQUIRE(queue.peek() == 20);
        
        queue.pop();
        REQUIRE(queue.peek() == 30);
        
        queue.pop();
        REQUIRE_THROWS_AS(queue.peek(), std::runtime_error);
    }
    
    SECTION("FIFO order") {
        for (int i = 0; i < 100; i++) {
            queue.push(i);
        }
        
        for (int i = 0; i < 100; i++) {
            REQUIRE(queue.peek() == i);
            queue.pop();
        }
        
        REQUIRE_THROWS_AS(queue.peek(), std::runtime_error);
    }
    
    SECTION("Serialization") {
        for (int i = 0; i < 10; i++) {
            queue.push(i);
        }
        
        const std::string filename = "test_queue.bin";
        queue.serialize(filename);
        
        MyQueue queue2;
        queue2.deserialize(filename);
        
        // Compare queues
        for (int i = 0; i < 10; i++) {
            REQUIRE(queue2.peek() == i);
            queue2.pop();
        }
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("AVLTree", "[avltree]") {
    AVLTree tree;
    
    SECTION("Empty tree") {
        REQUIRE_FALSE(tree.find(10));
    }
    
    SECTION("Insert and find") {
        tree.insert(50);
        tree.insert(30);
        tree.insert(70);
        tree.insert(20);
        tree.insert(40);
        tree.insert(60);
        tree.insert(80);
        
        REQUIRE(tree.find(50));
        REQUIRE(tree.find(30));
        REQUIRE(tree.find(70));
        REQUIRE(tree.find(20));
        REQUIRE(tree.find(40));
        REQUIRE(tree.find(60));
        REQUIRE(tree.find(80));
        REQUIRE_FALSE(tree.find(100));
    }
    
    SECTION("Remove operations") {
        tree.insert(50);
        tree.insert(30);
        tree.insert(70);
        tree.insert(20);
        tree.insert(40);
        
        REQUIRE(tree.find(30));
        tree.remove(30);
        REQUIRE_FALSE(tree.find(30));
        
        // Remove root
        REQUIRE(tree.find(50));
        tree.remove(50);
        REQUIRE_FALSE(tree.find(50));
        
        // Remove non-existent
        tree.remove(100); // Should not crash
    }
    
    SECTION("Rotation cases") {
        // Test LL rotation
        tree.insert(30);
        tree.insert(20);
        tree.insert(10);
        
        // Test RR rotation
        AVLTree tree2;
        tree2.insert(10);
        tree2.insert(20);
        tree2.insert(30);
        
        // Test LR rotation
        AVLTree tree3;
        tree3.insert(30);
        tree3.insert(10);
        tree3.insert(20);
        
        // Test RL rotation
        AVLTree tree4;
        tree4.insert(10);
        tree4.insert(30);
        tree4.insert(20);
    }
    
    SECTION("Serialization") {
        int values[] = {50, 30, 70, 20, 40, 60, 80};
        for (int val : values) {
            tree.insert(val);
        }
        
        const std::string filename = "test_avltree.bin";
        tree.serialize(filename);
        
        AVLTree tree2;
        tree2.deserialize(filename);
        
        // Verify all values exist
        for (int val : values) {
            REQUIRE(tree2.find(val));
        }
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("HashTableOpen", "[hashtableopen]") {
    HashTableOpen ht(4);
    
    SECTION("Empty table") {
        REQUIRE_FALSE(ht.get(10).has_value());
    }
    
    SECTION("Insert and get") {
        ht.insert(1, 100);
        ht.insert(2, 200);
        ht.insert(3, 300);
        
        REQUIRE(ht.get(1).value() == 100);
        REQUIRE(ht.get(2).value() == 200);
        REQUIRE(ht.get(3).value() == 300);
        REQUIRE_FALSE(ht.get(4).has_value());
    }
    
    SECTION("Update value") {
        ht.insert(1, 100);
        ht.insert(1, 150); // Update
        
        REQUIRE(ht.get(1).value() == 150);
    }
    
    SECTION("Remove operations") {
        ht.insert(1, 100);
        ht.insert(2, 200);
        ht.insert(3, 300);
        
        REQUIRE(ht.get(1).has_value());
        ht.remove(1);
        REQUIRE_FALSE(ht.get(1).has_value());
        
        // Remove non-existent
        ht.remove(100); // Should not crash
        
        // Check other values still exist
        REQUIRE(ht.get(2).has_value());
        REQUIRE(ht.get(3).has_value());
    }
    
    SECTION("Collision handling") {
        // Force collisions by using small capacity
        HashTableOpen smallTable(3);
        smallTable.insert(1, 100);  // hash(1) = 1
        smallTable.insert(4, 400);  // hash(4) = 1 (collision)
        smallTable.insert(7, 700);  // hash(7) = 1 (collision)
        
        REQUIRE(smallTable.get(1).value() == 100);
        REQUIRE(smallTable.get(4).value() == 400);
        REQUIRE(smallTable.get(7).value() == 700);
    }
    
    SECTION("Resize operation") {
        // Insert enough items to trigger resize
        for (int i = 0; i < 10; i++) {
            ht.insert(i, i * 100);
        }
        
        // All values should still be accessible
        for (int i = 0; i < 10; i++) {
            REQUIRE(ht.get(i).value() == i * 100);
        }
    }
    
    SECTION("Serialization") {
        for (int i = 0; i < 10; i++) {
            ht.insert(i, i * 10);
        }
        
        const std::string filename = "test_hashtableopen.bin";
        ht.serialize(filename);
        
        HashTableOpen ht2;
        ht2.deserialize(filename);
        
        // Verify all values
        for (int i = 0; i < 10; i++) {
            REQUIRE(ht2.get(i).value() == i * 10);
        }
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("HashTableChain", "[hashtablechain]") {
    HashTableChain ht(4);
    
    SECTION("Empty table") {
        REQUIRE_FALSE(ht.get(10).has_value());
    }
    
    SECTION("Insert and get") {
        ht.insert(1, 100);
        ht.insert(2, 200);
        ht.insert(3, 300);
        
        REQUIRE(ht.get(1).value() == 100);
        REQUIRE(ht.get(2).value() == 200);
        REQUIRE(ht.get(3).value() == 300);
        REQUIRE_FALSE(ht.get(4).has_value());
    }
    
    SECTION("Update value") {
        ht.insert(1, 100);
        ht.insert(1, 150); // Chain adds new node
        
        // Both nodes exist in chain, but get returns first
        REQUIRE(ht.get(1).value() == 150);
    }
    
    SECTION("Remove operations") {
        ht.insert(1, 100);
        ht.insert(2, 200);
        ht.insert(3, 300);
        
        REQUIRE(ht.get(1).has_value());
        ht.remove(1);
        REQUIRE_FALSE(ht.get(1).has_value());
        
        // Remove non-existent
        ht.remove(100); // Should not crash
        
        // Check other values still exist
        REQUIRE(ht.get(2).has_value());
        REQUIRE(ht.get(3).has_value());
    }
    
    SECTION("Collision chains") {
        // Force collisions
        HashTableChain smallTable(3);
        smallTable.insert(1, 100);  // hash(1) = 1
        smallTable.insert(4, 400);  // hash(4) = 1 (collision)
        smallTable.insert(7, 700);  // hash(7) = 1 (collision)
        
        REQUIRE(smallTable.get(1).value() == 100);
        REQUIRE(smallTable.get(4).value() == 400);
        REQUIRE(smallTable.get(7).value() == 700);
        
        // Remove from middle of chain
        smallTable.remove(4);
        REQUIRE_FALSE(smallTable.get(4).has_value());
        REQUIRE(smallTable.get(1).has_value());
        REQUIRE(smallTable.get(7).has_value());
    }
    
    SECTION("Serialization") {
        for (int i = 0; i < 10; i++) {
            ht.insert(i, i * 10);
        }
        
        const std::string filename = "test_hashtablechain.bin";
        ht.serialize(filename);
        
        HashTableChain ht2;
        ht2.deserialize(filename);
        
        // Verify all values
        for (int i = 0; i < 10; i++) {
            REQUIRE(ht2.get(i).value() == i * 10);
        }
        
        // Cleanup
        std::remove(filename.c_str());
    }
}

TEST_CASE("Stack print") {
    MyStack s;
    for (int i = 1; i <= 3; ++i) s.push(i);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    s.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}

TEST_CASE("Queue print") {
    MyQueue q;
    for (int i = 1; i <= 3; ++i) q.push(i);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    q.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}

TEST_CASE("Array print") {
    MyArray arr;
    arr.add_at_index(0, 10);
    arr.add_at_index(1, 20);
    arr.add_at_index(2, 30);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    arr.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}

TEST_CASE("DoublyList print") {
    DoublyLinkedList list;
    list.push_back(1);
    list.push_back(2);
    list.push_back(3);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    list.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}

TEST_CASE("SinglyList print") {
    SinglyLinkedList list;
    list.push_back(1);
    list.push_back(2);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    list.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}

TEST_CASE("HashTableOpen print") {
    HashTableOpen ht;
    ht.insert(1, 10);
    ht.insert(2, 20);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    ht.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}

TEST_CASE("HashTableChain print") {
    HashTableChain ht;
    ht.insert(1, 10);
    ht.insert(2, 20);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    ht.print();
    std::cout.rdbuf(old);

    REQUIRE_FALSE(oss.str().empty());
}
