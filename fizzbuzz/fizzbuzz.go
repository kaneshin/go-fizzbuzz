package fizzbuzz

import (
	"fmt"
)

func PrintFizzBuzz(n int) {
	res := FizzBuzz(n)
	for _, v := range res {
		fmt.Println(v)
	}
}

// Pattern 1.
func FizzBuzz(n int) []interface{} {
	ret := make([]interface{}, n)
	for i := 1; i <= n; i++ {
		var d interface{} = i
		if i%15 == 0 {
			d = "FizzBuzz"
		} else if i%5 == 0 {
			d = "Buzz"
		} else if i%3 == 0 {
			d = "Fizz"
		}
		ret[i-1] = d
	}
	return ret
}

// Pattern 2. switch statement
// func FizzBuzz(n int) []interface{} {
// 	ret := make([]interface{}, n)
// 	for i := 1; i <= n; i++ {
// 		switch {
// 		case i%15 == 0:
// 			ret[i-1] = "FizzBuzz"
// 		case i%5 == 0:
// 			ret[i-1] = "Buzz"
// 		case i%3 == 0:
// 			ret[i-1] = "Fizz"
// 		default:
// 			ret[i-1] = i
// 		}
// 	}
// 	return ret
// }
