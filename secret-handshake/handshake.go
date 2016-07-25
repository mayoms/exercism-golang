package secret

const REVERSE = 16

var handShakeMap = map[int]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func reverseArray(array []string) []string {

	for i := len(array)/2 - 1; i >= 0; i-- {
		opp := len(array) - 1 - i
		array[i], array[opp] = array[opp], array[i]
	}
	return array
}

func Handshake(code int) []string {

	if code < 1 {
		return nil
	}
	steps := []string{}
	for k := range handShakeMap {
		if code&k != 0 {
			steps = append(steps, handShakeMap[k])
		}
	}
	if code&REVERSE != 0 {
		steps = reverseArray(steps)
	}
	return steps
}
