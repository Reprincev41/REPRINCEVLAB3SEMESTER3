#ifndef ARRAY_H
#define ARRAY_H

#include <cstddef>
#include <fstream>
#include <stdexcept>

class MyArray {
private:
    int* data;
    size_t capacity;
    size_t size;
    void resize(size_t new_capacity);

public:
    MyArray();
    ~MyArray();
    
    void add_to_end(int value);
    void add_at_index(size_t index, int value);
    int get_at_index(size_t index) const;
    void remove_at_index(size_t index);
    void replace_at_index(size_t index, int value);
    size_t get_length() const;
    void print() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif
