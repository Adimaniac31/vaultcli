// cmd/store.go
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Store(key, value, token string) {
	data := map[string]string{
		"key":   key,
		"value": value,
	}
	jsonData, _ := json.Marshal(data)

	req, err := http.NewRequest("POST", "http://localhost:8080/store", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("❌ Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("❌ Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
