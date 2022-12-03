package models

var (
	RomanArabic = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	ArabicRoman = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C"}
)

var (
	Arabic = "arabic"
	Roman  = "roman"
)

type MathData struct {
	TypeNum string
	Data    string
}

type DataForCalc struct {
	Nums      []int
	Operation string
	TypeNum   string
}

type Answer struct {
	Num     int
	TypeNum string
}
