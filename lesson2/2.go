package main
import "fmt"
func main() {
	var a, b, c, max, min, avr int
	var res, res1, res2 string
	res = " Максимальное число"
	res1 = " Минимальное число"
	res2 = " Среднее значение"
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if        a>=b && b>=c { max=a
	} else if a<=b && b<=c { max=c
	} else if a<=b && b>=c { max=b
	} else if a<=b && b<=c { min=a
	} else if a<=b && b<=c { min=c
	} else if a>=b && b>=c { min=b
	} else if a>=b && a<=c { avr=a
	} else if a<=b && b<=c { avr=b
	} else if a<=c && b>=c { avr=c
	}
	fmt.Printf("%s %d %s %d %s %d", res, max, res1, min, res2, avr )
}

/*вводятся a,b,c
найти маскимальное из чисел
найти минимальное из чисел
среднее число это то число которое больше первого но меньше второго*/