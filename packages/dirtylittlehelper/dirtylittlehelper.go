// Package dirtylittlehelper provides general helper functions / utils
package dirtylittlehelper

// helper function to remove element of a slice, does not maintain sort order
func RemoveFromSlice(fullSlice []int, indexToRemove int) []int {
	fullSlice[indexToRemove] = fullSlice[len(fullSlice)-1]
	return fullSlice[:len(fullSlice)-1]
}
