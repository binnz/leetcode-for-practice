"""
    Leetcode 122
    https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

Say you have an array prices for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).


dp[i][0] 第i天没有持有股票的利润
dp[i][1] 第i天持有股票的利润
"""


def maxProfit(prices):
    dp = [[0, 0] for _ in range(len(prices))]
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    for i, price in enumerate(prices):
        if i == 0:
            continue
        dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + price)
        dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - price)
    return dp[len(prices) - 1][0]


def maxProfit2(prices):
    last_hold_cash = 0
    last_hold_stock = -prices[0]
    for i in range(1, len(prices)):
        last_hold_cash = max(last_hold_cash, last_hold_stock + prices[i])
        last_hold_stock = max(last_hold_stock, last_hold_cash - prices[i])
    return last_hold_cash
