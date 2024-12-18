package internal

// Rule represents a page ordering rule where Before must be printed before After
type Rule struct {
	Before int
	After  int
}
