package main

import "fmt"

// 整数拆分，将一个整数拆分成乘积最大的
func integerBreak(n int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return 1
	} else if n <= 3 {
		return n - 1
	}
	dp[1], dp[2], dp[3] = 1, 2, 3
	for i := 4; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// 尝试所有的剪法
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}
	return dp[n]
}
func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	} else {
		return n1
	}
}

// 剪绳子
func cuttingRope(n int) int {
	dp := make([]int, n+1)
	if n < 2 {
		return 1
	} else if n <= 3 {
		return n - 1
	}
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i; j++ {
			// 尝试所有的剪法
			dp[i] = max(dp[i], j*max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

// 买卖股票的最佳时机
func maxProfit1(prices []int) int {
	// 前i天的最大收益 = max{前i-1天的最大收益，第i天的价格-前i-1天中的最小价格}
	if len(prices) <= 1 {
		return 0
	}
	minNum, maxNum := prices[0], 0
	for i := 1; i < len(prices); i++ {
		maxNum = max(maxNum, prices[i]-minNum)
		minNum = min(minNum, prices[i])
	}
	return maxNum
}

// 买卖股票的最佳时Ⅱ-多次买卖
func maxProfit2(prices []int) int {
	//只要今天比昨天大，就卖
	lastNum, maxNum := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-lastNum > 0 {
			maxNum += prices[i] - lastNum
		}
		lastNum = prices[i]
	}
	return maxNum
}

// 买卖股票的最佳时Ⅲ-最多可以卖两次
func maxProfit3(prices []int) int {
	n := len(prices)
	buy1, profit1 := prices[0], 0
	buy2, profit2 := prices[0], 0
	for i := 1; i < n; i++ {
		buy1 = min(buy1, prices[i])
		profit1 = max(profit1, prices[i]-buy1)
		buy2 = min(buy2, prices[i]-profit1) // 第二次买入要减去第一次赚的利润,相当于降低第二次买入的成本了
		profit2 = max(profit2, prices[i]-buy2)
	}
	return profit2
}

func main() {
	// 数字拆分
	result := integerBreak(6)
	fmt.Println(result)
	fmt.Println(cuttingRope(6))
}
