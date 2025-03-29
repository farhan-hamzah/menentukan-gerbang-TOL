package main
import "fmt"

func gerbangTol(tol string, a, b, c *int){
	if tol == "A"{
		*a++
	}else if tol == "B"{
		*b++
	}else if tol == "C"{
		*c++
	}
}

func main(){
	var tol string
	var a, b, c int
	fmt.Scan(&tol)
	gerbangTol(tol, &a, &b, &c)
	for tol >= "A" && tol <= "C"{
		fmt.Scan(&tol)
		gerbangTol(tol, &a, &b, &c)
	}
	fmt.Println("Tipe A = ", a)
	fmt.Println("Tipe B = ", b)
	fmt.Println("Tipe C = ", c)
}