package watchers

import (
	"fmt"
	"net/url"
	"os"

	"github.com/Baldomo/webapi-dav/pkg/comunicati"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/appleboy/go-fcm"
)

const (
	apikey   = os.Getenv("WEBAPI_FCM_KEY")
	duration = 86400
)

var (
	client *fcm.Client = nil
)

func NotifyComunicato(filename string, tipo string) {
	if apikey == "" {
		return
	}

	message := &fcm.Message{
		To: fmt.Sprintf("/topics/comunicati-%s", tipo),
		Notification: &fcm.Notification{
			Title: fmt.Sprintf("Comunicati %s", tipo),
			Body:  fmt.Sprintf("Nuovo comunicato: %s", filename),
		},
		Data: map[string]interface{}{
			"": fmt.Sprint(comunicati.UrlPrefix, "comunicati-", tipo, "/", url.PathEscape(filename)),
		},
		TimeToLive: func(i uint) *uint { return &i }(duration),
	}

	if client == nil {
		client, _ = fcm.NewClient(apikey)
	}
	resp, err := client.Send(message)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Notifica inviata con successo!")
		log.Infof("%#v", message)
		log.Infof("%#v", resp)
	}
}
