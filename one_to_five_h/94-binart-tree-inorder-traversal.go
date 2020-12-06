package main

import (
  "fmt"
)

type TreeNode struct{
  Val int
  Left *TreeNode
  Right * TreeNode
}

func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    stack:= []*TreeNode{}
    visit := []bool{}
    if root ==nil{
        return res
    }
    stack =append(stack, root)
    visit = append(visit, false)
    for len(stack)!= 0 {
        node:=stack[len(stack)-1]
        visited :=visit[len(visit)-1]
        stack=stack[:len(stack)-1]
        visit=visit[:len(visit)-1]
        if visited{
            res= append(res, node.Val)
        }else{
           if node.Right != nil{
            stack= append(stack, node.Right)
            visit=append(visit,false)
           }
            stack=append(stack, node)
            visit= append(visit,true)
           if node.Left !=nil{
            stack=append(stack,node.Left)
            visit=append(visit,false)
           }
        }
    }
    return res
}

