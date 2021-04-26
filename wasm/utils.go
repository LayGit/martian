package main

func ReverseString(str string) string {
	var bytes = []byte(str)
	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	return string(bytes)
}