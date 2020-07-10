"""
    Leetcode 3
    https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

"""


def lengthOfLongestSubstring(s):
    start, res = 0, 0
    m = {}
    for i, e in enumerate(s):
        if e not in m:
            m[e] = i
        else:
            start = max(m[e] + 1, start)
            m[e] = i
        res = max(res, i - start + 1)
    return res
