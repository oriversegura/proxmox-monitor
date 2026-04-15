package proxmox

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oriversegura/proxmox-monitor/internal/models"
)

type vmsResponse struct {
	Data []vmsModel `json:"data"`
}

type vmsModel struct {
	Name      string `json:"name"`
	Status    string `json:"status"`
	IP        string `json:"ip"`
	Timestamp int64  `json:"uptime"`
	ID        string `json:"vm_id"`
	Kind      string `json:"type"`
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

	var result vmsResponse
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	for _, vmsModel := range result.Data {

		var status models.Status

		if vmsModel.Status == "running" {
			status = models.StatusOnline
		} else {
			status = models.StatusOffline
		}

		vms := models.Vm{
			Base: models.Base{
				Name:   vmsModel.Name,
				IP:     vmsModel.IP,
				Status: status,
			},
			ID:   vmsModel.ID,
			Kind: models.Kind(vmsModel.Kind),
		}
		nodes = append(nodes, vms)
	}

	return nodes, nil
}
