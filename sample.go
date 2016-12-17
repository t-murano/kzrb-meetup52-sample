package main

import "fmt"

func kzrb(num int) string {
	return fmt.Sprintf("kzrb meetup%d", num)
}

func main() {
	var num = 52
	fmt.Println(kzrb(num))
}
