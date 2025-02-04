// service resource is yet to be implemented

package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type RunbookGenerator struct {
	SCService
}

type Runbook struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *RunbookGenerator) createResources(runbooks []Runbook) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, runbook := range runbooks {
		resourceList = append(resourceList, terraformutils.NewResource(
			runbook.ID,
			fmt.Sprintf("runbook_%s", runbook.Name),
			"squadcast_runbook",
			g.GetProviderName(),
			map[string]string{
				"team_id": g.Args["team_id"].(string),
			},
			[]string{},
			map[string]interface{}{},
		))
	}
	return resourceList
}

func (g *RunbookGenerator) InitResources() error {
	getRunbooksURL := "/v3/runbooks"
	response, err := Request[[]Runbook](getRunbooksURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
