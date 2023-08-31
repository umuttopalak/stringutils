/*
Package stringutils implements basic string utility functions for demo
purposes only!
*/
package stringutils

// Reverse reverses given string!
func Reverse(s string) string {
	r := []rune(s)
	lr := len(r)
	ss := make([]rune, lr)

	for i := 0; i < lr; i++ {
		ss[lr-1-i] = r[i]
	}

	return string(ss)
}
