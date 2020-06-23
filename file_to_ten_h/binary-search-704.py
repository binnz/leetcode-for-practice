"""
    Leetcode 704
    https://leetcode-cn.com/problems/binary-search/solution/
"""

"""
    Binary search return most left element index
    Examples:
        Input: [2, 3, 4, 4, 4, 7], 4
        Ouput: 2
"""


def binarySearchMostLeft(nums, target):
    if not nums:
        return -1
    left = 0
    right = len(nums)
    while left < right:
        mid = (right + left) >> 1
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid
    if left == len(nums):
        return -1
    return left if nums[left] == target else -1


"""
    Binary search return most right element index
    Examples:
        Input: [2, 3, 4, 4, 4, 7], 4
        Ouput: 4
"""


def binarySearchMostRight(nums, target):
    if not nums:
        return -1
    left = 0
    right = len(nums)

    while left < right:
        mid = (right + left) >> 1
        if nums[mid] > target:
            right = mid
        else:
            left = mid + 1

    if left == 0:
        return -1
    return left - 1 if nums[left - 1] == target else -1


if __name__ == '__main__':
    assert 4 == binarySearchMostLeft([2, 3, 4, 5, 6, 6, 6, 12], 6)
    assert 5 == binarySearchMostRight([3, 4, 5, 6, 6, 6, 9, 12], 6)
    assert 0 == binarySearchMostLeft([1], 1)
    assert 0 == binarySearchMostRight([1], 1)
