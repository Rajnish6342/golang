package main

import "math"

func maxProfit(prices []int) int {
	mProfit := 0
	buy := math.MaxInt
	for _, price := range prices {
		if buy > price {
			buy = price
		} else {
			mProfit = max(price-buy, mProfit)
		}
	}
	return mProfit
}
