#include "Array.h"
#include <iostream>

MyArray::MyArray() : capacity(2), size(0) {
    data = new int[capacity];
}

MyArray::~MyArray() { delete[] data; }

void MyArray::resize(size_t new_capacity) {
    int* new_data = new int[new_capacity];
    for (size_t i = 0; i < size; i++) { new_data[i] = data[i]; }
    delete[] data;
    data = new_data;
    capacity = new_capacity;
}

void MyArray::add_to_end(int value) {
    if (size == capacity) { resize(capacity * 2); }
    data[size++] = value;
}

void MyArray::add_at_index(size_t index, int value) {
    if (index > size) { throw std::out_of_range("Index out of bounds"); }
    if (size == capacity) { resize(capacity * 2); }
    for (size_t i = size; i > index; i--) { data[i] = data[i - 1]; }
    data[index] = value;
    size++;
}

int MyArray::get_at_index(size_t index) const {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    return data[index];
}

void MyArray::remove_at_index(size_t index) {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    for (size_t i = index; i < size - 1; i++) { data[i] = data[i + 1]; }
    size--;
}

void MyArray::replace_at_index(size_t index, int value) {
    if (index >= size) { throw std::out_of_range("Index out of bounds"); }
    data[index] = value;
}

size_t MyArray::get_length() const { return size; }

void MyArray::print() const {
    std::cout << "Array [";
    for (size_t i = 0; i < size; i++) {
        std::cout << data[i];
        if (i < size - 1) std::cout << ", ";
    }
    std::cout << "]\n";
}

void MyArray::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    file.write(reinterpret_cast<const char*>(data), size * sizeof(int));
    file.close();
}

void MyArray::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    size_t new_size;
    file.read(reinterpret_cast<char*>(&new_size), sizeof(new_size));
    
    if (new_size > capacity) { resize(new_size); }
    size = new_size;
    
    file.read(reinterpret_cast<char*>(data), size * sizeof(int));
    file.close();
}
