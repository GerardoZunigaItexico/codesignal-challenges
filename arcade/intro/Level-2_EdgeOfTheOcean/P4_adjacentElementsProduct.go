package main

func adjacentElementsProduct(inputArray []int) int {
	max := inputArray[0]*inputArray[1]
	for i:=1; i<len(inputArray)-1; i++{
		if newMax := inputArray[i]*inputArray[i+1]; newMax > max {
			max = newMax
		}
	}
	return max
}


func adjacentElementsProductRecursive(inputArray []int) int {
	if len(inputArray) == 2{
		return inputArray[0]*inputArray[1]
	}
	max := inputArray[0]*inputArray[1]
	prod := adjacentElementsProduct(inputArray[1:])
	if prod > max{
		max = prod
	}
	return max
}

func adjacentElementsProductRecursive2(inputArray []int) int {
	if len(inputArray) == 2{
		return inputArray[0]*inputArray[1]
	}
	max := inputArray[0]*inputArray[1]
	prod := adjacentElementsProduct(inputArray[1:])
	if max < prod {
		return prod
	}
	return prod
}