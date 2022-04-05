package models

type AlertSQS struct {
	Description     string `json:"description" validate:"required,max=140"`
	Severity        string `json:"severity" validate:"required"`
	Percentage      int    `json:"percentage" validate:"required"`
	Time            int    `json:"time" validate:"required"`
	Type            string `json:"type" validate:"required"`
	Threshold       int    `json:"threshold" validate:"required"`
	ApplicationUUID string `json:"application_uuid" validate:"required"`
}
