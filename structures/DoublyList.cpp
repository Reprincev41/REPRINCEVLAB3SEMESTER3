#include "DoublyList.h"
#include <iostream>

DoublyLinkedList::DoublyLinkedList() : head(nullptr), tail(nullptr), size(0) {}

DoublyLinkedList::~DoublyLinkedList() {
    while (head != nullptr) { DNode* temp = head; head = head->next; delete temp; }
}

size_t DoublyLinkedList::get_size() const { return size; }

void DoublyLinkedList::push_front(int value) {
    auto* newNode = new DNode(value);
    if (head == nullptr) { head = tail = newNode; }
    else { newNode->next = head; head->prev = newNode; head = newNode; }
    size++;
}

void DoublyLinkedList::push_back(int value) {
    auto* newNode = new DNode(value);
    if (tail == nullptr) { head = tail = newNode; }
    else { tail->next = newNode; newNode->prev = tail; tail = newNode; }
    size++;
}

void DoublyLinkedList::insert_after(size_t index, int value) {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    if (index == size - 1) { push_back(value); return; }
    DNode* curr = head;
    for (size_t i = 0; i < index; ++i) { curr = curr->next; }
    auto* newNode = new DNode(value);
    newNode->next = curr->next;
    newNode->prev = curr;
    curr->next->prev = newNode;
    curr->next = newNode;
    size++;
}

void DoublyLinkedList::insert_before(size_t index, int value) {
    if (index > size) { throw std::out_of_range("Index out of bounds"); }
    if (index == 0) { push_front(value); return; }
    insert_after(index - 1, value);
}

void DoublyLinkedList::pop_front() {
    if (head == nullptr) { return; }
    DNode* temp = head;
    head = head->next;
    if (head != nullptr) { head->prev = nullptr; } else { tail = nullptr; }
    delete temp;
    size--;
}

void DoublyLinkedList::pop_back() {
    if (tail == nullptr) { return; }
    DNode* temp = tail;
    tail = tail->prev;
    if (tail != nullptr) { tail->next = nullptr; } else { head = nullptr; }
    delete temp;
    size--;
}

void DoublyLinkedList::remove_at(size_t index) {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    if (index == 0) { pop_front(); return; }
    if (index == size - 1) { pop_back(); return; }
    DNode* curr = head;
    for (size_t i = 0; i < index; ++i) { curr = curr->next; }
    curr->prev->next = curr->next;
    curr->next->prev = curr->prev;
    delete curr;
    size--;
}

void DoublyLinkedList::remove_by_value(int value) {
    DNode* curr = head;
    while (curr != nullptr) {
        if (curr->data == value) {
            if (curr == head) { pop_front(); } 
            else if (curr == tail) { pop_back(); } 
            else {
                curr->prev->next = curr->next;
                curr->next->prev = curr->prev;
                delete curr;
                size--;
            }
            return;
        }
        curr = curr->next;
    }
}

bool DoublyLinkedList::find(int value) const {
    DNode* curr = head;
    while (curr != nullptr) { if (curr->data == value) { return true; } curr = curr->next; }
    return false;
}

void DoublyLinkedList::print() const {
    std::cout << "DoublyLinkedList [";
    DNode* curr = head;
    while (curr != nullptr) {
        std::cout << curr->data;
        if (curr->next != nullptr) std::cout << " <-> ";
        curr = curr->next;
    }
    std::cout << "]\n";
}

void DoublyLinkedList::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    DNode* curr = head;
    while (curr != nullptr) {
        file.write(reinterpret_cast<const char*>(&curr->data), sizeof(int));
        curr = curr->next;
    }
    file.close();
}

void DoublyLinkedList::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    // Clear existing list FIRST
    while (head != nullptr) {
        DNode* temp = head;
        head = head->next;
        delete temp;
    }
    head = tail = nullptr;
    size = 0;
    
    // THEN read size into a temporary variable
    size_t file_size;
    file.read(reinterpret_cast<char*>(&file_size), sizeof(file_size));
    
    // Add sanity check
    if (file_size > 1000000) {
        throw std::runtime_error("Suspiciously large size in file");
    }
    
    // Use the temporary variable
    for (size_t i = 0; i < file_size; i++) {
        int value;
        file.read(reinterpret_cast<char*>(&value), sizeof(int));
        push_back(value);
    }
    
    file.close();
}

