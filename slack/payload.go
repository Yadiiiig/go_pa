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
				SelectedDate string `json:"selected_date"`
				Type         string `json:"type"`
			} `json:"before"`
		} `json:"dates"`
	} `json:"values"`
}

type addAgendaItemPayload struct {
	Values struct {
		Date struct {
			DueDateInput struct {
				SelectedDate string `json:"selected_date"`
				Type         string `json:"type"`
			} `json:"due_date_input"`
		} `json:"date"`
		Information struct {
			InformationInput struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"information_input"`
		} `json:"information"`
		Title struct {
			TitleInput struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"title_input"`
		} `json:"title"`
	} `json:"values"`
}
