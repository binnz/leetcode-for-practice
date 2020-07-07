"""
    Leetcode 33
    https://leetcode-cn.com/problems/search-in-rotated-sorted-array/

"""


def search(nums, target):
    length = len(nums)
    left = 0
    right = length - 1
    if not nums:
        return -1
    while left < right:
        mid = (left + right) >> 1
        if nums[mid] == nums[right]:
            right -= 1
        elif nums[mid] > nums[right]:
            left = mid + 1
        elif nums[mid] < nums[right]:
            right = mid
    assert left == right
    if nums[0] == target:
        return 0
    if left == 0:
        return find(nums, 0, length, target)
    if target < nums[0]:
        return find(nums, left, length, target)
    if target > nums[0]:
        return find(nums, 0, left, target)


"""
This method is a common method for find the origin start element index in a rotated sorted array,
it also works with duplicate exists.
So it can be used in other places.

The rotated method describe in https://leetcode-cn.com/problems/search-in-rotated-sorted-array/:

an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

Examples:
    Input: [5, 6, 7, 0, 2, 4]
    Ouput: 3

    Input: [0, 0, 0, 1, 0, 0]
    Output: 4

    Input: [1, 2, 3, 5, 7]
    Output: 0
"""


def findFirstEleIndexInRotatedSortedArray(nums):
    length = len(nums)
    left = 0
    right = length - 1
    if not nums:
        return -1
    while left < right:
        mid = (left + right) >> 1
        if nums[mid] == nums[right]:
            right -= 1
        elif nums[mid] > nums[right]:
            left = mid + 1
        elif nums[mid] < nums[right]:
            right = mid
    assert left == right
    return left


"""
This method is binary search method, return any target element index if has duplicate exists.

The method about binary search return most left or right target element index can be found in method binary-search.py.
"""


def find(nums, left, right, target):
    while left < right:
        mid = (left + right) >> 1
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid
    assert left == right
    if left == len(nums):
        return -1
    return left if nums[left] == target else -1


if __name__ == '__main__':
    assert 4 == findFirstEleIndexInRotatedSortedArray([4, 5, 6, 7, 0, 1, 3])
    assert 3 == findFirstEleIndexInRotatedSortedArray([0, 0, 1, 0, 0])
