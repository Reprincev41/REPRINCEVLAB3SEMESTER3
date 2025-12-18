#ifndef QUEUE_H
#define QUEUE_H

#include <fstream>
#include <stdexcept>

class MyQueue {
private:
    struct QueueNode {
        int data;
        QueueNode* next;
        QueueNode(int val) : data(val), next(nullptr) {}
    } *frontNode, *rearNode;

public:
    MyQueue();
    ~MyQueue();
    
    void push(int value);
    void pop();
    int peek() const;
    void print() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif