"""
    Leetcode 39
    From https://leetcode-cn.com/problems/combination-sum/
"""


def combinationSum(self, candidates, target):
    if not candidates:
        return []
    res = []
    candidates.sort()

    def dfs(start, sums, path):
        if sums == target:
            res.append(path)
            return
        if sums > target:
            return
        for i in range(start, len(candidates)):
            if sums + candidates[i] > target:
                break
            dfs(i, sums + candidates[i], path + [candidates[i]])
    dfs(0, 0, [])
    return res
