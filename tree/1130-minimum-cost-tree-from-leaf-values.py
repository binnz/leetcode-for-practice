"""
    Leetcode 1130
    https://leetcode-cn.com/problems/minimum-cost-tree-from-leaf-values/

"""


def mctFromLeafValues(arr):
    memo = {}

    def find(left, right):
        if (left, right) in memo:
            return memo[(left, right)]
        if right - left <= 1:
            return 0
        res = float('inf')
        for i in range(left + 1, right):
            mid = find(left, i) + find(i, right) + \
                max(arr[left:i]) * max(arr[i:right])
            res = min(res, mid)
        memo[(left, right)] = res
        return res
    find(0, len(arr))
    return memo[(0, len(arr))]


def mctFromLeafValues2(arr):
    stack = [float('inf')]
    cost = 0
    for e in arr:
        while stack[-1] < e:
            mid = stack.pop()
            cost += mid * min(stack[-1], e)
        stack.append(e)
    while len(stack) > 2:
        mid = stack.pop()
        cost += stack[-1] * mid
    return cost
