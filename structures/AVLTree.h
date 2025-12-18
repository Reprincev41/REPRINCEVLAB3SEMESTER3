#ifndef AVLTREE_H
#define AVLTREE_H

#include <fstream>
#include <vector>

struct AVLNode {
    int key;
    AVLNode* left;
    AVLNode* right;
    int height;
    AVLNode(int k) : key(k), left(nullptr), right(nullptr), height(1) {}
};

class AVLTree {
private:
    AVLNode* root;
    
    static int height(AVLNode* N);
    static int max(int a, int b);
    AVLNode* rightRotate(AVLNode* y);
    AVLNode* leftRotate(AVLNode* x);
    int getBalance(AVLNode* N);
    AVLNode* insertNode(AVLNode* node, int key);
    static AVLNode* minValueNode(AVLNode* node);
    AVLNode* deleteNode(AVLNode* root, int key);
    void inOrder(AVLNode* root) const;
    void destroyTree(AVLNode* node);
    void serializeHelper(AVLNode* node, std::ofstream& file) const;
    AVLNode* deserializeHelper(std::ifstream& file);

public:
    AVLTree();
    ~AVLTree();
    
    void insert(int key);
    void remove(int key);
    bool find(int key) const;
    void print() const;
    
    // Serialization/Deserialization
    void serialize(const std::string& filename) const;
    void deserialize(const std::string& filename);
};

#endif
