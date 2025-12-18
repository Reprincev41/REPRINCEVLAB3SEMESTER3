#include "Stack.h"
#include <iostream>

MyStack::MyStack() : topNode(nullptr) {}

MyStack::~MyStack() { while (topNode != nullptr) { pop(); } }

void MyStack::push(int value) {
    auto* newNode = new StackNode(value);
    newNode->next = topNode;
    topNode = newNode;
}

void MyStack::pop() {
    if (topNode != nullptr) {
        StackNode* temp = topNode;
        topNode = topNode->next;
        delete temp;
    }
}

int MyStack::peek() const {
    if (topNode == nullptr) { throw std::runtime_error("Stack empty"); }
    return topNode->data;
}

void MyStack::print() const {
    std::cout << "Stack [";
    StackNode* curr = topNode;
    while (curr != nullptr) {
        std::cout << curr->data;
        if (curr->next != nullptr) std::cout << " -> ";
        curr = curr->next;
    }
    std::cout << "]\n";
}

void MyStack::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    size_t count = 0;
    StackNode* curr = topNode;
    while (curr != nullptr) { count++; curr = curr->next; }
    
    file.write(reinterpret_cast<const char*>(&count), sizeof(count));
    curr = topNode;
    while (curr != nullptr) {
        file.write(reinterpret_cast<const char*>(&curr->data), sizeof(int));
        curr = curr->next;
    }
    file.close();
}

void MyStack::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    // Clear existing stack
    while (topNode != nullptr) { pop(); }
    
    size_t count;
    file.read(reinterpret_cast<char*>(&count), sizeof(count));
    for (size_t i = 0; i < count; i++) {
        int value;
        file.read(reinterpret_cast<char*>(&value), sizeof(int));
        push(value);
    }
    file.close();
}