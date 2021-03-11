package main

var getAgendaBetweenModal = `{
	"type": "modal",
	"callback_id": "get_agenda_item_between",
	"title": {
		"type": "plain_text",
		"text": "My App",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Submit",
		"emoji": true
	},
	"close": {
		"type": "plain_text",
		"text": "Cancel",
		"emoji": true
	},
	"blocks": [
		{
			"type": "section",
			"text": {
				"type": "plain_text",
				"text": "Get agenda items between dates:",
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
		}
	]
}`

var addAgendaItemModal = `{
	"type": "modal",
	"callback_id": "add_agenda_item",
	"title": {
		"type": "plain_text",
		"text": "My App",
		"emoji": true
	},
	"submit": {
		"type": "plain_text",
		"text": "Submit",
		"emoji": true
	},
	"close": {
		"type": "plain_text",
		"text": "Cancel",
		"emoji": true
	},
	"blocks": [
		{
			"type": "header",
			"text": {
				"type": "plain_text",
				"text": "Create agenda item",
				"emoji": true
			}
		},
		{
			"type": "divider"
		},
		{
			"type": "input",
			"block_id": "title",
			"label": {
				"type": "plain_text",
				"text": "Title"
			},
			"element": {
				"type": "plain_text_input",
				"action_id": "title_input",
				"placeholder": {
					"type": "plain_text",
					"text": "Enter a title"
				}
			}
		},
		{
			"type": "input",
			"block_id": "information",
			"label": {
				"type": "plain_text",
				"text": "Information"
			},
			"element": {
				"type": "plain_text_input",
				"action_id": "information_input",
				"placeholder": {
					"type": "plain_text",
					"text": "Enter some information"
				}
			}
		},
		{
			"type": "section",
			"text": {
				"type": "mrkdwn",
				"text": "Pick a due date."
			},
			"accessory": {
				"type": "datepicker",
				"placeholder": {
					"type": "plain_text",
					"text": "Select a date",
					"emoji": true
				},
				"action_id": "due_date_input"
			}
		}
	]
}`
