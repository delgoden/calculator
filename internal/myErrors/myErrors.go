package myErrors

import "errors"

var (
	DifferentNumberSystemsError = errors.New("it is forbidden to use different number systems at the same time")
	NotMathOperationError       = errors.New("a string is not a mathematical operation")
	NotSatisfyTaskError         = errors.New("the format of the mathematical operation does not satisfy the task — two operands and one operator (+, -, /, *)")
	NumError                    = errors.New("the numbers must be in the range from n 1(I) to 10(X)")
	RomanZeroNegativeError      = errors.New("there are no zero and negative numbers in the Roman system of calculus")
)
