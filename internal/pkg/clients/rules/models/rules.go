package models

type GetRuleDto struct {
	ID     string `json:"id"`
	TypeId string `json:"typeId"`
	Params string `json:"params"`
}

type CreateRuleDto struct {
	TypeId string `json:"typeId" binding:"required"`
	Params string `json:"params" binding:"required"`
}

type GetPolicyDto struct {
	ID        string   `json:"id"`
	ProjectID string   `json:"projectId"`
	RulesIDs  []string `json:"rulesIds"`
}

type CreatePolicyDto struct {
	ProjectID string   `json:"projectId" binding:"required"`
	RulesIDs  []string `json:"rulesIds" binding:"required"`
}
