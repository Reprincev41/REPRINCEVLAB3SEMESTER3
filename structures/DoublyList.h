#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <cstddef>
#include <fstream>
#include <stdexcept>

struct DNode {
    int data;
    DNode* next;
    DNode* prev;
    DNode(int val) : data(val), next(nullptr), prev(nullptr) {}
};

class DoublyLinkedList {
private:
    DNode* head;
    DNode* tail;
    size_t size;

public:
    DoublyLinkedList();
    ~DoublyLinkedList();
    
    size_t get_size() const;
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
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif