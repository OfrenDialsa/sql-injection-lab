package inject

import "fmt"

func runBasicTest(baseURL string) {
	fmt.Println("\n=== BASIC SQLi TEST ===")

	normal := send(baseURL, "1")
	injected := send(baseURL, "1 OR 1=1")

	fmt.Println("Normal:\n", normal)
	fmt.Println("Injected:\n", injected)

	if injected != normal {
		fmt.Println("⚠️ SQL Injection detected!")
	}
}
