package myParser

import (
	"fmt"
	"github.com/delgoden/calculator/internal/models"
	"github.com/delgoden/calculator/internal/myErrors"
	"strconv"
	"strings"
)

func ParsData(data models.MathData) (dataForCalc models.DataForCalc, err error) {
	if data.TypeNum == models.Arabic {
		dataForCalc, err = parsArabic(data)
	}
	if data.TypeNum == models.Roman {
		dataForCalc, err = parsRoman(data)
	}
	return
}

func parsArabic(data models.MathData) (dataForCalc models.DataForCalc, err error) {
	newStr := strings.Replace(data.Data, " ", "", -1)
	dataForCalc.Operation = findOperation(newStr)
	dataForCalc.Nums, err = convStrTONum(strings.Split(newStr, dataForCalc.Operation))
	if err != nil {
		return
	}
	dataForCalc.TypeNum = data.TypeNum
	return
}

func parsRoman(data models.MathData) (dataForCalc models.DataForCalc, err error) {
	newStr := strings.Replace(data.Data, " ", "", -1)
	dataForCalc.Operation = findOperation(newStr)
	dataForCalc.Nums, err = romanToArabic(strings.Split(newStr, dataForCalc.Operation))
	if err != nil {
		return
	}
	dataForCalc.TypeNum = data.TypeNum
	return
}

func findOperation(data string) (operation string) {
	switch {
	case strings.Contains(data, "-"):
		operation = "-"
	case strings.Contains(data, "+"):
		operation = "+"
	case strings.Contains(data, "/"):
		operation = "/"
	case strings.Contains(data, "*"):
		operation = "*"
	}
	return
}

func convStrTONum(data []string) (nums []int, err error) {
	for _, numStr := range data {
		num, _ := strconv.Atoi(numStr)
		if num > 10 || num < 1 {
			err = fmt.Errorf("%w", myErrors.NumError)
			break
		}
		nums = append(nums, num)
	}
	return
}

func romanToArabic(data []string) (nums []int, err error) {
	for _, romanNum := range data {
		if num, ok := models.RomanArabic[romanNum]; ok {
			nums = append(nums, num)
		} else {
			err = fmt.Errorf("%w", myErrors.NumError)
		}
	}
	return
}
