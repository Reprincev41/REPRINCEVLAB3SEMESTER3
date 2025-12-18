#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#include <cstddef>
#include <fstream>
#include <stdexcept>

// ==========================================
// Singly Linked List
// ==========================================

struct SNode {
    int data;
    SNode* next;
    SNode(int val) : data(val), next(nullptr) {}
};

class SinglyLinkedList {
private:
    SNode* head;
    SNode* tail;
    size_t size;

public:
    SinglyLinkedList();
    ~SinglyLinkedList();
    
    void push_front(int value);
    void push_back(int value);
    void insert_after(size_t index, int value);
    void insert_before(size_t index, int value);
    void pop_front();
    void pop_back();
    void remove_at(size_t index);
    void remove_by_value(int value);
    bool find(int value) const;
    void print() const;
    size_t get_size() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif
