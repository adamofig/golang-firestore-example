package entity

// Este request lo obtuve directamente de los datos que vienen de dialogflow, con una tecnica para imprimir el json

type Request struct {
	OriginalDetectIntentRequest struct {
		Payload struct {
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	QueryResult struct {
		AllRequiredParamsPresent bool `json:"allRequiredParamsPresent"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		} `json:"fulfillmentMessages"`
		Intent struct {
			DisplayName string `json:"displayName"`
			Name        string `json:"name"`
		} `json:"intent"`
		IntentDetectionConfidence float64 `json:"intentDetectionConfidence"`
		LanguageCode              string  `json:"languageCode"`
		OutputContexts            []struct {
			LifespanCount int    `json:"lifespanCount,omitempty"`
			Name          string `json:"name"`
			Parameters    struct {
				DateTime         string      `json:"date-time"`
				DateTimeOriginal string      `json:"date-time.original"`
				NoInput          interface{} `json:"no-input"`
				NoMatch          interface{} `json:"no-match"`
			} `json:"parameters,omitempty"`
		} `json:"outputContexts"`
		Parameters map[string]interface{} `json:"parameters"`
		QueryText  string                 `json:"queryText"`
	} `json:"queryResult"`
	ResponseID string `json:"responseId"`
	Session    string `json:"session"`
}
