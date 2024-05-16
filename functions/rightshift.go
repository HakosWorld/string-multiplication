package reload

func Rightshift(arry []string, i int) []string {
	if i < len(arry)-1 {
		for j := i; j < len(arry)-1; j++ {
			arry[j] = arry[j+1]
		}
		arry = arry[:len(arry)-1]
		return (arry)
	}
	arry = arry[:len(arry)-1]
	return (arry)
}
