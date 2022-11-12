package helper

import "fmt"

func ScanUint(msg string) uint {
	var result uint

	fmt.Print(msg)
	fmt.Scan(&result)

	return result
}
