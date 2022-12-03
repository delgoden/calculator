package validation

import (
	"fmt"
	"github.com/delgoden/calculator/internal/models"
	"github.com/delgoden/calculator/internal/myErrors"
	"regexp"
	"strings"
)

func Valid(data string) (mathData models.MathData, err error) {
	arabicNum := regexp.MustCompile(`(^\d+ ?[-+/*] ?\d+$)`)
	romanNum := regexp.MustCompile(`^[IVXLCDM]+ ?[-+/*] ?[IVXLCDM]+$`)
	differentCalculusSystems, _ := regexp.Compile(`(^\d+ ?[-+/*] ?[IVXLCDM]+$)|(^[IVXLCDM]+ ?[-+/*] ?\d+$)`)
	charCount, _ := regexp.Compile(`[a-zA-Z]`)
	mathOperation, _ := regexp.Compile(`[-+/*]`)
	if arabicNum.MatchString(data) {
		mathData.TypeNum = models.Arabic
		mathData.Data = data
		return
	}
	if romanNum.MatchString(data) {
		mathData.TypeNum = models.Roman
		mathData.Data = data
		return
	}
	if differentCalculusSystems.MatchString(data) {
		err = fmt.Errorf("%w", myErrors.DifferentNumberSystemsError)
		return
	}
	if mathOpSls := mathOperation.FindAll([]byte(data), -1); len(mathOpSls) > 1 {
		err = fmt.Errorf("%w", myErrors.NotSatisfyTaskError)
		return
	}
	if !strings.ContainsAny(data, "-+/*") || charCount.MatchString(data) {
		err = fmt.Errorf("%w", myErrors.NotMathOperationError)
		return
	}
	return
}
