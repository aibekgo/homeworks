package main
import "fmt"
func main() {
	var a,b,c int
	var res string
	fmt.Scanf("%d %d %d",&a, &b, &c)
	if a==b && b==c && a==c  { res = " Равносторонний треугольник"
	} else if a!=b && b!=c && a!=c { res = "Разносторонний треугольник"
	} else { res = " Равнобедренный треугольник"}
	fmt.Printf("%s", res)
	}