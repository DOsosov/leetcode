package main
func isPalindrome(x int) bool {
    var pruff bool = true
    conv := []byte(strconv.Itoa(x))
    for i := 0; i < len(conv)/2; i++ {
        conv[i], conv[len(conv) - 1 - i] = conv[len(conv) - 1 - i], conv[i]
    }
    convobr := []byte(strconv.Itoa(x))
    for i := 0; i < len(convobr); i++ {
        if convobr[i] != conv[i] {
            pruff = false
            return pruff
        }
    }
    return pruff
}