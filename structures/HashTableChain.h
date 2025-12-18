#ifndef HASHTABLECHAIN_H
#define HASHTABLECHAIN_H

#include <string>
#include <optional>

// ==========================================
// Hash Table (Chaining Method)
// ==========================================

class HashTableChain {
private:
    struct ChainNode {
        int key;
        int value;
        ChainNode* next;
        ChainNode(int k, int v) : key(k), value(v), next(nullptr) {}
    };
    
    ChainNode** table;
    size_t size;
    size_t capacity;
    
    size_t hash(int key) const;

public:
    HashTableChain(size_t init_cap = 8);
    ~HashTableChain();
    
    void insert(int key, int value);
    std::optional<int> get(int key);
    void remove(int key);
    void print() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif
