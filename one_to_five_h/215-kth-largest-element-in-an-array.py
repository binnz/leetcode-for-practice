"""
    Leetcode 215
    https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

"""


def findKthLargest(nums, k):
    if not nums:
        return None
    if len(nums) < 2:
        return nums[0]
    n = len(nums)
    build_heap(nums, k)
    for i in range(k, n):
        if nums[i] > nums[0]:
            nums[i], nums[0] = nums[0], nums[i]
            heapify(0, nums, k)
    return nums[0]


def build_heap(nums, n):
    for i in range((n - 1) // 2, -1, -1):
        heapify(i, nums, n)


def heapify(i, nums, n):
    left = 2 * i + 1
    right = 2 * i + 2
    maxi = i
    if left < n and nums[left] < nums[maxi]:
        maxi = left
    if right < n and nums[right] < nums[maxi]:
        maxi = right
    if maxi != i:
        nums[i], nums[maxi] = nums[maxi], nums[i]
        heapify(maxi, nums, n)
