package main

import (
	"fmt"
	"github.com/delgoden/calculator/internal/calculation"
	"github.com/delgoden/calculator/internal/models"
	"github.com/delgoden/calculator/internal/myErrors"
	"github.com/delgoden/calculator/internal/myParser"
	"github.com/delgoden/calculator/internal/reader"
	"github.com/delgoden/calculator/internal/validation"
)

func main() {
	for {
		fmt.Println("Enter two numbers from 1(I) to 10(X)\nand a mathematical operation (+ - * /) separated by a space\nexample: 8 + 9(VIII + IX)\nto complete the work, enter 'end'")
		data := reader.ReadData()
		if data == "end" {
			fmt.Println("GoodBye")
			break
		}
		mathData, err := validation.Valid(data)
		if err != nil {
			fmt.Println(err)
			break
		}
		dataForCalc, err := myParser.ParsData(mathData)
		if err != nil {
			fmt.Println(err)
			break
		}
		dataForAnswer := calculation.Calculation(dataForCalc)
		if dataForAnswer.TypeNum == models.Arabic {
			fmt.Println(dataForAnswer.Num)
		} else if dataForAnswer.TypeNum == models.Roman {
			if dataForAnswer.Num < 1 {
				fmt.Println(myErrors.RomanZeroNegativeError)
			} else {
				answer := arabicToRoman(dataForAnswer.Num)
				fmt.Println(answer)
			}
		}
		fmt.Println()
	}
}

func arabicToRoman(data int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for data >= conversion.value {
			roman += conversion.digit
			data -= conversion.value
		}
	}
	return roman
}
