package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "lem/go/auth/proto"
)

const apiBaseURL = "http://localhost:8080/v1"

func main() {
	httpClinet := http.Client{
		Timeout: 10 * time.Second,
	}

	token, err := login(httpClinet, "admin", "admin")
	if err != nil {
		log.Printf("Login failed: %v\n", err)
	} else {
		log.Printf("Login successful, token: %s\n", token)
	}
}

func login(client http.Client, username, passowrd string) (string, error) {
	loginRequest := pb.LoginRequest{
		Username: username,
		Password: passowrd,
	}

	reqBody, err := json.Marshal(&loginRequest)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, apiBaseURL+"login", bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}

	var loginResponse pb.LoginResponse
	if err = json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return loginResponse.Token, nil
}
