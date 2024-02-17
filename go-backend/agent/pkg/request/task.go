package request

import (
	"distributed-calculator/orchestrator/pkg/models"
	"fmt"
	"io"
	"net/http"
)

func GetTask() *models.Task {
	resp, err := http.Get("http://localhost:8080/waiting-task/")
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	fmt.Println(string(body))
	return nil
}
