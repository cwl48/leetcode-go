package exitssearch

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 *
 * https://leetcode-cn.com/problems/word-search/description/
 *
 * algorithms
 * Medium (34.99%)
 * Total Accepted:    5.6K
 * Total Submissions: 16K
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
 *
 * 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
 *
 * 示例:
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * 给定 word = "ABCCED", 返回 true.
 * 给定 word = "SEE", 返回 true.
 * 给定 word = "ABCB", 返回 false.
 *
 */
var visited [][]bool

func exist(board [][]byte, word string) bool {

	m, n := len(board), len(board[0])
	visited = make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			opp := word[0]
			if board[i][j] == opp && search(board, word, i, j, 0) {
				return true
			}
		}
	}

	return false
}

// count 匹配的次数
func search(board [][]byte, word string, i int, j int, count int) bool {

	if count == len(word) {
		return true
	}
	if i+1 > len(board) || i < 0 || j+1 > len(board[0]) || j < 0 || visited[i][j] || board[i][j] != word[count] {
		return false
	}
	visited[i][j] = true
	if search(board, word, i, j-1, count+1) || search(board, word, i, j+1, count+1) || search(board, word, i+1, j, count+1) || search(board, word, i-1, j, count+1) {
		return true
	}
	visited[i][j] = false
	return false
}
