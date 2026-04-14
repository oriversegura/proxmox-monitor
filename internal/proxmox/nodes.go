package proxmox

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oriversegura/proxmox-monitor/internal/models"
)

type nodesResponse struct {
	Data []nodesModel `json:"data"`
}

type nodesModel struct {
	Name      string `json:"node"`
	Status    string `json:"status"`
	IP        string `json:"ip"`
	Timestamp int64  `json:"uptime"`
}

func GetNodes(client *http.Client, url string, apiToken string) (nodes []models.Node, error error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "PVEAPIToken="+apiToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var result nodesResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	for _, nodesModel := range result.Data {

		var status models.Status

		if nodesModel.Status == "online" {
			status = models.StatusOnline
		} else {
			status = models.StatusOffline
		}

		node := models.Node{
			Hostname: nodesModel.Name,
			Base: models.Base{
				Name:   nodesModel.Name,
				IP:     nodesModel.IP,
				Status: status,
			},
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
