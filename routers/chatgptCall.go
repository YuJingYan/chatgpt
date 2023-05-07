package routers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const (
	OpenAIURL     = "https://api.openai.com/v1/chat/completions"
	APIKey        = "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"
	ServerAddress = ":8080"
)

type Content struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Prompt struct {
	Messages []Content `json:"messages"`
	Model    string    `json:"model"`
}

func chatHandler(c *gin.Context) {
	var prompt Prompt
	if err := c.ShouldBindJSON(&prompt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	answer, err := callChatGPT(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error calling ChatGPT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": answer})
}

func callChatGPT(prompt Prompt) (string, error) {
	data, err := json.Marshal(prompt)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", OpenAIURL, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", APIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	choices := response["choices"].([]interface{})
	answer := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	con := answer["content"].(string)

	return con, nil
}
