"""
    Leetcode 18
    https://leetcode-cn.com/problems/4sum/

"""


def fourSum(nums, target):
    nums.sort()
    length = len(nums)
    res = []
    for i in range(length - 2):
        if i > 0 and nums[i] == nums[i - 1]:
            continue
        if nums[i] * 4 > target:
            break
        if nums[i] + nums[-1] * 3 < target:
            continue
        for j in range(i + 1, length - 2):
            if nums[i] + nums[j] * 3 > target:
                break
            if nums[i] + nums[j] + nums[-1] * 2 < target:
                continue
            k = j + 1
            m = length - 1
            if j > i + 1 and nums[j] == nums[j - 1]:
                continue
            if nums[i] + nums[j] + nums[k] * 2 > target:
                break
            while k < m:
                sums = nums[i] + nums[j] + nums[k] + nums[m]
                if sums == target:
                    res.append([nums[i], nums[j], nums[k], nums[m]])
                    k += 1
                    m -= 1
                    while nums[k] == nums[k - 1] and k < m:
                        k += 1
                    while nums[m] == nums[m + 1] and k < m:
                        m -= 1
                elif sums > target:
                    m -= 1
                else:
                    k += 1
    return res
