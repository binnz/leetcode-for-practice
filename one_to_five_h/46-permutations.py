"""
    Leetcode 46
    From https://leetcode-cn.com/problems/permutations/
"""


def permute(nums):
    if not nums:
        return []
    res = []
    length = len(nums)

    def dfs(depth, path, visited):
        if depth == length:
            res.append(path)
            return

        for i in range(0, length):
            if visited[i]:
                continue
            visited[i] = True
            dfs(depth + 1, path + [nums[i]], visited)
            visited[i] = False
    visited = [False for i in range(length)]
    dfs(0, [], visited)
    return res
