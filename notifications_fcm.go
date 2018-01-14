package main

import (
	"github.com/appleboy/go-fcm"
	"fmt"
)

var (
	client *fcm.Client = nil
)

func NotifyComunicato(filename string, tipo string) {
	message := &fcm.Message{
		To: fmt.Sprintf("comunicati-%s", tipo),
		Notification: &fcm.Notification{
			Title: fmt.Sprintf("Comunicati %s", tipo),
			Body: fmt.Sprintf("Nuovo comunicato: %s", filename),
		},
	}

	if client == nil {
		client, _ = fcm.NewClient("AAAA_mxAYMw:APA91bFGCkamxX8_BkjIQP4mWdSc7ZI_3QarBGYW205oGFVyYTuFZipU8WKzO1tDfypspEQrsAcnbW4Wk4720lBhKRrBhCPELFp9YH8Wfujo4KW_QLUwG3pO2E44M6_emEEllthkHPhU")
	}
	_, err := client.SendWithRetry(message, 3)
	if err != nil {
		Log.Error(err.Error())
	}
}

func formatFilename(filename string) string {
	return ""
}