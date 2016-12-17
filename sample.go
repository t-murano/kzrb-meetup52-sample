package main

import (
	"fmt"
	"strings"
)

func beer(count int) string {
	return strings.Repeat("🍺", count)
}

func kzrb(num int) string {
	return fmt.Sprintf("kzrb meetup%d", num)
}

func pizza(count int) string {
	return strings.Repeat("🍕", count)
}

func main() {
	var num = 52
	fmt.Println(beer(num))
	fmt.Println(kzrb(num))
	fmt.Println(pizza(num))
}
