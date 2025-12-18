#define BOOST_TEST_MODULE DataStructuresTest
#include <boost/test/included/unit_test.hpp>
#include <boost/test/tools/output_test_stream.hpp>
#include <filesystem>
#include <random>

#include "Stack.h"
#include "Queue.h"
#include "Array.h"
#include "DoublyList.h"
#include "SinglyList.h"
#include "HashTableOpen.h"
#include "HashTableChain.h"
#include "AVLTree.h"

#include <sstream>
#include <iostream>

namespace fs = std::filesystem;

// Вспомогательные функции
std::string generate_temp_filename() {
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<> dis(10000, 99999);
    return "test_temp_" + std::to_string(dis(gen)) + ".bin";
}

// ==========================================
// Тесты для MyArray
// ==========================================
BOOST_AUTO_TEST_SUITE(ArrayTestSuite)

BOOST_AUTO_TEST_CASE(Array_Constructor_Destructor) {
    MyArray arr;
    BOOST_TEST(arr.get_length() == 0);
}

BOOST_AUTO_TEST_CASE(Array_AddToEnd_GetAtIndex) {
    MyArray arr;
    
    arr.add_to_end(10);
    arr.add_to_end(20);
    arr.add_to_end(30);
    
    BOOST_TEST(arr.get_length() == 3);
    BOOST_TEST(arr.get_at_index(0) == 10);
    BOOST_TEST(arr.get_at_index(1) == 20);
    BOOST_TEST(arr.get_at_index(2) == 30);
    
    // Проверка исключения при неверном индексе
    BOOST_CHECK_THROW(arr.get_at_index(5), std::out_of_range);
}

BOOST_AUTO_TEST_CASE(Array_AddAtIndex) {
    MyArray arr;
    
    arr.add_to_end(10);
    arr.add_to_end(30);
    arr.add_at_index(1, 20);
    
    BOOST_TEST(arr.get_length() == 3);
    BOOST_TEST(arr.get_at_index(0) == 10);
    BOOST_TEST(arr.get_at_index(1) == 20);
    BOOST_TEST(arr.get_at_index(2) == 30);
    
    // Добавление в начало
    arr.add_at_index(0, 5);
    BOOST_TEST(arr.get_at_index(0) == 5);
    BOOST_TEST(arr.get_length() == 4);
    
    // Добавление в конец
    arr.add_at_index(4, 40);
    BOOST_TEST(arr.get_at_index(4) == 40);
    BOOST_TEST(arr.get_length() == 5);
}

BOOST_AUTO_TEST_CASE(Array_RemoveAtIndex) {
    MyArray arr;
    
    for (int i = 0; i < 5; ++i) {
        arr.add_to_end(i * 10);
    }
    
    arr.remove_at_index(2); // Удаляем 20
    
    BOOST_TEST(arr.get_length() == 4);
    BOOST_TEST(arr.get_at_index(0) == 0);
    BOOST_TEST(arr.get_at_index(1) == 10);
    BOOST_TEST(arr.get_at_index(2) == 30);
    BOOST_TEST(arr.get_at_index(3) == 40);
    
    // Удаление из начала
    arr.remove_at_index(0);
    BOOST_TEST(arr.get_at_index(0) == 10);
    
    // Удаление из конца
    arr.remove_at_index(2);
    BOOST_TEST(arr.get_length() == 2);
    
    BOOST_CHECK_THROW(arr.remove_at_index(5), std::out_of_range);
}

BOOST_AUTO_TEST_CASE(Array_ReplaceAtIndex) {
    MyArray arr;
    
    arr.add_to_end(10);
    arr.add_to_end(20);
    arr.add_to_end(30);
    
    arr.replace_at_index(1, 99);
    
    BOOST_TEST(arr.get_at_index(1) == 99);
    BOOST_CHECK_THROW(arr.replace_at_index(5, 100), std::out_of_range);
}

BOOST_AUTO_TEST_CASE(Array_Resize) {
    MyArray arr;
    
    // Добавляем больше элементов, чем начальная ёмкость (2)
    for (int i = 0; i < 10; ++i) {
        arr.add_to_end(i);
    }
    
    BOOST_TEST(arr.get_length() == 10);
    for (int i = 0; i < 10; ++i) {
        BOOST_TEST(arr.get_at_index(i) == i);
    }
}

BOOST_AUTO_TEST_CASE(Array_Serialization) {
    MyArray arr1;
    
    for (int i = 0; i < 5; ++i) {
        arr1.add_to_end(i * 10);
    }
    
    std::string filename = generate_temp_filename();
    
    try {
        // Сериализация
        arr1.serialize(filename);
        
        // Десериализация
        MyArray arr2;
        arr2.deserialize(filename);
        
        BOOST_TEST(arr2.get_length() == arr1.get_length());
        for (size_t i = 0; i < arr1.get_length(); ++i) {
            BOOST_TEST(arr2.get_at_index(i) == arr1.get_at_index(i));
        }
        
        // Тест с несуществующим файлом
        BOOST_CHECK_THROW(arr2.deserialize("nonexistent_file.bin"), std::runtime_error);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для SinglyLinkedList
// ==========================================
BOOST_AUTO_TEST_SUITE(SinglyLinkedListTestSuite)

BOOST_AUTO_TEST_CASE(SinglyList_BasicOperations) {
    SinglyLinkedList list;
    
    BOOST_TEST(list.get_size() == 0);
    BOOST_TEST(list.find(10) == false);
}

BOOST_AUTO_TEST_CASE(SinglyList_PushFront_PopFront) {
    SinglyLinkedList list;
    
    list.push_front(30);
    list.push_front(20);
    list.push_front(10);
    
    BOOST_TEST(list.get_size() == 3);
    
    list.pop_front();
    BOOST_TEST(list.get_size() == 2);
    
    list.pop_front();
    BOOST_TEST(list.get_size() == 1);
    
    list.pop_front();
    BOOST_TEST(list.get_size() == 0);
    
    // Попытка удаления из пустого списка
    list.pop_front();
    BOOST_TEST(list.get_size() == 0);
}

BOOST_AUTO_TEST_CASE(SinglyList_PushBack_PopBack) {
    SinglyLinkedList list;
    
    list.push_back(10);
    list.push_back(20);
    list.push_back(30);
    
    BOOST_TEST(list.get_size() == 3);
    
    list.pop_back();
    BOOST_TEST(list.get_size() == 2);
    
    list.pop_back();
    BOOST_TEST(list.get_size() == 1);
    
    list.pop_back();
    BOOST_TEST(list.get_size() == 0);
    
    list.pop_back(); // Должно безопасно обрабатывать
}

BOOST_AUTO_TEST_CASE(SinglyList_InsertAfter) {
    SinglyLinkedList list;
    
    list.push_back(10);
    list.push_back(30);
    
    list.insert_after(0, 20); // Вставляем после 10
    
    BOOST_TEST(list.get_size() == 3);
    
    BOOST_CHECK_THROW(list.insert_after(5, 100), std::out_of_range);
}

BOOST_AUTO_TEST_CASE(SinglyList_InsertBefore) {
    SinglyLinkedList list;
    
    list.push_back(10);
    list.push_back(30);
    
    list.insert_before(1, 20); // Вставляем перед 30
    
    BOOST_TEST(list.get_size() == 3);
    
    list.insert_before(0, 5); // Вставка в начало
    BOOST_TEST(list.get_size() == 4);
}

BOOST_AUTO_TEST_CASE(SinglyList_RemoveAt) {
    SinglyLinkedList list;
    
    for (int i = 0; i < 5; ++i) {
        list.push_back(i * 10);
    }
    
    list.remove_at(2); // Удаляем 20
    BOOST_TEST(list.get_size() == 4);
    
    list.remove_at(0); // Удаляем начало
    BOOST_TEST(list.get_size() == 3);
    
    list.remove_at(2); // Удаляем конец
    BOOST_TEST(list.get_size() == 2);
    
    BOOST_CHECK_THROW(list.remove_at(5), std::out_of_range);
}

BOOST_AUTO_TEST_CASE(SinglyList_RemoveByValue) {
    SinglyLinkedList list;
    
    list.push_back(10);
    list.push_back(20);
    list.push_back(30);
    list.push_back(20); // Дубликат
    
    list.remove_by_value(20); // Удаляет первый 20
    BOOST_TEST(list.get_size() == 3);
    
    list.remove_by_value(99); // Несуществующее значение
    BOOST_TEST(list.get_size() == 3);
    
    list.remove_by_value(10); // Удаление из начала
    BOOST_TEST(list.get_size() == 2);
}

BOOST_AUTO_TEST_CASE(SinglyList_Serialization) {
    SinglyLinkedList list1;
    
    for (int i = 0; i < 5; ++i) {
        list1.push_back(i * 10);
    }
    
    std::string filename = generate_temp_filename();
    
    try {
        list1.serialize(filename);
        
        SinglyLinkedList list2;
        list2.deserialize(filename);
        
        BOOST_TEST(list2.get_size() == list1.get_size());
        
        // Проверка через find (поскольку нет прямого доступа по индексу)
        BOOST_TEST(list2.find(0) == true);
        BOOST_TEST(list2.find(10) == true);
        BOOST_TEST(list2.find(40) == true);
        BOOST_TEST(list2.find(99) == false);
        
        // Тест с повреждённым файлом (слишком большой размер)
        std::ofstream badfile("bad_file.bin", std::ios::binary);
        size_t huge_size = 1000000000;
        badfile.write(reinterpret_cast<const char*>(&huge_size), sizeof(huge_size));
        badfile.close();
        
        BOOST_CHECK_THROW(list2.deserialize("bad_file.bin"), std::runtime_error);
        fs::remove("bad_file.bin");
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для DoublyLinkedList
// ==========================================
BOOST_AUTO_TEST_SUITE(DoublyLinkedListTestSuite)

BOOST_AUTO_TEST_CASE(DoublyList_BasicOperations) {
    DoublyLinkedList list;
    
    BOOST_TEST(list.get_size() == 0);
    BOOST_TEST(list.find(10) == false);
}

BOOST_AUTO_TEST_CASE(DoublyList_PushPopOperations) {
    DoublyLinkedList list;
    
    list.push_front(30);
    list.push_front(20);
    list.push_front(10);
    
    BOOST_TEST(list.get_size() == 3);
    
    list.push_back(40);
    list.push_back(50);
    
    BOOST_TEST(list.get_size() == 5);
    
    list.pop_front();
    BOOST_TEST(list.get_size() == 4);
    
    list.pop_back();
    BOOST_TEST(list.get_size() == 3);
}

BOOST_AUTO_TEST_CASE(DoublyList_InsertOperations) {
    DoublyLinkedList list;
    
    list.push_back(10);
    list.push_back(30);
    
    list.insert_after(0, 20);
    BOOST_TEST(list.get_size() == 3);
    
    list.insert_before(2, 25);
    BOOST_TEST(list.get_size() == 4);
    
    BOOST_CHECK_THROW(list.insert_after(10, 100), std::out_of_range);
}

BOOST_AUTO_TEST_CASE(DoublyList_RemoveOperations) {
    DoublyLinkedList list;
    
    for (int i = 0; i < 5; ++i) {
        list.push_back(i * 10);
    }
    
    list.remove_at(2);
    BOOST_TEST(list.get_size() == 4);
    
    list.remove_by_value(30);
    BOOST_TEST(list.get_size() == 3);
    
    list.remove_by_value(0); // Удаление из начала
    BOOST_TEST(list.get_size() == 2);
    
    list.remove_by_value(40); // Удаление из конца
    BOOST_TEST(list.get_size() == 1);
}

BOOST_AUTO_TEST_CASE(DoublyList_Serialization) {
    DoublyLinkedList list1;
    
    list1.push_back(100);
    list1.push_back(200);
    list1.push_back(300);
    
    std::string filename = generate_temp_filename();
    
    try {
        list1.serialize(filename);
        
        DoublyLinkedList list2;
        list2.deserialize(filename);
        
        BOOST_TEST(list2.get_size() == list1.get_size());
        BOOST_TEST(list2.find(100) == true);
        BOOST_TEST(list2.find(200) == true);
        BOOST_TEST(list2.find(300) == true);
        BOOST_TEST(list2.find(999) == false);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для MyStack
// ==========================================
BOOST_AUTO_TEST_SUITE(StackTestSuite)

BOOST_AUTO_TEST_CASE(Stack_BasicOperations) {
    MyStack stack;
    
    BOOST_CHECK_THROW(stack.peek(), std::runtime_error);
}

BOOST_AUTO_TEST_CASE(Stack_PushPopPeek) {
    MyStack stack;
    
    stack.push(10);
    stack.push(20);
    stack.push(30);
    
    BOOST_TEST(stack.peek() == 30);
    
    stack.pop();
    BOOST_TEST(stack.peek() == 20);
    
    stack.pop();
    BOOST_TEST(stack.peek() == 10);
    
    stack.pop();
    BOOST_CHECK_THROW(stack.peek(), std::runtime_error);
    
    // Попытка удаления из пустого стека
    stack.pop(); // Должно безопасно обрабатывать
}

BOOST_AUTO_TEST_CASE(Stack_Serialization) {
    MyStack stack1;
    
    stack1.push(100);
    stack1.push(200);
    stack1.push(300);
    
    std::string filename = generate_temp_filename();
    
    try {
        stack1.serialize(filename);
        
        MyStack stack2;
        stack2.deserialize(filename);
        
        BOOST_TEST(stack2.peek() == 300);
        stack2.pop();
        BOOST_TEST(stack2.peek() == 200);
        stack2.pop();
        BOOST_TEST(stack2.peek() == 100);
        stack2.pop();
        BOOST_CHECK_THROW(stack2.peek(), std::runtime_error);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для MyQueue
// ==========================================
BOOST_AUTO_TEST_SUITE(QueueTestSuite)

BOOST_AUTO_TEST_CASE(Queue_BasicOperations) {
    MyQueue queue;
    
    BOOST_CHECK_THROW(queue.peek(), std::runtime_error);
}

BOOST_AUTO_TEST_CASE(Queue_PushPopPeek) {
    MyQueue queue;
    
    queue.push(10);
    queue.push(20);
    queue.push(30);
    
    BOOST_TEST(queue.peek() == 10);
    
    queue.pop();
    BOOST_TEST(queue.peek() == 20);
    
    queue.pop();
    BOOST_TEST(queue.peek() == 30);
    
    queue.pop();
    BOOST_CHECK_THROW(queue.peek(), std::runtime_error);
    
    queue.pop(); // Должно безопасно обрабатывать
}

BOOST_AUTO_TEST_CASE(Queue_Serialization) {
    MyQueue queue1;
    
    queue1.push(100);
    queue1.push(200);
    queue1.push(300);
    
    std::string filename = generate_temp_filename();
    
    try {
        queue1.serialize(filename);
        
        MyQueue queue2;
        queue2.deserialize(filename);
        
        BOOST_TEST(queue2.peek() == 100);
        queue2.pop();
        BOOST_TEST(queue2.peek() == 200);
        queue2.pop();
        BOOST_TEST(queue2.peek() == 300);
        queue2.pop();
        BOOST_CHECK_THROW(queue2.peek(), std::runtime_error);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для AVLTree
// ==========================================
BOOST_AUTO_TEST_SUITE(AVLTreeTestSuite)

BOOST_AUTO_TEST_CASE(AVLTree_Insert_Find) {
    AVLTree tree;
    
    BOOST_TEST(tree.find(10) == false);
    
    tree.insert(10);
    tree.insert(20);
    tree.insert(5);
    tree.insert(15);
    tree.insert(25);
    
    BOOST_TEST(tree.find(10) == true);
    BOOST_TEST(tree.find(20) == true);
    BOOST_TEST(tree.find(5) == true);
    BOOST_TEST(tree.find(15) == true);
    BOOST_TEST(tree.find(25) == true);
    BOOST_TEST(tree.find(99) == false);
    
    // Вставка дубликата
    tree.insert(10);
    BOOST_TEST(tree.find(10) == true);
}

BOOST_AUTO_TEST_CASE(AVLTree_Remove) {
    AVLTree tree;
    
    // Создаём сбалансированное дерево
    for (int i = 0; i < 10; ++i) {
        tree.insert(i * 10);
    }
    
    BOOST_TEST(tree.find(30) == true);
    tree.remove(30);
    BOOST_TEST(tree.find(30) == false);
    
    // Удаление корня
    tree.remove(0);
    BOOST_TEST(tree.find(0) == false);
    
    // Удаление несуществующего элемента
    tree.remove(999);
    // Не должно вызывать ошибку
    
    // Удаление всех элементов
    for (int i = 0; i < 10; ++i) {
        tree.remove(i * 10);
    }
}

BOOST_AUTO_TEST_CASE(AVLTree_Rotations) {
    AVLTree tree;
    
    // Сценарий, вызывающий LL-вращение
    tree.insert(30);
    tree.insert(20);
    tree.insert(10);
    
    BOOST_TEST(tree.find(10) == true);
    BOOST_TEST(tree.find(20) == true);
    BOOST_TEST(tree.find(30) == true);
    
    // Сценарий, вызывающий RR-вращение
    AVLTree tree2;
    tree2.insert(10);
    tree2.insert(20);
    tree2.insert(30);
    
    BOOST_TEST(tree2.find(10) == true);
    BOOST_TEST(tree2.find(20) == true);
    BOOST_TEST(tree2.find(30) == true);
}

BOOST_AUTO_TEST_CASE(AVLTree_Serialization) {
    AVLTree tree1;
    
    tree1.insert(50);
    tree1.insert(30);
    tree1.insert(70);
    tree1.insert(20);
    tree1.insert(40);
    tree1.insert(60);
    tree1.insert(80);
    
    std::string filename = generate_temp_filename();
    
    try {
        tree1.serialize(filename);
        
        AVLTree tree2;
        tree2.deserialize(filename);
        
        // Проверяем, что все элементы присутствуют
        BOOST_TEST(tree2.find(50) == true);
        BOOST_TEST(tree2.find(30) == true);
        BOOST_TEST(tree2.find(70) == true);
        BOOST_TEST(tree2.find(20) == true);
        BOOST_TEST(tree2.find(40) == true);
        BOOST_TEST(tree2.find(60) == true);
        BOOST_TEST(tree2.find(80) == true);
        BOOST_TEST(tree2.find(99) == false);
        
        // Проверяем балансировку
        tree2.remove(40);
        tree2.remove(60);
        tree2.remove(80);
        
        BOOST_TEST(tree2.find(50) == true);
        BOOST_TEST(tree2.find(30) == true);
        BOOST_TEST(tree2.find(20) == true);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для HashTableOpen
// ==========================================
BOOST_AUTO_TEST_SUITE(HashTableOpenTestSuite)

BOOST_AUTO_TEST_CASE(HashTableOpen_BasicOperations) {
    HashTableOpen table(4);
    
    BOOST_TEST(table.get(10).has_value() == false);
}

BOOST_AUTO_TEST_CASE(HashTableOpen_Insert_Get) {
    HashTableOpen table(4);
    
    table.insert(10, 100);
    table.insert(20, 200);
    table.insert(30, 300);
    
    auto val1 = table.get(10);
    auto val2 = table.get(20);
    auto val3 = table.get(30);
    
    BOOST_TEST(val1.has_value() == true);
    BOOST_TEST(val1.value() == 100);
    
    BOOST_TEST(val2.has_value() == true);
    BOOST_TEST(val2.value() == 200);
    
    BOOST_TEST(val3.has_value() == true);
    BOOST_TEST(val3.value() == 300);
    
    BOOST_TEST(table.get(99).has_value() == false);
}

BOOST_AUTO_TEST_CASE(HashTableOpen_Update) {
    HashTableOpen table(4);
    
    table.insert(10, 100);
    table.insert(10, 200); // Обновление значения
    
    auto val = table.get(10);
    BOOST_TEST(val.has_value() == true);
    BOOST_TEST(val.value() == 200);
}

BOOST_AUTO_TEST_CASE(HashTableOpen_Remove) {
    HashTableOpen table(4);
    
    table.insert(10, 100);
    table.insert(20, 200);
    table.insert(30, 300);
    
    BOOST_TEST(table.get(20).has_value() == true);
    
    table.remove(20);
    BOOST_TEST(table.get(20).has_value() == false);
    
    table.remove(99); // Удаление несуществующего элемента
    
    // Проверка, что остальные элементы на месте
    BOOST_TEST(table.get(10).value() == 100);
    BOOST_TEST(table.get(30).value() == 300);
}

BOOST_AUTO_TEST_CASE(HashTableOpen_Resize) {
    HashTableOpen table(2); // Маленькая начальная ёмкость
    
    for (int i = 0; i < 10; ++i) {
        table.insert(i, i * 10);
    }
    
    // Проверяем, что все элементы доступны после ресайза
    for (int i = 0; i < 10; ++i) {
        auto val = table.get(i);
        BOOST_TEST(val.has_value() == true);
        BOOST_TEST(val.value() == i * 10);
    }
}

BOOST_AUTO_TEST_CASE(HashTableOpen_Serialization) {
    HashTableOpen table1(4);
    
    table1.insert(1, 100);
    table1.insert(2, 200);
    table1.insert(3, 300);
    table1.insert(4, 400);
    
    std::string filename = generate_temp_filename();
    
    try {
        table1.serialize(filename);
        
        HashTableOpen table2(1); // Начальная ёмкость будет перезаписана
        table2.deserialize(filename);
        
        BOOST_TEST(table2.get(1).value() == 100);
        BOOST_TEST(table2.get(2).value() == 200);
        BOOST_TEST(table2.get(3).value() == 300);
        BOOST_TEST(table2.get(4).value() == 400);
        BOOST_TEST(table2.get(5).has_value() == false);
        
        // Проверка удаления после десериализации
        table2.remove(2);
        BOOST_TEST(table2.get(2).has_value() == false);
        BOOST_TEST(table2.get(1).value() == 100);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_CASE(HashTableOpen_Collisions) {
    HashTableOpen table(5);
    
    // Создаём коллизии
    table.insert(0, 100);  // hash(0) = 0
    table.insert(5, 200);  // hash(5) = 0 (коллизия)
    table.insert(10, 300); // hash(10) = 0 (коллизия)
    
    BOOST_TEST(table.get(0).value() == 100);
    BOOST_TEST(table.get(5).value() == 200);
    BOOST_TEST(table.get(10).value() == 300);
    
    table.remove(5);
    BOOST_TEST(table.get(0).value() == 100);
    BOOST_TEST(table.get(5).has_value() == false);
    BOOST_TEST(table.get(10).value() == 300);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты для HashTableChain
// ==========================================
BOOST_AUTO_TEST_SUITE(HashTableChainTestSuite)

BOOST_AUTO_TEST_CASE(HashTableChain_BasicOperations) {
    HashTableChain table(4);
    
    BOOST_TEST(table.get(10).has_value() == false);
}

BOOST_AUTO_TEST_CASE(HashTableChain_Insert_Get) {
    HashTableChain table(4);
    
    table.insert(10, 100);
    table.insert(20, 200);
    table.insert(30, 300);
    
    BOOST_TEST(table.get(10).value() == 100);
    BOOST_TEST(table.get(20).value() == 200);
    BOOST_TEST(table.get(30).value() == 300);
    BOOST_TEST(table.get(99).has_value() == false);
}

BOOST_AUTO_TEST_CASE(HashTableChain_Update) {
    HashTableChain table(4);
    
    table.insert(10, 100);
    table.insert(10, 200); // Цепочки обновляют значение
    
    BOOST_TEST(table.get(10).value() == 200);
}

BOOST_AUTO_TEST_CASE(HashTableChain_Remove) {
    HashTableChain table(4);
    
    table.insert(10, 100);
    table.insert(20, 200);
    table.insert(30, 300);
    
    table.remove(20);
    BOOST_TEST(table.get(20).has_value() == false);
    BOOST_TEST(table.get(10).value() == 100);
    BOOST_TEST(table.get(30).value() == 300);
    
    // Удаление из середины цепочки
    table.insert(2, 400);  // Предположим, hash(2) = 2
    table.insert(6, 500);  // hash(6) = 2 (коллизия)
    table.insert(10, 600); // hash(10) = 2 (коллизия)
    
    table.remove(6);
    BOOST_TEST(table.get(2).value() == 400);
    BOOST_TEST(table.get(6).has_value() == false);
    BOOST_TEST(table.get(10).value() == 600);
}

BOOST_AUTO_TEST_CASE(HashTableChain_Serialization) {
    HashTableChain table1(4);
    
    table1.insert(1, 100);
    table1.insert(2, 200);
    table1.insert(3, 300);
    table1.insert(5, 400); // Коллизия с 1
    
    std::string filename = generate_temp_filename();
    
    try {
        table1.serialize(filename);
        
        HashTableChain table2(1);
        table2.deserialize(filename);
        
        BOOST_TEST(table2.get(1).value() == 100);
        BOOST_TEST(table2.get(2).value() == 200);
        BOOST_TEST(table2.get(3).value() == 300);
        BOOST_TEST(table2.get(5).value() == 400);
        
        // Проверка цепочек
        table2.remove(1);
        BOOST_TEST(table2.get(1).has_value() == false);
        BOOST_TEST(table2.get(5).value() == 400);
    }
    catch (...) {
        fs::remove(filename);
        throw;
    }
    
    fs::remove(filename);
}

BOOST_AUTO_TEST_CASE(HashTableChain_LongChains) {
    HashTableChain table(2); // Маленькая ёмкость для создания длинных цепочек
    
    for (int i = 0; i < 20; ++i) {
        table.insert(i, i * 10);
    }
    
    for (int i = 0; i < 20; ++i) {
        BOOST_TEST(table.get(i).value() == i * 10);
    }
    
    // Удаление из разных позиций цепочек
    table.remove(0);  // Начало цепочки
    table.remove(10); // Середина цепочки
    table.remove(19); // Конец цепочки
    
    BOOST_TEST(table.get(0).has_value() == false);
    BOOST_TEST(table.get(10).has_value() == false);
    BOOST_TEST(table.get(19).has_value() == false);
    BOOST_TEST(table.get(5).value() == 50);
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Комплексные тесты взаимодействия
// ==========================================
BOOST_AUTO_TEST_SUITE(IntegrationTestSuite)

BOOST_AUTO_TEST_CASE(MultipleStructuresInteraction) {
    // Тест использования нескольких структур вместе
    MyArray arr;
    SinglyLinkedList list;
    MyStack stack;
    MyQueue queue;
    
    // Заполняем структуры
    for (int i = 0; i < 5; ++i) {
        arr.add_to_end(i);
        list.push_back(i);
        stack.push(i);
        queue.push(i);
    }
    
    // Проверяем согласованность
    BOOST_TEST(arr.get_length() == 5);
    BOOST_TEST(list.get_size() == 5);
    
    // Проверяем LIFO/FIFO
    BOOST_TEST(stack.peek() == 4); // LIFO - последний добавленный
    BOOST_TEST(queue.peek() == 0); // FIFO - первый добавленный
}

BOOST_AUTO_TEST_CASE(LargeDatasetTest) {
    // Тест с большим объёмом данных
    HashTableOpen hashTable(100);
    AVLTree tree;
    
    const int NUM_ELEMENTS = 1000;
    
    // Вставляем одинаковые данные в обе структуры
    for (int i = 0; i < NUM_ELEMENTS; ++i) {
        int key = i * 2;
        int value = i * 10;
        
        hashTable.insert(key, value);
        tree.insert(key);
    }
    
    // Проверяем наличие всех элементов
    for (int i = 0; i < NUM_ELEMENTS; ++i) {
        int key = i * 2;
        BOOST_TEST(hashTable.get(key).has_value() == true);
        BOOST_TEST(tree.find(key) == true);
    }
    
    // Удаляем половину элементов
    for (int i = 0; i < NUM_ELEMENTS; i += 2) {
        int key = i * 2;
        hashTable.remove(key);
        tree.remove(key);
    }
    
    // Проверяем, что правильные элементы удалены
    for (int i = 0; i < NUM_ELEMENTS; ++i) {
        int key = i * 2;
        bool shouldExist = (i % 2 == 1);
        
        BOOST_TEST(hashTable.get(key).has_value() == shouldExist);
        BOOST_TEST(tree.find(key) == shouldExist);
    }
}

BOOST_AUTO_TEST_CASE(SerializationChain) {
    // Тест цепочки сериализаций
    DoublyLinkedList list1;
    
    for (int i = 0; i < 10; ++i) {
        list1.push_back(i * 100);
    }
    
    std::string filename1 = generate_temp_filename();
    std::string filename2 = generate_temp_filename();
    
    try {
        // Сериализуем в первый файл
        list1.serialize(filename1);
        
        // Десериализуем и сразу сериализуем во второй
        DoublyLinkedList list2;
        list2.deserialize(filename1);
        list2.serialize(filename2);
        
        // Десериализуем из второго файла
        DoublyLinkedList list3;
        list3.deserialize(filename2);
        
        BOOST_TEST(list3.get_size() == list1.get_size());
        
        // Модифицируем и проверяем
        list3.remove_by_value(0);
        list3.push_front(999);
        BOOST_TEST(list3.get_size() == 10); // Один удалили, один добавили
        
        fs::remove(filename1);
        fs::remove(filename2);
    }
    catch (...) {
        fs::remove(filename1);
        fs::remove(filename2);
        throw;
    }
}

BOOST_AUTO_TEST_SUITE_END()

// ==========================================
// Тесты пограничных случаев
// ==========================================
BOOST_AUTO_TEST_SUITE(EdgeCaseTestSuite)

BOOST_AUTO_TEST_CASE(EmptyStructuresOperations) {
    // Тесты операций на пустых структурах
    MyArray emptyArr;
    BOOST_TEST(emptyArr.get_length() == 0);
    BOOST_CHECK_THROW(emptyArr.get_at_index(0), std::out_of_range);
    BOOST_CHECK_THROW(emptyArr.remove_at_index(0), std::out_of_range);
    
    SinglyLinkedList emptyList;
    emptyList.pop_front(); // Не должно падать
    emptyList.pop_back();  // Не должно падать
    BOOST_TEST(emptyList.get_size() == 0);
    
    MyStack emptyStack;
    BOOST_CHECK_THROW(emptyStack.peek(), std::runtime_error);
    emptyStack.pop(); // Не должно падать
    
    MyQueue emptyQueue;
    BOOST_CHECK_THROW(emptyQueue.peek(), std::runtime_error);
    emptyQueue.pop(); // Не должно падать
    
    AVLTree emptyTree;
    BOOST_TEST(emptyTree.find(10) == false);
    emptyTree.remove(10); // Не должно падать
    
    HashTableOpen emptyHash(10);
    BOOST_TEST(emptyHash.get(10).has_value() == false);
    emptyHash.remove(10); // Не должно падать
}

BOOST_AUTO_TEST_CASE(NegativeAndZeroValues) {
    // Тесты с отрицательными и нулевыми значениями
    MyArray arr;
    arr.add_to_end(-10);
    arr.add_to_end(0);
    arr.add_to_end(10);
    
    BOOST_TEST(arr.get_at_index(0) == -10);
    BOOST_TEST(arr.get_at_index(1) == 0);
    BOOST_TEST(arr.get_at_index(2) == 10);
    
    HashTableOpen hashTable(10);
    hashTable.insert(-5, -100);
    hashTable.insert(0, 0);
    hashTable.insert(5, 100);
    
    BOOST_TEST(hashTable.get(-5).value() == -100);
    BOOST_TEST(hashTable.get(0).value() == 0);
    BOOST_TEST(hashTable.get(5).value() == 100);
    
    AVLTree tree;
    tree.insert(-10);
    tree.insert(0);
    tree.insert(10);
    
    BOOST_TEST(tree.find(-10) == true);
    BOOST_TEST(tree.find(0) == true);
    BOOST_TEST(tree.find(10) == true);
}

BOOST_AUTO_TEST_CASE(DuplicateValues) {
    // Тесты с дублирующимися значениями
    HashTableChain hashTable(10);
    
    // Множественные вставки с одним ключом
    hashTable.insert(1, 100);
    hashTable.insert(1, 200);
    hashTable.insert(1, 300);
    
    // Должен остаться последний
    BOOST_TEST(hashTable.get(1).value() == 300);
    
    // AVL дерево игнорирует дубликаты
    AVLTree tree;
    tree.insert(10);
    tree.insert(10);
    tree.insert(10);
    
    tree.remove(10);
    BOOST_TEST(tree.find(10) == false);
}

BOOST_AUTO_TEST_CASE(SingleElementStructures) {
    // Тесты структур с одним элементом
    SinglyLinkedList singleList;
    singleList.push_front(42);
    
    BOOST_TEST(singleList.get_size() == 1);
    singleList.pop_front();
    BOOST_TEST(singleList.get_size() == 0);
    
    MyStack singleStack;
    singleStack.push(99);
    
    BOOST_TEST(singleStack.peek() == 99);
    singleStack.pop();
    BOOST_CHECK_THROW(singleStack.peek(), std::runtime_error);
    
    HashTableOpen singleHash(1);
    singleHash.insert(7, 777);
    
    BOOST_TEST(singleHash.get(7).value() == 777);
    singleHash.remove(7);
    BOOST_TEST(singleHash.get(7).has_value() == false);
}

BOOST_AUTO_TEST_CASE(stack_print) {
    MyStack s;
    for (int i = 1; i <= 3; ++i) s.push(i);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    s.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_CASE(queue_print) {
    MyQueue q;
    for (int i = 1; i <= 3; ++i) q.push(i);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    q.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_CASE(array_print) {
    MyArray arr;
    arr.add_at_index(0, 10);
    arr.add_at_index(1, 20);
    arr.add_at_index(2, 30);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    arr.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_CASE(doublylist_print) {
    DoublyLinkedList list;
    list.push_back(1);
    list.push_back(2);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    list.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_CASE(singlylist_print) {
    SinglyLinkedList list;
    list.push_back(1);
    list.push_back(2);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    list.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_CASE(hashtableopen_print) {
    HashTableOpen ht;
    ht.insert(1, 10);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    ht.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_CASE(hashtablechain_print) {
    HashTableChain ht;
    ht.insert(1, 10);

    std::ostringstream oss;
    std::streambuf* old = std::cout.rdbuf(oss.rdbuf());
    ht.print();
    std::cout.rdbuf(old);

    BOOST_TEST(!oss.str().empty());
}

BOOST_AUTO_TEST_SUITE_END()