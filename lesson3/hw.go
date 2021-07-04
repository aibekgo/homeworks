package main

import "fmt"

func main() {
	var s, e, sum1, sum2 int
	fmt.Scanf("%d %d", &s, &e)
	for i:=s; i<e; i++ {
		if i%2!=0 {sum1 = sum1+i}
		if i%2==0 {sum2 = sum2+i}


	}
	fmt.Println(sum1, sum2)
}
/*вам дается промежуток чисел от s до e
найти сумму четных числе от s до e
найти сумму нечетных чисел от s до e*/