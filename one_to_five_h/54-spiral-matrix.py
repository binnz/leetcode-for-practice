"""
    Leetcode 54
    https://leetcode-cn.com/problems/spiral-matrix/

"""


def spiralOrder(matrix):
    if len(matrix) == 0 or len(matrix[0]) == 0:
        return []
    x_min = 0
    x_max = len(matrix)
    y_min = 0
    y_max = len(matrix[0])
    direct = 0
    res = []
    i = j = 0
    while len(res) != len(matrix) * len(matrix[0]):
        if i in range(x_min, x_max) and j in range(y_min, y_max):
            res.append(matrix[i][j])
            if direct == 0:
                j += 1
            elif direct == 1:
                i += 1
            elif direct == 2:
                j -= 1
            elif direct == 3:
                i -= 1
            else:
                raise Exception
        else:
            if direct == 0:
                j -= 1
                i += 1
                direct += 1
                x_min += 1
            elif direct == 1:
                i -= 1
                j -= 1
                direct += 1
                y_max -= 1
            elif direct == 2:
                i -= 1
                j += 1
                direct += 1
                x_max -= 1
            elif direct == 3:
                j += 1
                i += 1
                y_min += 1
                direct = 0
            else:
                raise Exception
    return res
