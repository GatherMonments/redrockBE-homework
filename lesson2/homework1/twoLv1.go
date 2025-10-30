package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) TotalValue() float64 {
	return p.Price * float64(p.Stock)
}

func (p Product) IsInStock() bool {
	return p.Stock > 0
}

func (p Product) Info() string {
	return fmt.Sprintf("商品: %s, 单价: %.2f, 库存: %d件", p.Name, p.Price, p.Stock)
}

func (p *Product) Restock(amount int) {
	p.Stock += amount
}

func (p *Product) Sell(amount int) (success bool, message string) {

	if amount > p.Stock {
		return false, "库存不足"
	} else {
		p.Stock -= amount
		return true, "售卖成功"
	}
}

func main() {
	str := ""
	pro := Product{
		Name:  "Go编程书",
		Price: 89.5,
		Stock: 10,
	}
	success, _ := pro.Sell(5)
	if success {
		str = "成功"
	} else {
		str = "失败"
	}
	fmt.Printf("售卖5本:%s, 剩余库存:%d\n", str, pro.Stock)
	pro.Restock(20)
	fmt.Printf("进货20本, 当前库存:%d\n", pro.Stock)
	success1, message := pro.Sell(30)
	if success1 {
		str = "成功"
	} else {
		str = "失败"
	}
	fmt.Printf("售卖20本:%s, %s\n", str, message)
	fmt.Println("商品信息:")
	fmt.Println(pro.Info())
	fmt.Printf("库存总价值:%.2f\n", pro.TotalValue())
}

// 预期输出:
// 售卖5本: 成功, 剩余库存: 5
// 进货20本, 当前库存: 25
// 售卖30本: 失败, 库存不足
//
// 商品信息:
// 商品: Go编程书, 单价: ¥89.5, 库存: 25件
// 库存总价值: ¥2237.50
