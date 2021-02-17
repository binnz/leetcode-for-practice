package main

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	var (
		dfs  func(*TreeNode)
		dfss func(*TreeNode, map[*TreeNode]struct{}, int)
		res  []int
		mapp map[*TreeNode]*TreeNode = map[*TreeNode]*TreeNode{}
		used map[*TreeNode]struct{}  = map[*TreeNode]struct{}{}
	)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		if tn.Left != nil {
			mapp[tn.Left] = tn
			dfs(tn.Left)
		}
		if tn.Right != nil {
			mapp[tn.Right] = tn
			dfs(tn.Right)
		}
	}
	dfss = func(tn *TreeNode, used map[*TreeNode]struct{}, d int) {
		_, ok := used[tn]
		if ok {
			return
		}
		if d == 0 {
			res = append(res, tn.Val)
			return
		}
		used[tn] = struct{}{}
		if tn.Left != nil {
			dfss(tn.Left, used, d-1)
		}
		if tn.Right != nil {
			dfss(tn.Right, used, d-1)
		}
		if e, ok := mapp[tn]; ok {
			dfss(e, used, d-1)
		}
	}
	dfs(root)
	dfss(target, used, K)
	return res
}
