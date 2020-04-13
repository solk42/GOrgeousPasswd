package dirtylittlehelper

// helper function to remove element of a slice
func RemoveFromSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
