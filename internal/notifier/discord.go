package notifier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Discord struct {
	URL    string
	Client *http.Client
}

func NewDiscord(url string) Discord {
	return Discord{URL: url, Client: &http.Client{Timeout: 10 * time.Second}}
}

func (d Discord) Notify(ctx context.Context, message string) error {
	payload, err := json.Marshal(struct {
		Content string `json:"content"`
	}{Content: message})
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, d.URL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := d.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook HTTP %d", resp.StatusCode)
	}
	return nil
}
