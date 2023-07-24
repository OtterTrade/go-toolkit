# go-toolkit
为 OtterTrade 整理和编写的工具包

## 工具包介绍
### arith
数据基本类型定义与实现 [背景介绍](https://qv7420ojfw2.feishu.cn/docx/ORWcdVZNDok7PexdCSMc2bbZnmc?from=from_lark_group_search)

Number：抽象计算指标使用的通用接口，具体有4个实现。

1. OtNumber 内部封装了Float64和expFloat64两种实现方式，实现了根据数据范围动态调整实现， 对于极小数或极大数使用expFloat64实现，对于float64正常精度表示的数使用Float64实现。
2. Float64 float64基础类型的封装，用于float64基础类型到Number接口的转换。
3. expFloat64 科学计数法表示方式，不对外暴露，作为OtNumber内部一种实现。
4. decimalNumber 基于decimal库的一种实现，目前主要用来测试OtNumber精度。

arith库对外暴露OtNumber和Float64两种类型，OtNumber主要用于保存价格数据以及和价格相关的中间指标，可以以16位有效数字表示任意范围的浮点数，同时OtNumber实现了Json和Bson的序列化和反序列化函数，可以在json或mongodb定义结构中直接使用；Float64用于正常数据范围内的浮点数表示。如计算两个价格平均值：

```golang
package main

import (
	"go-toolkit/arith"
)

type kline struct {
	Open   arith.OtNumber `json:"open" bson:"open"`
	High   arith.OtNumber `json:"high" bson:"high"`
	Low    arith.OtNumber `json:"low" bson:"low"`
	Close  arith.OtNumber `json:"close" bson:"close"`
	Volume arith.OtNumber `json:"volume" bson:"volume"`
}
// avg_price 计算价格平均值
func avg_price(prices ...arith.OtNumber) arith.OtNumber {
	r := arith.OtNumber{}
	cnt := 0.
	for _, v := range(prices) {
		cnt += 1
		r = r.Add(v)
	}
	return r.Div(arith.Float64(cnt))
}

```
