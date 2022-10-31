package main

import (
	"fmt"
	"gentoo-packages-bot/scraping"
)

func main() {
	fl := scraping.GetUseFlag("claasdasdg")

	fmt.Println(fl)
}
