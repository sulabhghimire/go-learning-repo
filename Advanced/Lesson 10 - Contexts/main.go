package main

import (
	"context"
	"fmt"
	"time"
)

func checkEvenOdd(ctx context.Context, num int) string {

	select {
	case <-ctx.Done():
		return "Operation cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is Even", num)
		} else {
			return fmt.Sprintf("%d is Odd", num)
		}
	}
}

func main() {

	ctx := context.TODO()

	result := checkEvenOdd(ctx, 10)
	fmt.Println("Result:", result)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second) // timer of context starts here
	defer cancel()

	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result from timeout ctx:", result)

	time.Sleep(2 * time.Second) // Simulate some delay
	result = checkEvenOdd(ctx, 32)
	fmt.Println("Result from timeout ctx:", result)
}

// // === Difference between context.TODO and context.Background ===
// func main() {

// 	// Context are kind of objects, i.e. instances of structs
// 	toDoContext := context.TODO()
// 	contextBkg := context.Background()

// 	// Will result in a new context and needs to be passed a existing parent context
// 	ctx := context.WithValue(toDoContext, "name", "Sulabh")
// 	fmt.Println(ctx)               // Output: context.TODO.WithValue("name", "Sulabh")
// 	fmt.Println(ctx.Value("name")) // Output: Sulabh

// 	ctx1 := context.WithValue(contextBkg, "city", "Kathmandu")
// 	fmt.Println(ctx1)               // Output: context.Background.WithValue("name", "Kathmandu")
// 	fmt.Println(ctx1.Value("city")) // Output: Kathmandu

// }
