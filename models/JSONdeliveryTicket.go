package models

type JSONDeliveryTicket struct { // Information sent as a post request from the ocr stage
	Success          bool                 `json:"success"`
	Weight           JSONValueInformation `json:"weight"` // Total weight of the delivery
	ImageUrl         string               `json:"image_url"`
	ErrorInformation JSONErrorInformation `json:"error"`
}

type JSONValueInformation struct {
	Value  float32               `json:"value"`
	Checks JSONChecksInformation `json:"checks"`
}

type JSONChecksInformation struct {
	ExtremeValueCheck bool `json:"extremeValueCheck"`
	DecimalPlaceCheck bool `json:"decimalPlaceCheck"`
}

type JSONErrorInformation struct {
	Type        string                     `json:"type"` //field name
	Description string                     `json:"description"`
	Likelihoods JSONLikelihoodsInformation `json:"likelihoods"`
}

type JSONLikelihoodsInformation struct {
	Threshold float32 `json:"threshold"`
	Actual    float32 `json:"actual"`
}
