package numbers

// This is a Public function because we are using a uppercase letter
// Only names that starts with a uppercase letter can be accessed from outside the package
func Even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

// We will get an errror if we try to access it outside this package because it has a lowercase letter
// however, we can access it from within the Package
// So, odd will be a private function
func odd(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}

func isPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func OddAndPrime(n int) bool {
	if odd(n) && isPrime(n) {
		return true
	}
	return false
}
