func isValid(s string) bool {
    stack := linkedliststack.New()
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if !stack.Empty() {
                top, ok := stack.Pop()
                if ok && top.(rune) != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack.Push(c)
        }
    }

    return stack.Empty()
}