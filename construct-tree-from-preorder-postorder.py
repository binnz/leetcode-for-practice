"""
    Leetcode 889
    From https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/submissions/
"""

# Definition for a binary tree node.


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def constructFromPrePost(pre, post):
    if not pre:
        return None
    if len(pre) == 1:
        return TreeNode(pre[0])
    root = TreeNode(pre[0])
    index = post.index(pre[1])
    left_pre = [i for i in pre[1:] if i in post[:index + 1]]
    right_pre = [i for i in pre[1:] if i in post[index + 1:-1]]
    root.left = constructFromPrePost(left_pre, post[:index + 1])
    root.right = constructFromPrePost(right_pre, post[index + 1:-1])
    return root
