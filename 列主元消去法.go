package main

import (
	"fmt"
	"math"
)

//以下程序为了和通常所说的矩阵第一行第一列适应，舍去第零行第零列

func main() {
	var (
		n int
		A [20][20]float64
		x [20]float64
		b [20]float64
	)
	fmt.Println("dimension")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("A[%d][%d] = ", i, j)
			fmt.Scanln(&A[i][j])
		}
		fmt.Printf("b[%d] = ", i)
		fmt.Scanln(&b[i])
	}

	Eliminate(n, &A, &b)
	Substitude(n, &A, &b, &x)
	for i := 1; i <= n; i++ {
		println(x[i])
	}
}

func Swap(n int, A *[20][20]float64, b *[20]float64, r, k int) {
	// 交换增广矩阵的 r 行和 k 行
	(*b)[r], (*b)[k] = (*b)[k], (*b)[r]
	for i := 1; i <= n; i++ {
		(*A)[r][i], (*A)[k][i] = (*A)[k][i], (*A)[r][i]
	}
}

func Eliminate(n int, A *[20][20]float64, b *[20]float64) {
	for k := 1; k <= n; k++ {
		// 选择主元
		tmp := k
		for r := k + 1; r <= n; r++ {
			if math.Abs((*A)[r][k]) > math.Abs((*A)[tmp][k]) {
				tmp = r
			}
		}
		//若主元与操作行不同，交换两行
		if tmp != k {
			Swap(n, A, b, tmp, k)
		}
		//消元
		for r := k + 1; r <= n; r++ {
			factor := (*A)[r][k] / (*A)[k][k]
			for j := k; j <= n; j++ {
				(*A)[r][j] -= factor * (*A)[k][j]
			}
			(*b)[r] -= factor * (*b)[k]
		}
	}
}

func Substitude(n int, A *[20][20]float64, b *[20]float64, x *[20]float64) {
	// 回代
	for i := n; i >= 1; i-- {
		sum := 0.0
		for j := i + 1; j <= n; j++ {
			sum += (*A)[i][j] * (*x)[j]
		}
		(*x)[i] = ((*b)[i] - sum) / (*A)[i][i]
	}
}
