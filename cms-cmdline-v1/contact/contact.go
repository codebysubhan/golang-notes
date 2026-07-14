package contact


func CreateContact (d *[][3]string, n string, ph string, e string) bool {
	*d = append(*d, [3]string{n,ph,e})
	return true
}

func GetContact (d [][3]string, n string) (int, bool) {
	for i := 0 ; i < len(d) ; i++ {
		if d[i][0] == n {
			return i, true
		}
	} // Go will automatically fill it with empty strings (the zero values) for you!
	return 0, false // return [3]string{}, false
}

func DeleteContact (d *[][3]string, n string) bool {
	idx , ok := GetContact(*d,n)
	if ok == false {
		return false
	}
	copy_d := make([][3]string, len(*d))	
	copy(copy_d, *d)

	// *d = copy_d[:idx]
	// *d = append(*d, copy_d[idx+1:]) // unpacking issue

	// gpt corrected output
	// 2. Perform the deletion trick safely entirely inside copy_d
    // We assign the result back to copy_d, unpacking the second half with ...
    copy_d = append(copy_d[:idx], copy_d[idx+1:]...)
    // 3. Overwrite the original pointer with our freshly modified copy
    *d = copy_d
	return true
}