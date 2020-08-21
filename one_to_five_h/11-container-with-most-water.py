"""
    Leetcode 11
    https://leetcode-cn.com/problems/container-with-most-water/

"""


def maxArea(height):
    left, right = 0, len(height) - 1
    res = 0
    while left < right:
        res = max(res, min(height[left], height[right]) * (right - left))
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1
    return res