package utils

func Crash(err error) {
	if err != nil {
		panic(err)
	}
}
