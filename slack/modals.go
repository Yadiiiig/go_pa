package main

var getAgendaBetweenModal = `[
	{
		"type": "section",
		"text": {
			"type": "plain_text",
			"text": "Get agenda item between:",
			"emoji": true
		}
	},
	{
		"type": "actions",
		"block_id": "dates",
		"elements": [
			{
				"type": "datepicker",
				"placeholder": {
					"type": "plain_text",
					"text": "Select a date",
					"emoji": true
				},
				"action_id": "after"
			},
			{
				"type": "datepicker",
				"placeholder": {
					"type": "plain_text",
					"text": "Select a date",
					"emoji": true
				},
				"action_id": "before"
			}
		]
	},
	{
		"type": "actions",
		"elements": [
			{
				"type": "button",
				"text": {
					"type": "plain_text",
					"text": "Search",
					"emoji": true
				},
				"value": "submit",
				"action_id": "submit_agenda_between"
			}
		]
	}
]`
