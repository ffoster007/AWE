package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HealthChecker ตรวจสอบสถานะเว็บไซต์
type HealthChecker struct {
	URL           string
	CheckInterval time.Duration
	Timeout       time.Duration
	AlertFunc     func(status string, err error)
}

// Check ทำการตรวจสอบสถานะ
func (h *HealthChecker) Check() (bool, error) {
	client := &http.Client{
		Timeout: h.Timeout,
	}

	resp, err := client.Get(h.URL)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// ตรวจสอบ status code
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("status code: %d", resp.StatusCode)
}

// StartMonitoring เริ่มต้นการตรวจสอบแบบต่อเนื่อง
func (h *HealthChecker) StartMonitoring() {
	ticker := time.NewTicker(h.CheckInterval)
	defer ticker.Stop()

	log.Printf("เริ่มตรวจสอบ %s ทุกๆ %v\n", h.URL, h.CheckInterval)

	for range ticker.C {
		isOnline, err := h.Check()

		if isOnline {
			h.AlertFunc("ONLINE", nil)
		} else {
			h.AlertFunc("OFFLINE", err)
		}
	}
}

// ฟังก์ชันแจ้งเตือนผ่าน Console
func consoleAlert(status string, err error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if status == "ONLINE" {
		log.Printf("[%s] ✅ สถานะ: %s\n", timestamp, status)
	} else {
		log.Printf("[%s] ❌ สถานะ: %s - Error: %v\n", timestamp, status, err)
	}
}

// ฟังก์ชันแจ้งเตือนผ่าน LINE Notify (ต้องมี Token)
func lineNotifyAlert(token string) func(string, error) {
	return func(status string, err error) {
		message := fmt.Sprintf("สถานะเว็บไซต์: %s", status)
		if err != nil {
			message += fmt.Sprintf("\nError: %v", err)
		}

		// ส่ง POST ไปที่ LINE Notify API
		client := &http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest("POST",
			"https://notify-api.line.me/api/notify",
			nil)

		q := req.URL.Query()
		q.Add("message", message)
		req.URL.RawQuery = q.Encode()

		req.Header.Set("Authorization", "Bearer "+token)

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("ส่ง LINE Notify ไม่สำเร็จ: %v\n", err)
			return
		}
		defer resp.Body.Close()

		log.Printf("ส่ง LINE Notify สำเร็จ (Status: %d)\n", resp.StatusCode)
	}
}

func main() {
	// ตัวอย่างการใช้งาน
	checker := &HealthChecker{
		URL:           "https://example.com", // เปลี่ยนเป็น URL ของคุณ
		CheckInterval: 30 * time.Second,      // ตรวจสอบทุก 30 วินาที
		Timeout:       10 * time.Second,      // Timeout 10 วินาที
		AlertFunc:     consoleAlert,          // ใช้ console alert

		// หรือใช้ LINE Notify
		// AlertFunc: lineNotifyAlert("YOUR_LINE_TOKEN"),
	}

	// เริ่มการตรวจสอบ
	checker.StartMonitoring()
}
