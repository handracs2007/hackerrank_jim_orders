// https://www.hackerrank.com/challenges/jim-and-the-orders/problem

package main

import (
	"fmt"
	"sort"
)

type CustomerOrder struct {
	ID          int32
	OrderNumber int32
	PrepTime    int32
}

func (co CustomerOrder) ServeTime() int32 {
	return co.OrderNumber + co.PrepTime
}

func jimOrders(orders [][]int32) []int32 {
	var customerOrders = make([]CustomerOrder, len(orders))

	for i := 0; i < len(orders); i++ {
		customerOrders[i] = CustomerOrder{
			ID:          int32(i + 1),
			OrderNumber: orders[i][0],
			PrepTime:    orders[i][1],
		}
	}

	// Now, let's sort the information based on the serving time, then ID
	sort.Slice(customerOrders, func(i, j int) bool {
		if customerOrders[i].ServeTime() != customerOrders[j].ServeTime() {
			return customerOrders[i].ServeTime() < customerOrders[j].ServeTime()
		} else {
			return customerOrders[i].ID < customerOrders[j].ID
		}
	})

	// Now let's build the result
	var result = make([]int32, len(customerOrders))

	for idx, customerOrder := range customerOrders {
		result[idx] = customerOrder.ID
	}

	return result
}

func main() {
	fmt.Println(jimOrders([][]int32{{8, 3}, {5, 6}, {6, 2}, {2, 3}, {4, 3}}))
}
