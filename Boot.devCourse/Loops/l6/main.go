package main

func countConnections(groupSize int) int {
	if groupSize<=1{
		return 0
	
	}else{
		connections:=1
		for i:=2;i<groupSize;i++{
			connections += i
		}
		return  connections
	}
}
