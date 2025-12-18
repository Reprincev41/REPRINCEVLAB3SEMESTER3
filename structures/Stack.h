#ifndef STACK_H
#define STACK_H

#include <fstream>
#include <stdexcept>

class MyStack {
private:
    struct StackNode {
        int data;
        StackNode* next;
        StackNode(int val) : data(val), next(nullptr) {}
    } *topNode;

public:
    MyStack();
    ~MyStack();
    
    void push(int value);
    void pop();
    int peek() const;
    void print() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif