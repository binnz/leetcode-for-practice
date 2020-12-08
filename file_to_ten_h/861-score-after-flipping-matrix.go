package main

func matrixScore(A [][]int) int {
    m, n := len(A), len(A[0])
    res := 1<< (n-1) * m

    for j:=1;j<n; j++{
        one := 0
        for _,e := range A{
            if e[j] ==e[0]{
                one++
            }
        }
        if one <=m/2{
            one = m-one
        }
        res += 1<<(n-j-1)*one
    }
    return res
}
