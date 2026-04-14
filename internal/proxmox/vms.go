package proxmox

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oriversegura/proxmox-monitor/internal/models"
)

type vmsResponse struct {
	Data []vmsResponse `json:"data"`
}

type vmsModel struct {
	Name      string `json:"node"`
	Status    string `json:"status"`
	IP        string `json:"ip"`
	Timestamp int64  `json:"uptime"`
}

func GetVMS(client *http.Client, url string, apiToken string) (nodes []models.Vm, error error) {

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

		vms := models.Vm{
			Hostname: vmsModel.Name,
			Base: models.Base{
				Name:   vmsModel.Name,
				IP:     vmsModel.IP,
				Status: status,
			},
		}
		nodes = append(nodes, vms)
	}

	return nodes, nil
}
