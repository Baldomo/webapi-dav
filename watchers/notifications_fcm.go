package watchers

import (
	"fmt"
	"github.com/Baldomo/webapi-dav/comunicati"
	. "github.com/Baldomo/webapi-dav/log"
	"github.com/appleboy/go-fcm"
	"net/url"
)

const (
	apikey   = "AAAA_mxAYMw:APA91bFGCkamxX8_BkjIQP4mWdSc7ZI_3QarBGYW205oGFVyYTuFZipU8WKzO1tDfypspEQrsAcnbW4Wk4720lBhKRrBhCPELFp9YH8Wfujo4KW_QLUwG3pO2E44M6_emEEllthkHPhU"
	duration = 86400
)

var (
	client *fcm.Client = nil
)

func NotifyComunicato(filename string, tipo string) {
	message := &fcm.Message{
		To: fmt.Sprintf("/topics/comunicati-%s", tipo),
		Notification: &fcm.Notification{
			Title: fmt.Sprintf("Comunicati %s", tipo),
			Body:  fmt.Sprintf("Nuovo comunicato: %s", filename),
		},
		Data: map[string]interface{}{
			"": fmt.Sprintf(comunicati.UrlPrefix + "comunicati-" + tipo + "/" + url.PathEscape(filename)),
		},
		TimeToLive: func(i uint) *uint { return &i }(duration),
	}

	if client == nil {
		client, _ = fcm.NewClient(apikey)
	}
	resp, err := client.Send(message)
	if err != nil {
		Log.Error(err.Error())
	} else {
		Log.Info("Notifica inviata con successo!")
		Log.Infof("%#v", message)
		Log.Infof("%#v", resp)
	}
}
