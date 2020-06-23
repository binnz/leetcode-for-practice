"""
    Leetcode 498
    https://leetcode-cn.com/problems/diagonal-traverse/
"""
"""
Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

Output:  [1,2,4,7,5,3,6,8,9]
"""

def findDiagonalOrder(matrix):
    from itertools import groupby

    all_ele = []
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            all_ele.append([i + j, i, matrix[i][j]])
    all_ele.sort()
    res = []
    i_add_j = 0
    for key, group in groupby(all_ele, key=lambda _: _[0]):
        group = [_ for _ in group]
        if i_add_j % 2 == 0:
            group = group[::-1]
        res.extend([_[2] for _ in group])
        i_add_j += 1
    return res
