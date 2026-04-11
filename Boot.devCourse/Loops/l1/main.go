package main

func bulkSend(numMessages int) float64 {
	total_cost :=0.0
	for i :=0; i < numMessages ; i++{
		total_cost += float64(i)*0.01+1
	}
	return  total_cost
}
