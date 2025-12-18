#include "SinglyList.h"
#include <iostream>


SinglyLinkedList::SinglyLinkedList() : head(nullptr), tail(nullptr), size(0) {}

SinglyLinkedList::~SinglyLinkedList() {
    while (head != nullptr) { SNode* temp = head; head = head->next; delete temp; }
}

void SinglyLinkedList::push_front(int value) {
    auto* newNode = new SNode(value);
    newNode->next = head;
    head = newNode;
    if (tail == nullptr) { tail = head; }
    size++;
}

void SinglyLinkedList::push_back(int value) {
    auto* newNode = new SNode(value);
    if (head == nullptr) { head = tail = newNode; }
    else { tail->next = newNode; tail = newNode; }
    size++;
}

void SinglyLinkedList::insert_after(size_t index, int value) {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    SNode* curr = head;
    for (size_t i = 0; i < index; ++i) { curr = curr->next; }
    auto* newNode = new SNode(value);
    newNode->next = curr->next;
    curr->next = newNode;
    if (curr == tail) { tail = newNode; }
    size++;
}

void SinglyLinkedList::insert_before(size_t index, int value) {
    if (index == 0) { push_front(value); return; }
    if (index > size) { throw std::out_of_range("Index out of bounds"); }
    insert_after(index - 1, value);
}

void SinglyLinkedList::pop_front() {
    if (head == nullptr) { return; }
    SNode* temp = head;
    head = head->next;
    delete temp;
    if (head == nullptr) { tail = nullptr; }
    size--;
}

void SinglyLinkedList::pop_back() {
    if (head == nullptr) { return; }
    if (head == tail) { delete head; head = tail = nullptr; size = 0; return; }
    SNode* curr = head;
    while (curr->next != tail) { curr = curr->next; }
    delete tail;
    tail = curr;
    tail->next = nullptr;
    size--;
}

void SinglyLinkedList::remove_at(size_t index) {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    if (index == 0) { pop_front(); return; }
    SNode* curr = head;
    for (size_t i = 0; i < index - 1; ++i) { curr = curr->next; }
    SNode* toDel = curr->next;
    curr->next = toDel->next;
    if (toDel == tail) { tail = curr; }
    delete toDel;
    size--;
}

void SinglyLinkedList::remove_by_value(int value) {
    if (head == nullptr) { return; }
    if (head->data == value) { pop_front(); return; }
    SNode* curr = head;
    while ((curr->next != nullptr) && curr->next->data != value) { curr = curr->next; }
    if (curr->next != nullptr) {
        SNode* toDel = curr->next;
        curr->next = toDel->next;
        if (toDel == tail) { tail = curr; }
        delete toDel;
        size--;
    }
}

bool SinglyLinkedList::find(int value) const {
    SNode* curr = head;
    while (curr != nullptr) { if (curr->data == value) { return true; } curr = curr->next; }
    return false;
}

size_t SinglyLinkedList::get_size() const { return size; }

void SinglyLinkedList::print() const {
    std::cout << "SinglyLinkedList [";
    SNode* curr = head;
    while (curr != nullptr) {
        std::cout << curr->data;
        if (curr->next != nullptr) std::cout << " -> ";
        curr = curr->next;
    }
    std::cout << "]\n";
}

void SinglyLinkedList::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    SNode* curr = head;
    while (curr != nullptr) {
        file.write(reinterpret_cast<const char*>(&curr->data), sizeof(int));
        curr = curr->next;
    }
    file.close();
}

void SinglyLinkedList::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    // Clear existing list FIRST
    while (head != nullptr) {
        SNode* temp = head;
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