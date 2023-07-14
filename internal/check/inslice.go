package check

// InSlice - check if variable is in slice
func InSlice(id int64, slice []int64) bool {

	for i := range slice {
		if slice[i] == id {
			return true
		}
	}

	return false
}
