'''
    Leetcode 910
    https://leetcode-cn.com/problems/smallest-range-ii/submissions/

'''


def smallestRangeII(A, K):
    A.sort()
    if len(A) < 2:
        return 0
    res = A[-1] - A[0]
    for i in range(1, len(A)):
        minval = min(A[0] + K, A[i] - K)
        maxval = max(A[i - 1] + K, A[-1] - K)
        res = min(maxval - minval, res)
    return res
