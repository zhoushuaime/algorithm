package leetcode

// 给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。
// 水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。
// 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

// flipAndInvertImage ...
func flipAndInvertImage(A [][]int) [][]int {

	columns := len(A[0])

	for _,row := range A {
		for i:=0;i<(columns+1)/2;i++ {
			temp := row[i] ^ 1 // 0异或任何数不变，1异或任何数取反
			row[i] = row[columns-i-1] ^ 1
			row[columns-i-1] = temp
		}
	}

	return A

}

// 输入: [[1,1,0],[1,0,1],[0,0,0]]
// 输出: [[1,0,0],[0,1,0],[1,1,1]]
// 解释: 首先翻转每一行: [[0,1,1],[1,0,1],[0,0,0]]；
//      然后反转图片: [[1,0,0],[0,1,0],[1,1,1]]
