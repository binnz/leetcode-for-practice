# leetcode-for-practice

## [Binary Tree Traversal This Way](/binary-search/binary-tree-traversal.py)

```
    Six ways of traverse the binary tree using one method

                   A
                 /   \
                /     \      
               B       C
    A-B-C  preorder
    B-A-C  inorder
    B-C-A  postorder
    A-C-B  reverse of postorder
    C-A-B  reverse of inorder
    C-B-A  reverse of preorder

    The key difference is the order of the three components below, just keep reverse order of the traverse order.

    if cur_node.right:
        stack.append((cur_node.right, False))
    if cur_node.left:
        stack.append((cur_node.left, False))
    stack.append((cur_node, True))
```

```
Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
```

### A-B-C
```
def preorderTraversal(root):
    if not root:
        return []
    res = []
    stack = [(root, False)]
    while stack:
        cur_node, visited = stack.pop()
        if visited:
            res.append(cur_node.val)
        else:
            if cur_node.right:
                stack.append((cur_node.right, False))
            if cur_node.left:
                stack.append((cur_node.left, False))
            stack.append((cur_node, True))
    return res
```

### B-A-C
```
def inorderTraversal(root):
    if not root:
        return []
    res = []
    stack = [(root, False)]
    while stack:
        cur_node, visited = stack.pop()
        if visited:
            res.append(cur_node.val)
        else:
            if cur_node.right:
                stack.append((cur_node.right, False))
            stack.append((cur_node, True))
            if cur_node.left:
                stack.append((cur_node.left, False))
    return res
```

### B-C-A
```
def postorderTraversal(root):
    if not root:
        return []
    res = []
    stack = [(root, False)]
    while stack:
        cur_node, visited = stack.pop()
        if visited:
            res.append(cur_node.val)
        else:
            stack.append((cur_node, True))
            if cur_node.right:
                stack.append((cur_node.right, False))
            if cur_node.left:
                stack.append((cur_node.left, False))
    return res
```

### C-B-A
```
def reverseOfpreorderTraversal(root):
    if not root:
        return []
    res = []
    stack = [(root, False)]
    while stack:
        cur_node, visited = stack.pop()
        if visited:
            res.append(cur_node.val)
        else:
            stack.append((cur_node, True))
            if cur_node.left:
                stack.append((cur_node.left, False))
            if cur_node.right:
                stack.append((cur_node.right, False))
    return res
```

## [Double Pointer](/double-pointer)

## [Linked List](/linked-list)

## [Tree](/tree)
