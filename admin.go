package mastodon

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type AdminAccount struct {
	ID                     ID        `json:"id"`
	Username               string    `json:"username"`
	Domain                 string    `json:"domain"`
	CreatedAt              time.Time `json:"created_at"`
	Email                  string    `json:"email"`
	IP                     string    `json:"ip"`
	IPs                    []IP      `json:"ips"`
	Locale                 string    `json:"locale"`
	InviteRequest          string    `json:"invite_request"`
	Role                   *Role     `json:"role"`
	Confirmed              bool      `json:"confirmed"`
	Approved               bool      `json:"approved"`
	Disabled               bool      `json:"disabled"`
	Silenced               bool      `json:"silenced"`
	Suspended              bool      `json:"suspended"`
	Account                *Account  `json:"account"`
	CreatedByApplicationID ID        `json:"created_by_application_id"`
	InvitedByAccountID     ID        `json:"invited_by_account_id"`
}

type IP struct {
	IP     string    `json:"ip"`
	UsedAt time.Time `json:"used_at"`
}

type Role struct {
	ID          ID     `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Permissions string `json:"permissions"`
	Highlighted bool   `json:"highlighted"`
}

type AdminViewAccountsInput struct {
	Origin      string   `json:"origin,omitempty"`
	Status      string   `json:"status,omitempty"`
	Permissions string   `json:"permissions,omitempty"`
	RoleIDs     []string `json:"role_ids,omitempty"`
	InvitedBy   ID       `json:"invited_by,omitempty"`
	Username    string   `json:"username,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	ByDomain    string   `json:"by_domain,omitempty"`
	Email       string   `json:"email,omitempty"`
	IP          string   `json:"ip,omitempty"`
	MaxID       ID       `json:"max_id,omitempty"`
	SinceID     ID       `json:"since_id,omitempty"`
	MinID       ID       `json:"min_id,omitempty"`
	Limit       int64    `json:"limit"`
}

func (c *Client) AdminViewAccounts(ctx context.Context, input *AdminViewAccountsInput, pg *Pagination) ([]*AdminAccount, error) {
	inputBytes, _ := json.Marshal(input)
	params := url.Values{}
	json.Unmarshal(inputBytes, &params)

	var adminAccounts []*AdminAccount
	err := c.doAPI(ctx, http.MethodGet, "/api/v2/admin/accounts", params, &adminAccounts, pg)
	if err != nil {
		return nil, err
	}
	return adminAccounts, nil
}
