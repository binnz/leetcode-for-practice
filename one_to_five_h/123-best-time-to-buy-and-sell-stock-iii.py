"""
    Leetcode 121
    https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/


"""


def maxProfit(prices):
    dp = [[[0, 0] for _ in range(3)] for _ in range(len(prices) + 1)]
    for i in range(3):
        dp[0][i][0], dp[0][i][1] = 0, float('-inf')
    for i in range(1, len(prices) + 1):
        for j in range(1, 3):
            dp[i][j][0] = max(dp[i - 1][j][0], dp[i - 1][j][1] + prices[i - 1])
            dp[i][j][1] = max(dp[i - 1][j][1], dp[i - 1]
                              [j - 1][0] - prices[i - 1])
    return dp[len(prices)][2][0]
