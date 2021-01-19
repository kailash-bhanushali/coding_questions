package rotationArray

func CalculateGCD(len, d int) int {
	val := 1
	if len > 1 && d > 1 {
		for i := d; i > 1; i-- {
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
	for i := 0; i < modFind; i++ {
		temp := arr[i]
		j := i
		counter := 0
		maxInnerLoop := (len(arr) / d) - 1
		if modFind == 1 {
			maxInnerLoop = len(arr) - 1
		}
		for counter < maxInnerLoop {
			arr[j] = arr[j+modFind]
			j += modFind
			counter++
		}
		arr[j] = temp
	}
	return arr
}
