package watchers

import (
	"fmt"
	"net/url"
	"os"

	"github.com/Baldomo/webapi-dav/pkg/comunicati"
	"github.com/Baldomo/webapi-dav/pkg/config"
	"github.com/Baldomo/webapi-dav/pkg/log"
	"github.com/appleboy/go-fcm"
)

const (
	duration = 86400
)

var (
	apikey             = os.Getenv("WEBAPI_FCM_KEY")
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
			"": fmt.Sprint("https://"+config.GetConfig().General.FQDN+comunicati.PathPrefix, "comunicati-", tipo, "/", url.PathEscape(filename)),
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
