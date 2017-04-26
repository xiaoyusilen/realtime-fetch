// author by @xiaoyusilen

package route

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/xiaoyusilen/realtime-fetch/model"
)

func SendToAPI(position model.Test) bool {

	values := &model.Test{
		ID:        position.ID,
		Name:      position.Name,
		UpdatedAt: position.UpdatedAt,
	}

	b, _ := json.Marshal(&values)

	body := bytes.NewBuffer([]byte(b))

	url := "http://host:port/api"

	resp, err := http.Post(url, "application/json;charset=utf-8", body)

	if err != nil {
		log.Printf("err is: %s", err)
	}

	if resp.StatusCode == 200 {
		return true
	}
	return false
}
