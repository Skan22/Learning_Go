package main

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int,rows)
	data := make([]int,rows*cols)
	for i:=range matrix{
			matrix[i] = data[i*cols:(i+1)*cols] 
			for j:=range(matrix[i]){
				matrix[i][j] = i*j}
	
	}
return matrix
}
