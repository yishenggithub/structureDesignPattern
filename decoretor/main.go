package main

import "fmt"

// 装饰是一种结构设计模式， 允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。

func main() {
	pizza := &veggeMania{}

	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
