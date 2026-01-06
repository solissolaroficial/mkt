package response

// CalendarPostResponse representa a resposta de um post do calendário
type CalendarPostResponse struct {
	ID                 string              `json:"id"`
	Title              string              `json:"title"`
	Description        *string             `json:"description,omitempty"`
	Date               string              `json:"date"`
	Time               string              `json:"time"`
	Caption            *string             `json:"caption,omitempty"`
	Category           string              `json:"category"`
	Type               string              `json:"type"`
	Status             string              `json:"status"`
	AssigneeID         *string             `json:"assignee_id,omitempty"`
	Assignee           *PublicUserResponse `json:"assignee,omitempty"`
	Platforms          []string            `json:"platforms"`
	PublishedPlatforms []string            `json:"published_platforms"`
	ImageURL           *string             `json:"image_url,omitempty"`
	History            []PostHistoryEvent  `json:"history,omitempty"`
	CreatedAt          string              `json:"created_at"`
	UpdatedAt          string              `json:"updated_at"`
}

// PostHistoryEvent representa um evento no histórico do post
type PostHistoryEvent struct {
	ID        string  `json:"id"`
	Action    string  `json:"action"`
	User      string  `json:"user"`
	Text      *string `json:"text,omitempty"`
	Timestamp string  `json:"timestamp"`
}

// CalendarPostsListResponse representa a resposta de lista de posts do calendário
type CalendarPostsListResponse struct {
	Data       []CalendarPostResponse `json:"data"`
	Total      int64                  `json:"total"`
	Page       int                    `json:"page"`
	Limit      int                    `json:"limit"`
	TotalPages int                    `json:"total_pages"`
}
