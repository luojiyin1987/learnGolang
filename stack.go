package main

type Stack []byte

func (s *Stack) IsEmpy() bool {
	return len(*s) == 0
}

func (s *Stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpy() {
		return 0, false
	}

	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *Stack) Length() int {
	return len(*s)
}
