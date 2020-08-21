"""
    Leetcode 121
    https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/


"""


def maxProfit(prices):
    res = 0
    minPrice = float('inf')
    for price in prices:
        minPrice = min(minPrice, price)
        res = max(res, price - minPrice)
    return res
