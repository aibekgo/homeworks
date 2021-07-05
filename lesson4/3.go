package main
import "fmt"
func main() {
	arr := []int{}
	indexs := []int{}
	value := []int{}
	n:= 30
	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			value = append(value, arr[i])
			indexs = append(indexs, i)
		}
		}
	    fmt.Println(value)
		fmt.Println(indexs)


}