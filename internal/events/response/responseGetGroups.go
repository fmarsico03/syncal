package response

import "time"

type ResponseGetGroups struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Creator     string    `json:"creator"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
	Description string    `json:"description"`
	MeetLink    string    `json:"meet_link"`
}
