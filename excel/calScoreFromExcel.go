package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
)

type Item struct {
	AllScore decimal.Decimal `json:"allScore"`
	Item     int             `json:"item"`
	Score    decimal.Decimal `json:"score"`
}

func main() {
	// JSON数据
	jsonData := `[{"allScore":29,"item":1,"score":23.5},{"allScore":11,"item":2,"score":8.5},{"allScore":22.5,"item":3,"score":20.5},{"allScore":21.5,"item":4,"score":16.05},{"allScore":16,"item":5,"score":12}]`

	// 解析JSON数据
	var items []Item
	err := json.Unmarshal([]byte(jsonData), &items)
	if err != nil {
		log.Fatalf("解析JSON数据失败: %s", err)
	}

	// 计算总分
	totalScore := decimal.NewFromFloat(0)
	for _, item := range items {
		totalScore = totalScore.Add(item.Score)
	}

	// 计算合格率
	passRate := totalScore.Div(decimal.NewFromInt(100))

	// 输出结果
	fmt.Printf("总分: %s\n", totalScore)
	fmt.Printf("合格率: %.2f%%\n", passRate.Mul(decimal.NewFromFloat(100)).Round(2).Div(decimal.NewFromInt(100)).BigFloat())
}
