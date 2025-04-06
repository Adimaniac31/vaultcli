// cmd/read.go
package cmd

import (
	"fmt"
	"io"
	"net/http"
)

func Read(key, token string) {
	url := fmt.Sprintf("http://localhost:8080/secret?key=%s", key)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("❌ Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("❌ Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
