"""
    Leetcode 1424
    https://leetcode-cn.com/problems/diagonal-traverse-ii/


Input: nums = [[1,2,3,4,5],
               [6,7],
               [8],
               [9,10,11],
               [12,13,14,15,16]]
Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]

"""


def findDiagonalOrder(nums):
    res = []
    for i in range(len(nums)):
        for j in range(len(nums[i])):
            res.append([i + j, j, nums[i][j]])
    res.sort()
    return [_[2] for _ in res]
