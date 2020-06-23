"""
    Leetcode 384
    https://leetcode-cn.com/problems/shuffle-an-array/
"""

import random


class Solution:

    def __init__(self, nums):
        self.origin = nums
        self.data = self.origin[::]

    def reset(self):
        """
        Resets the array to its original configuration and return it.
        """
        return self.origin

    def shuffle(self):
        """
        Returns a random shuffling of the array.
        """
        for i in range(len(self.data)):
            index = random.randint(i, len(self.data) - 1)
            self.data[i], self.data[index] = self.data[index], self.data[i]
        return self.data
