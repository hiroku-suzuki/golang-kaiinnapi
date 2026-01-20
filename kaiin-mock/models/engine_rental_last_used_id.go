package models

type EngineRentalLastUsedEngineRentalIdResponse struct {
	Model    string        `json:"model"`
	Errors   []ErrorItem   `json:"errors,omitempty"`
	Warnings []WarningItem `json:"warnings,omitempty"`
}
