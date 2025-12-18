#include "HashTableOpen.h"
#include <iostream>
#include <fstream>
#include <cmath>

// ==========================================
// HashTableOpen Implementation
// ==========================================

HashTableOpen::HashTableOpen(size_t init_cap) : size(0), capacity(init_cap) {
    table = new HashEntry[capacity];
    for (size_t i = 0; i < capacity; i++) {
        table[i].isOccupied = false;
        table[i].isDeleted = false;
    }
}

HashTableOpen::~HashTableOpen() {
    delete[] table;
}

size_t HashTableOpen::hash(int key) const {
    return std::abs(key) % capacity;
}

void HashTableOpen::resize() {
    size_t old_capacity = capacity;
    HashEntry* old_table = table;
    
    capacity *= 2;
    table = new HashEntry[capacity];
    for (size_t i = 0; i < capacity; i++) {
        table[i].isOccupied = false;
        table[i].isDeleted = false;
    }
    
    size = 0;
    for (size_t i = 0; i < old_capacity; i++) {
        if (old_table[i].isOccupied && !old_table[i].isDeleted) {
            insert(old_table[i].key, old_table[i].value);
        }
    }
    
    delete[] old_table;
}

void HashTableOpen::insert(int key, int value) {
    if (size >= capacity * 0.7) {
        resize();
    }
    
    size_t idx = hash(key);
    size_t start_idx = idx;
    
    while (table[idx].isOccupied && !table[idx].isDeleted && table[idx].key != key) {
        idx = (idx + 1) % capacity;
        if (idx == start_idx) return;
    }
    
    table[idx].key = key;
    table[idx].value = value;
    table[idx].isOccupied = true;
    table[idx].isDeleted = false;
    size++;
}

std::optional<int> HashTableOpen::get(int key) {
    size_t idx = hash(key);
    size_t start_idx = idx;
    
    while (table[idx].isOccupied) {
        if (!table[idx].isDeleted && table[idx].key == key) {
            return table[idx].value;
        }
        idx = (idx + 1) % capacity;
        if (idx == start_idx) break;
    }
    
    return std::nullopt;
}

void HashTableOpen::remove(int key) {
    size_t idx = hash(key);
    size_t start_idx = idx;
    
    while (table[idx].isOccupied) {
        if (!table[idx].isDeleted && table[idx].key == key) {
            table[idx].isDeleted = true;
            size--;
            return;
        }
        idx = (idx + 1) % capacity;
        if (idx == start_idx) break;
    }
}

void HashTableOpen::print() const {
    std::cout << "HashTableOpen:\n";
    for (size_t i = 0; i < capacity; i++) {
        if (table[i].isOccupied && !table[i].isDeleted) {
            std::cout << "[" << i << "]: " << table[i].key << " -> " << table[i].value << "\n";
        }
    }
}

void HashTableOpen::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    file.write(reinterpret_cast<const char*>(&capacity), sizeof(capacity));
    
    for (size_t i = 0; i < capacity; i++) {
        file.write(reinterpret_cast<const char*>(&table[i].key), sizeof(int));
        file.write(reinterpret_cast<const char*>(&table[i].value), sizeof(int));
        file.write(reinterpret_cast<const char*>(&table[i].isOccupied), sizeof(bool));
        file.write(reinterpret_cast<const char*>(&table[i].isDeleted), sizeof(bool));
    }
    
    file.close();
}

void HashTableOpen::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    file.read(reinterpret_cast<char*>(&capacity), sizeof(capacity));
    
    delete[] table;
    table = new HashEntry[capacity];
    
    for (size_t i = 0; i < capacity; i++) {
        file.read(reinterpret_cast<char*>(&table[i].key), sizeof(int));
        file.read(reinterpret_cast<char*>(&table[i].value), sizeof(int));
        file.read(reinterpret_cast<char*>(&table[i].isOccupied), sizeof(bool));
        file.read(reinterpret_cast<char*>(&table[i].isDeleted), sizeof(bool));
    }
    
    file.close();
}
