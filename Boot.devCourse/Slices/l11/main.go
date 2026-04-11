package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	daycosts := make([]float64,0)
	for i:=0;i<len(costs);i++{
		if costs[i].day == day {
			daycosts = append(daycosts,costs[i].value)
		}
	}
	return daycosts
}
