package main

import "fmt"

const N int = 5

var (
	
	x1= [N]int{1, 2, 3, 4, 5,}
	y1= [N]int{1, 3, 6, 3, 1,}
	
	x2= [N]int{1, 0, -1, -1, 1,}
	y2= [N]int{1, 3, 1, -1, -1,}
	
	x3= [N]int{0, 0, 0, 0, 0,}
	y3= [N]int{1, 6, 7, 12, 40,}

)

func InData(x[N] int, y[N] int) {
	fmt.Println("== INPUT ==")
	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
}

func OutData(ans bool){
    fmt.Println("== OUTPUT ==\n", ans)
}

func CheckingValue(x[N] int, y[N] int, a int, b int, c int) int {
	turn := (x[b] - x[a])*(y[c] - y[b]) - (y[b] - y[a])*(x[c] - x[b])
	return turn
}

func Algoritm(x[N] int, y[N] int) bool {
	var turn int = CheckingValue(x, y, N-1, 0, 1)
	var i int = 1
	for ; i < N-1; i++{
	    turn = CheckingValue(x, y, i-1, i, i+1)
	    if turn != 0 { break }
	}
	
	var ans bool = false
	    for ; i<N-1; i++{
	        ans = CheckingValue(x, y, i-1, i, i+1)*turn >=0
	        if !ans {break}
	    }
	
	return ans
}

func main() {

    fmt.Println("~~~ TESTS ~~~")
    
	InData(x1, y1)
	OutData(Algoritm(x1, y1))
	
	InData(x2, y2)
	OutData(Algoritm(x2, y2))
	
	InData(x3, y3)
	OutData(Algoritm(x3, y3))
}