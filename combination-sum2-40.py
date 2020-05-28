"""
    Leetcode 40
    From https://leetcode-cn.com/problems/combination-sum-ii/
"""


def combinationSum2(candidates, target):
    if not candidates:
        return []
    candidates.sort()
    length = len(candidates)
    res = []

    def dfs(start, sums, path):
        if sums == 0:
            res.append(path)
            return

        for i in range(start, length):
            if candidates[i] > sums:
                break
            if i > start and candidates[i] == candidates[i - 1]:
                continue
            dfs(i + 1, sums - candidates[i], path + [candidates[i]])
    dfs(0, target, [])
    return res
