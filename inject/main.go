package inject

import "fmt"

func Run(baseURL string, mode string) {

	switch mode {
	case "basic":
		runBasicTest(baseURL)
	case "boolean":
		runBooleanTest(baseURL)
	case "time":
		runTimeTest(baseURL)
	default:
		fmt.Println("Unknown mode, fallback to basic")
		runBasicTest(baseURL)
	}
}
