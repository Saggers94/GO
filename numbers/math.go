// package mymath // this is not allowed because the package is a directory that contains this file,
// so it should be numbers
package numbers

func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
