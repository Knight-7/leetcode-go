/*
* @Author: Knight
* @Date:   2020/8/16 8:14
* @Email:  knight2347@163.com
* @Ideal:  leetcode图像渲染，简单的dfs
 */

package DFS

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	dfsNewColor(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfsNewColor(image [][]int, sr, sc, oldColor, newColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = newColor
	dfsNewColor(image, sr-1, sc, oldColor, newColor)
	dfsNewColor(image, sr+1, sc, oldColor, newColor)
	dfsNewColor(image, sr, sc-1, oldColor, newColor)
	dfsNewColor(image, sr, sc+1, oldColor, newColor)
}
