package Week_05

func findWords(board [][]byte, words []string) []string {
	m := len(board)
	n := len(board[0])
	res := []string{}
	maxLen := 0
	h := map[string]bool{}

	for _, word := range words {
		h[word] = true
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			recursion(board, "", i, j, maxLen, h, &res)
		}
	}

	return res
}

func recursion(board [][]byte, w string, i, j, maxLen int, h map[string]bool, res *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == '#' {
		return
	}

	w = w + string(board[i][j])

	if len(w) > maxLen {
		return
	}

	if _, has := h[w]; has {
		*res = append(*res, w)
		delete(h, w)
		maxLen = 0
		for word, _ := range h {
			if len(word) > maxLen {
				maxLen = len(word)
			}
		}
	}

	tmp := board[i][j]
	board[i][j] = '#'
	recursion(board, w, i-1, j, maxLen, h, res)
	recursion(board, w, i+1, j, maxLen, h, res)
	recursion(board, w, i, j-1, maxLen, h, res)
	recursion(board, w, i, j+1, maxLen, h, res)
	board[i][j] = tmp
}
