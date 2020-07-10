"""
    Leetcode 222
    https://leetcode-cn.com/problems/count-complete-tree-nodes/

"""


def countNodes(root):
    def dfs(node):
        if not node:
            return 0
        if not node.left and not node.right:
            return 1
        return dfs(node.left) + dfs(node.right) + 1
    return dfs(root)


def countNodes2(root):
    res = 0
    if not root:
        return res
    stack = [root]
    while stack:
        cur_node = stack.pop()
        res += 1
        if cur_node.left:
            stack.append(cur_node.left)
        if cur_node.right:
            stack.append(cur_node.right)
    return res
