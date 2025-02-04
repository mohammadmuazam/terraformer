package squadcast

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type EscalationPolicyGenerator struct {
	SCService
}

type EscalationPolicy struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (g *EscalationPolicyGenerator) createResources(policies []EscalationPolicy) []terraformutils.Resource {
	var resourceList []terraformutils.Resource
	for _, policy := range policies {
		resourceList = append(resourceList, terraformutils.NewResource(
			policy.ID,
			fmt.Sprintf("policy_%s", policy.Name),
			"squadcast_escalation_policy",
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

func (g *EscalationPolicyGenerator) InitResources() error {
	getEscalationPolicyURL := fmt.Sprintf("/v3/escalation-policies?owner_id=%s", g.Args["team_id"].(string))
	response, err := Request[[]EscalationPolicy](getEscalationPolicyURL, g.Args["access_token"].(string), g.Args["region"].(string), true)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(*response)
	return nil
}
