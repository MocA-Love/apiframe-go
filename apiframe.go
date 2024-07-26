package apiframe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ApiframeClient struct {
	baseURL   string
	apiKey    string
	verbose   bool
}

func NewApiframeClient(apiKey string, verbose bool) (*ApiframeClient, error) {
	if apiKey == "" {
		return nil, errors.New("The api_key is required!")
	}
	return &ApiframeClient{
		baseURL: "https://api.apiframe.pro",
		apiKey:  apiKey,
		verbose: verbose,
	}, nil
}

func (c *ApiframeClient) postRequest(endpoint string, data interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, err
	}

	if c.verbose {
		fmt.Printf("Response: %v\n", result)
	}

	return result, nil
}

func (c *ApiframeClient) Imagine(prompt string, aspectRatio string, processMode string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"prompt":        prompt,
		"aspect_ratio":  aspectRatio,
		"process_mode":  processMode,
		"webhook_url":   webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/imagine", data)
}

func (c *ApiframeClient) Upscale1x(parentTaskID string, index string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"index":          index,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/upscale-1x", data)
}

func (c *ApiframeClient) UpscaleAlt(parentTaskID string, upscaleType string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"type":           upscaleType,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/upscale-alt", data)
}

func (c *ApiframeClient) UpscaleHighres(parentTaskID string, upscaleType string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"type":           upscaleType,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/upscale-highres", data)
}

func (c *ApiframeClient) Reroll(parentTaskID string, prompt string, aspectRatio string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"prompt":         prompt,
		"aspect_ratio":   aspectRatio,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/reroll", data)
}

func (c *ApiframeClient) Variations(parentTaskID string, index string, prompt string, aspectRatio string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"index":          index,
		"prompt":         prompt,
		"aspect_ratio":   aspectRatio,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/variations", data)
}

func (c *ApiframeClient) Inpaint(parentTaskID string, mask string, prompt string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"mask":           mask,
		"prompt":         prompt,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/inpaint", data)
}

func (c *ApiframeClient) Outpaint(parentTaskID string, zoomRatio string, aspectRatio string, prompt string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"zoom_ratio":     zoomRatio,
		"aspect_ratio":   aspectRatio,
		"prompt":         prompt,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/outpaint", data)
}

func (c *ApiframeClient) Pan(parentTaskID string, direction string, prompt string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"parent_task_id": parentTaskID,
		"direction":      direction,
		"prompt":         prompt,
		"webhook_url":    webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/pan", data)
}

func (c *ApiframeClient) Describe(imageURL string, processMode string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"image_url":     imageURL,
		"process_mode":  processMode,
		"webhook_url":   webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/describe", data)
}

func (c *ApiframeClient) Blend(imageURLs []string, dimension string, processMode string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"image_urls":    imageURLs,
		"dimension":     dimension,
		"process_mode":  processMode,
		"webhook_url":   webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/blend", data)
}

func (c *ApiframeClient) Seed(taskID string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"task_id":       taskID,
		"webhook_url":   webhookURL,
		"webhook_secret": webhookSecret,
	}
	return c.postRequest("/seed", data)
}

func (c *ApiframeClient) Faceswap(targetImageURL string, swapImageURL string, webhookURL string, webhookSecret string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"target_image_url": targetImageURL,
		"swap_image_url":   swapImageURL,
		"webhook_url":      webhookURL,
		"webhook_secret":   webhookSecret,
	}
	return c.postRequest("/faceswap", data)
}

func (c *ApiframeClient) Fetch(taskID string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"task_id": taskID,
	}
	return c.postRequest("/fetch", data)
}

func (c *ApiframeClient) FetchMany(taskIDs []string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"task_ids": taskIDs,
	}
	return c.postRequest("/fetch-many", data)
}

func (c *ApiframeClient) Account() (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/account", c.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, err
	}

	if c.verbose {
		fmt.Printf("Response: %v\n", result)
	}

	return result, nil
}
