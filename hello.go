package main

import "fmt"
import "time"
import "math/rand"
import "math"

var i, j int = 1, 3

func main() {
	fmt.Printf("hello, world\n")
	fmt.Println("The time is", time.Now())
    fmt.Println("Random:", rand.Intn(10))
    fmt.Println("Pi:", math.Pi)

    fmt.Println(add(22, 333))
}

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string){
    return y, x;
}