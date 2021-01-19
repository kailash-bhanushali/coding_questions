package rotationArray

func CalculateGCD(len, d int) int {
	val := 1
	temp := len
	if len > 1 && d > 1 {
		if len > d {
			temp = d
		}
		for i := temp; i > 1; i-- {
			if len%i == 0 && d%i == 0 {
				val = i
				break
			}
		}
	}
	return val
}
func Rotation_Array(arr []int, d int) []int {
	modFind := CalculateGCD(len(arr), d)
	if modFind > 1 {
		for i := 0; i < modFind; i++ {
			temp := arr[i]
			j := i
			counter := 0
			for counter < modFind {
				arr[j] = arr[j+modFind]
				j += modFind
				counter++
			}
			arr[j] = temp
		}
	} else {
		newarr := arr[0:d]
		arr = arr[d:len(arr)]
		arr = append(arr, newarr...)
	}
	return arr
}
