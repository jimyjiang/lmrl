package arithmetic

import (
	"fmt"
	"math"
)

func printDP(dp [][]float64) {
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			fmt.Printf("%d\t", int(dp[i][j]))
		}
		fmt.Println()
	}
}
func LongCommonSubstring(s1, s2 string) float64 {
	if len(s1) == 0 || len(s2) == 0 {
		return 0
	}
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	m := len(runes1)
	n := len(runes2)
	var dp [][]float64
	dp = make([][]float64, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]float64, n+1)
	}
	// printDP(dp)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if runes1[i-1] == runes2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = math.Max(float64(dp[i-1][j]), float64(dp[i][j-1]))
			}
		}
	}
	// printDP(dp)
	return dp[m][n]
}
