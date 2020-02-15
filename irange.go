// Provides a way to create sequence of integers. Simple, but helpful.
package irange

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func getParams(params ...int) (start int, end int, step int) {
	numParams := 0

	for i, param := range params {
		numParams++
		switch i {
		case 0:
			start = param
		case 1:
			end = param
		case 2:
			step = param
		}
	}

	switch numParams {
	case 0:
		return 0, 0, 1
	case 1:
		return 0, start, 1
	case 2:
		return start, end, 1
	default:
		return start, end, step
	}
}

// Produce a sequence of numbers.
// Accepts at least 1, and up to 3 parameters.
// If a single parameter is given, 'end', returns the numbers from 0 to end, exclusive
// Range(3) returns [0, 1, 2].
// If more than 1 parameter is given "Range(start, end, step?)", Range returns the numbers from
// 'start' (inclusive) to 'end' (exclusive), with increments of 'step'.
func Range(params ...int) []int {
	start, end, step := getParams(params...)

	var len int
	if step > 0 && start < end {
		len = (end-1-start)/step + 1
	} else if step < 0 && start > end {
		len = 1 + (start-1-end)/-step
	} else {
		len = 0
	}

	if len == 0 {
		return []int{}
	}

	len = max(len, 0)
	res := make([]int, len)

	for i := 0; i < len; i++ {
		res[i] = start + step*i
	}

	return res
}
