#include "AVLTree.h"
#include <iostream>

AVLTree::AVLTree() : root(nullptr) {}

AVLTree::~AVLTree() { destroyTree(root); }

void AVLTree::destroyTree(AVLNode* node) {
    if (node != nullptr) { 
        destroyTree(node->left); 
        destroyTree(node->right); 
        delete node; 
    }
}

int AVLTree::height(AVLNode* N) { return (N == nullptr) ? 0 : N->height; }

int AVLTree::max(int a, int b) { return (a > b) ? a : b; }

AVLNode* AVLTree::rightRotate(AVLNode* y) {
    AVLNode* x = y->left;
    AVLNode* T2 = x->right;
    x->right = y;
    y->left = T2;
    y->height = max(height(y->left), height(y->right)) + 1;
    x->height = max(height(x->left), height(x->right)) + 1;
    return x;
}

AVLNode* AVLTree::leftRotate(AVLNode* x) {
    AVLNode* y = x->right;
    AVLNode* T2 = y->left;
    y->left = x;
    x->right = T2;
    x->height = max(height(x->left), height(x->right)) + 1;
    y->height = max(height(y->left), height(y->right)) + 1;
    return y;
}

int AVLTree::getBalance(AVLNode* N) { return (N == nullptr) ? 0 : height(N->left) - height(N->right); }

AVLNode* AVLTree::insertNode(AVLNode* node, int key) {
    if (node == nullptr) { return new AVLNode(key); }
    
    if (key < node->key) { node->left = insertNode(node->left, key); }
    else if (key > node->key) { node->right = insertNode(node->right, key); }
    else { return node; }
    
    node->height = 1 + max(height(node->left), height(node->right));
    int balance = getBalance(node);
    
    // LL Case
    if (balance > 1 && key < node->left->key) { return rightRotate(node); }
    
    // RR Case
    if (balance < -1 && key > node->right->key) { return leftRotate(node); }
    
    // LR Case
    if (balance > 1 && key > node->left->key) {
        node->left = leftRotate(node->left);
        return rightRotate(node);
    }
    
    // RL Case
    if (balance < -1 && key < node->right->key) {
        node->right = rightRotate(node->right);
        return leftRotate(node);
    }
    
    return node;
}

AVLNode* AVLTree::minValueNode(AVLNode* node) {
    AVLNode* current = node;
    while (current->left != nullptr) { current = current->left; }
    return current;
}

AVLNode* AVLTree::deleteNode(AVLNode* root, int key) {
    if (root == nullptr) { return root; }
    
    if (key < root->key) { root->left = deleteNode(root->left, key); }
    else if (key > root->key) { root->right = deleteNode(root->right, key); }
    else {
        if ((root->left == nullptr) || (root->right == nullptr)) {
            AVLNode* temp = (root->left != nullptr) ? root->left : root->right;
            if (temp == nullptr) { temp = root; root = nullptr; }
            else { *root = *temp; }
            delete temp;
        } else {
            AVLNode* temp = minValueNode(root->right);
            root->key = temp->key;
            root->right = deleteNode(root->right, temp->key);
        }
    }
    
    if (root == nullptr) { return root; }
    
    root->height = 1 + max(height(root->left), height(root->right));
    int balance = getBalance(root);
    
    if (balance > 1 && getBalance(root->left) >= 0) { return rightRotate(root); }
    
    if (balance > 1 && getBalance(root->left) < 0) {
        root->left = leftRotate(root->left);
        return rightRotate(root);
    }
    
    if (balance < -1 && getBalance(root->right) <= 0) { return leftRotate(root); }
    
    if (balance < -1 && getBalance(root->right) > 0) {
        root->right = rightRotate(root->right);
        return leftRotate(root);
    }
    
    return root;
}

void AVLTree::inOrder(AVLNode* root) const {
    if (root != nullptr) {
        inOrder(root->left);
        std::cout << root->key << " ";
        inOrder(root->right);
    }
}

void AVLTree::insert(int key) { root = insertNode(root, key); }

void AVLTree::remove(int key) { root = deleteNode(root, key); }

bool AVLTree::find(int key) const {
    AVLNode* curr = root;
    while(curr != nullptr) {
        if (key == curr->key) { return true; }
        if (key < curr->key) curr = curr->left;
        else curr = curr->right;
    }
    return false;
}

void AVLTree::print() const {
    std::cout << "AVLTree (In-order): ";
    inOrder(root);
    std::cout << "\n";
}

void AVLTree::serializeHelper(AVLNode* node, std::ofstream& file) const {
    if (node == nullptr) {
        int null_marker = -1;
        file.write(reinterpret_cast<const char*>(&null_marker), sizeof(int));
        return;
    }
    
    file.write(reinterpret_cast<const char*>(&node->key), sizeof(int));
    serializeHelper(node->left, file);
    serializeHelper(node->right, file);
}

AVLNode* AVLTree::deserializeHelper(std::ifstream& file) {
    int key;
    file.read(reinterpret_cast<char*>(&key), sizeof(int));
    
    if (key == -1) { return nullptr; }
    
    AVLNode* node = new AVLNode(key);
    node->left = deserializeHelper(file);
    node->right = deserializeHelper(file);
    node->height = 1 + max(height(node->left), height(node->right));
    
    return node;
}

void AVLTree::serialize(const std::string& filename) const {
    std::ofstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for writing");
    
    serializeHelper(root, file);
    file.close();
}

void AVLTree::deserialize(const std::string& filename) {
    std::ifstream file(filename, std::ios::binary);
    if (!file) throw std::runtime_error("Cannot open file for reading");
    
    destroyTree(root);
    root = deserializeHelper(file);
    file.close();
}
