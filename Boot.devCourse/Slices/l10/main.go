package main

func sum(nums ...int) int {
	total_cost:=0
	for i:=0;i<len(nums);i++{
		total_cost+=nums[i]
	}
	return total_cost
}
