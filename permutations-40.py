"""
    Leetcode 40
    From https://leetcode-cn.com/problems/permutations/
"""

def permute(nums):
    if not nums:
        return []
    res = []
    length = len(nums)
    def dfs(start, depth, path, visited):
        if depth == length:
            res.append(path)
            return
        
        for i in range(0, length):
            if visited[i]:
                continue
            visited[i] = True
            dfs(i + 1, depth + 1, path + [nums[i]], visited)
            visited[i] = False
    visited = [False for i in range(length)]
    dfs(0, 0, [], visited)
    return res
