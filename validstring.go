package main

func validString(s string) bool {
	pairs := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	// make the stack
	stack := make([]byte, 0)

	for _, char := range []byte(s) {
		pair, ok := pairs[char]
		if !ok {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		if pair != stack[len(stack)-1] {
			return false
		}

		stack = stack[:len(stack)-1] // stripping the slice to emulate pop() function
	}
	return len(stack) == 0
}
