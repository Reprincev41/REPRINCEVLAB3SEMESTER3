#include "Queue.h"
#include <iostream>

MyQueue::MyQueue() : frontNode(nullptr), rearNode(nullptr) {}

MyQueue::~MyQueue() { while (frontNode != nullptr) { pop(); } }

void MyQueue::push(int value) {
    auto* newNode = new QueueNode(value);
    if (rearNode == nullptr) { frontNode = rearNode = newNode; return; }
    rearNode->next = newNode;
    rearNode = newNode;
}

void MyQueue::pop() {
    if (frontNode == nullptr) { return; }
    QueueNode* temp = frontNode;
    frontNode = frontNode->next;
    if (frontNode == nullptr) { rearNode = nullptr; }
    delete temp;
}

int MyQueue::peek() const {
    if (frontNode == nullptr) { throw std::runtime_error("Queue empty"); }
    return frontNode->data;
}

void MyQueue::print() const {
    std::cout << "Queue [";
    QueueNode* curr = frontNode;
    while (curr != nullptr) {
        std::cout << curr->data;
        if (curr->next != nullptr) std::cout << " -> ";
        curr = curr->next;
    }
    std::cout << "]\n";
}

void MyQueue::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    size_t count = 0;
    QueueNode* curr = frontNode;
    while (curr != nullptr) { count++; curr = curr->next; }
    
    file.write(reinterpret_cast<const char*>(&count), sizeof(count));
    curr = frontNode;
    while (curr != nullptr) {
        file.write(reinterpret_cast<const char*>(&curr->data), sizeof(int));
        curr = curr->next;
    }
    file.close();
}

void MyQueue::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    // Clear existing queue
    while (frontNode != nullptr) { pop(); }
    
    size_t count;
    file.read(reinterpret_cast<char*>(&count), sizeof(count));
    for (size_t i = 0; i < count; i++) {
        int value;
        file.read(reinterpret_cast<char*>(&value), sizeof(int));
        push(value);
    }
    file.close();
}
