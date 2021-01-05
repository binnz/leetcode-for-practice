'''
    Leetcode 97
    https://leetcode-cn.com/problems/interleaving-string/
'''


def isInterleave(a, b, s):
    m = len(a)
    n = len(b)
    if m + n != len(s):
        return False
    dp = [[False for _ in range(n + 1)] for _ in range(m + 1)]

    dp[0][0] = True
    for i in range(1, m + 1):
        dp[i][0] = dp[i - 1][0] and a[i - 1] == s[i - 1]
    for i in range(1, n + 1):
        dp[0][i] = dp[0][i - 1] and b[i - 1] == s[i - 1]
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            dp[i][j] = (a[i - 1] == s[i + j - 1] and dp[i - 1][j]
                        ) or (b[j - 1] == s[i + j - 1] and dp[i][j - 1])

    return dp[m][n]


def isInterleave2(a, b, s):
    m = len(a)
    n = len(b)
    if m + n != len(s):
        return False
    if not a:
        return b == s
    if not b:
        return a == s
    res = False
    if a[0] == s[0]:
        res = isInterleave2(a[1:], b, s[1:])
    if b[0] == s[0]:
        res = res or isInterleave2(a, b[1:], s[1:])
    return res
