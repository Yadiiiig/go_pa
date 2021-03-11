package main

type modalPayload struct {
	CallbackID string                 `json:"callback_id"`
	State      map[string]interface{} `json:"state"`
}

type agendaBetweenPayload struct {
	Values struct {
		Dates struct {
			After struct {
				SelectedDate string `json:"selected_date"`
				Type         string `json:"type"`
			} `json:"after"`
			Before struct {
				SelectedDate interface{} `json:"selected_date"`
				Type         string      `json:"type"`
			} `json:"before"`
		} `json:"dates"`
	} `json:"values"`
}
