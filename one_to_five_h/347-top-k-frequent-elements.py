"""
    Leetcode 347
    https://leetcode-cn.com/problems/top-k-frequent-elements/

"""


def topKFrequent(nums, k):
    result_map = {}
    for e in nums:
        if e not in result_map:
            result_map[e] = 1
        else:
            result_map[e] += 1
    res = []
    for key, val in result_map.items():
        res.append((key, val))
        build_heap(res, k)
    for i in range(k, len(res)):
        if res[0][1] < res[i][1]:
            res[0] = res[i]
            heapify(res, 0, k)
    return [x[0] for x in res[:k]]


def build_heap(nums, k):
    for i in range((k - 1) // 2, -1, -1):
        heapify(nums, i, k)


def heapify(nums, i, l):
    left = 2 * i + 1
    right = 2 * i + 2
    minv = i
    if left < l and nums[left][1] < nums[minv][1]:
        minv = left
    if right < l and nums[right][1] < nums[minv][1]:
        minv = right
    if minv != i:
        nums[i], nums[minv] = nums[minv], nums[i]
        heapify(nums, minv, l)
