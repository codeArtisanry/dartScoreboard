package types

import "time"

// User user
// swagger:model user
type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	AvatarURL string     `json:"avatarUrl,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// swagger:response UsersPaginationResponse
type UsersPaginationResponse struct {
	// in: body
	UserResponses []User `json:"user_responses"`
	PrePageLink   string `json:"pre_page_link,omitempty"`
	PostPageLink  string `json:"post_page_link,omitempty"`
}
