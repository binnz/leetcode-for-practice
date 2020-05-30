"""
    Leetcode 47
    From https://leetcode-cn.com/problems/permutations-ii/
"""


def permuteUnique(nums):
    if not nums:
        return []

    res = []
    length = len(nums)

    def dfs(depth, path, visited):
        if depth == length:
            res.append(path)
            return
        depth_visited = set()
        for i in range(0, length):
            if nums[i] not in depth_visited and visited >> i & 1 == 0:
                visited ^= 1 << i
                depth_visited.add(nums[i])
                dfs(depth + 1, path + [nums[i]], visited)
                visited ^= 1 << i

    dfs(0, [], 0)
    return res
