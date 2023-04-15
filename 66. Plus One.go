package main
func main() {
	digits := []int{1, 2, 3}
	plusOne(digits)
}
func plusOne(digits []int) []int {
	digits[len(digits) - 1] = digits[len(digits) - 1] + 1
	for i := 1; i < len(digits); i++ {
		if digits[len(digits) - i] == 10 {
			digits[len(digits) - i] = 0
			digits[len(digits) - i - 1] = digits[len(digits) - i - 1] + 1
		}
	}
	if digits[0] == 10 {
		digits[0] = 0
		digits2 := []int{1}
		for i := 0; i < len(digits); i++ {
		 digits2 = append(digits2, digits[i])
		}
		return digits2
	}
	return digits
 }