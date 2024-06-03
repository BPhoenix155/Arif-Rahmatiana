package main

import "fmt"

func berlibur(naikgaji, dapetbonus bool) {
	if naikgaji == true || dapetbonus == true {
		fmt.Println("berlibur")
	} else {
		fmt.Println("tidak berlibur")
	}
}

func main() {
	var naikGaji, dapatBonus bool
	fmt.Scan(&naikGaji, &dapatBonus)
	berlibur(naikGaji, dapatBonus)
}
