package calculation

import "github.com/delgoden/calculator/internal/models"

func Calculation(data models.DataForCalc) (answer models.Answer) {
	switch data.Operation {
	case "-":
		answer.Num = data.Nums[0] - data.Nums[1]
		answer.TypeNum = data.TypeNum
	case "+":
		answer.Num = data.Nums[0] + data.Nums[1]
		answer.TypeNum = data.TypeNum
	case "/":
		answer.Num = data.Nums[0] / data.Nums[1]
		answer.TypeNum = data.TypeNum
	case "*":
		answer.TypeNum = data.TypeNum
		answer.Num = data.Nums[0] * data.Nums[1]
	}
	return
}
