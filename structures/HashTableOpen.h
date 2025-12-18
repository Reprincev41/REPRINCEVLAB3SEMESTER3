#ifndef HASHTABLEOPEN_H
#define HASHTABLEOPEN_H

#include <string>
#include <optional>

// ==========================================
// Hash Table (Open Addressing - Linear Probing)
// ==========================================

class HashTableOpen {
private:
    struct HashEntry {
        int key;
        int value;
        bool isOccupied;
        bool isDeleted;
    };
    
    HashEntry* table;
    size_t size;
    size_t capacity;
    
    size_t hash(int key) const;
    void resize();

public:
    HashTableOpen(size_t init_cap = 8);
    ~HashTableOpen();
    
    void insert(int key, int value);
    std::optional<int> get(int key);
    void remove(int key);
    void print() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif
