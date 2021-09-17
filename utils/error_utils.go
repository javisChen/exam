package utils

func TryThrowError(e error) {
	if e != nil {
		panic(e)
	}
}
