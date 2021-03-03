package main

type agendaBetweenPayload struct {
	Values struct {
		Dates struct {
			After struct {
				Type         string `json:"type"`
				SelectedDate string `json:"selected_date"`
			} `json:"after"`
			Before struct {
				Type         string `json:"type"`
				SelectedDate string `json:"selected_date"`
			} `json:"before"`
		} `json:"dates"`
	} `json:"values"`
}

type actionsStruct struct {
	Type     string `json:"type"`
	BlockID  string `json:"block_id"`
	ActionID string `json:"action_id"`
	Text     struct {
		Type  string `json:"type"`
		Text  string `json:"text"`
		Emoji bool   `json:"emoji"`
	} `json:"text"`
	Value    string `json:"value"`
	ActionTs string `json:"action_ts"`
}
