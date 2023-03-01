package models

type JSONDeliveryTicket struct { // Information sent as a post request from the ocr stage
	ExtractedFields []string             `json:"extractedFields"`
	TotalWeight     JSONValueInformation `json:"total weight"` // Total weight of the delivery
	ImageUrl        string               `json:"image_url"`
}

type JSONValueInformation struct {
	Success          bool                  `json:"success"`
	Value            string                `json:"value"`
	ErrorInformation JSONErrorInformation  `json:"error"`
	Checks           JSONChecksInformation `json:"checks"`
}

type JSONChecksInformation struct {
	ExtremeValueCheck bool `json:"extremeValueCheck"`
	DecimalPlaceCheck bool `json:"decimalPlaceCheck"`
}

type JSONErrorInformation struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}
