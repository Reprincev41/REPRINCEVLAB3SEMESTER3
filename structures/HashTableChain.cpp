#include "HashTableChain.h"
#include <iostream>
#include <fstream>
#include <cmath>


// ==========================================
// HashTableChain Implementation
// ==========================================

HashTableChain::HashTableChain(size_t init_cap) : size(0), capacity(init_cap) {
    table = new ChainNode*[capacity];
    for (size_t i = 0; i < capacity; i++) {
        table[i] = nullptr;
    }
}

HashTableChain::~HashTableChain() {
    for (size_t i = 0; i < capacity; i++) {
        ChainNode* node = table[i];
        while (node != nullptr) {
            ChainNode* temp = node;
            node = node->next;
            delete temp;
        }
    }
    delete[] table;
}

size_t HashTableChain::hash(int key) const {
    return std::abs(key) % capacity;
}

void HashTableChain::insert(int key, int value) {
    size_t idx = hash(key);
    ChainNode* newNode = new ChainNode(key, value);
    newNode->next = table[idx];
    table[idx] = newNode;
    size++;
}

std::optional<int> HashTableChain::get(int key) {
    size_t idx = hash(key);
    ChainNode* curr = table[idx];
    
    while (curr != nullptr) {
        if (curr->key == key) {
            return curr->value;
        }
        curr = curr->next;
    }
    
    return std::nullopt;
}

void HashTableChain::remove(int key) {
    size_t idx = hash(key);
    ChainNode* curr = table[idx];
    ChainNode* prev = nullptr;
    
    while (curr != nullptr) {
        if (curr->key == key) {
            if (prev != nullptr) {
                prev->next = curr->next;
            } else {
                table[idx] = curr->next;
            }
            delete curr;
            size--;
            return;
        }
        prev = curr;
        curr = curr->next;
    }
}

void HashTableChain::print() const {
    std::cout << "HashTableChain:\n";
    for (size_t i = 0; i < capacity; i++) {
        if (table[i] != nullptr) {
            std::cout << "[" << i << "]: ";
            ChainNode* curr = table[i];
            while (curr != nullptr) {
                std::cout << "(" << curr->key << "->" << curr->value << ")";
                if (curr->next != nullptr) std::cout << " -> ";
                curr = curr->next;
            }
            std::cout << "\n";
        }
    }
}

void HashTableChain::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    file.write(reinterpret_cast<const char*>(&size), sizeof(size));
    file.write(reinterpret_cast<const char*>(&capacity), sizeof(capacity));
    
    for (size_t i = 0; i < capacity; i++) {
        size_t chain_size = 0;
        ChainNode* temp = table[i];
        while (temp != nullptr) {
            chain_size++;
            temp = temp->next;
        }
        
        file.write(reinterpret_cast<const char*>(&chain_size), sizeof(chain_size));
        
        ChainNode* curr = table[i];
        while (curr != nullptr) {
            file.write(reinterpret_cast<const char*>(&curr->key), sizeof(int));
            file.write(reinterpret_cast<const char*>(&curr->value), sizeof(int));
            curr = curr->next;
        }
    }
    
    file.close();
}

void HashTableChain::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    // Clear existing data
    for (size_t i = 0; i < capacity; i++) {
        ChainNode* node = table[i];
        while (node != nullptr) {
            ChainNode* temp = node;
            node = node->next;
            delete temp;
        }
    }
    delete[] table;
    
    file.read(reinterpret_cast<char*>(&size), sizeof(size));
    file.read(reinterpret_cast<char*>(&capacity), sizeof(capacity));
    
    table = new ChainNode*[capacity];
    for (size_t i = 0; i < capacity; i++) {
        table[i] = nullptr;
    }
    
    for (size_t i = 0; i < capacity; i++) {
        size_t chain_size;
        file.read(reinterpret_cast<char*>(&chain_size), sizeof(chain_size));
        
        for (size_t j = 0; j < chain_size; j++) {
            int key, value;
            file.read(reinterpret_cast<char*>(&key), sizeof(int));
            file.read(reinterpret_cast<char*>(&value), sizeof(int));
            insert(key, value);
        }
    }
    
    file.close();
}
