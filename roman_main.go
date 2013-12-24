package main

import roman "github.com/camerazn/romanNumerals/romanNumerals"
import ("fmt"
		"flag")

func main(){
	toA := ""
	flag.StringVar(&toA, "toA", "", "Roman Numeral to translate to Integer")
	flag.Parse()

fmt.Println(roman.ToArabic(toA))
}
