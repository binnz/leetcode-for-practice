"""
    Leetcode 16
    https://leetcode-cn.com/problems/3sum-closest/

"""


def threeSumClosest(nums, target):
    nums.sort()
    res = 0
    diff = float('inf')
    for i in range(len(nums) - 2):
        left = i + 1
        right = len(nums) - 1
        while left < right:
            sums = nums[i] + nums[left] + nums[right]
            if abs(target - sums) < diff:
                res = sums
                diff = abs(target - sums)
            if sums == target:
                return target
            elif sums < target:
                left += 1
            else:
                right -= 1
    return res
