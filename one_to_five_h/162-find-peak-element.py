"""
    Leetcode 162
    https://leetcode-cn.com/problems/find-peak-element/

"""


def findPeakElement(nums):
    left = 0
    right = len(nums) - 1
    while left < right:
        mid = (left + right) >> 1
        if nums[mid] < nums[mid + 1]:
            left = mid + 1
        elif nums[mid] > nums[mid + 1]:
            right = mid
    assert left == right
    return left
