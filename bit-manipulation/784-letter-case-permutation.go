package main

func letterCasePermutation(S string) []string {
	var (
		res    []string
		dfs    func(idx int, path []byte)
		str    = []byte(S)
		length = len(S)
	)

	dfs = func(idx int, path []byte) {
		if idx == length {
			res = append(res, string(path))
			return
		}
		dfs(idx+1, path)
		letter := path[idx]
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			path[idx] ^= 32
			dfs(idx+1, path)
		}
	}

	dfs(0, str)
	return res

}
