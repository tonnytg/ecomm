package client

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type ClientInterface interface {
	Enable() error
	Disable() error
	GetID() uuid.UUID
	GetName() string
	GetUsername() string
}

type Client struct {
	gorm.Model
	Address    string    `json:"address"`
	Email      string    `json:"email"`
	ValidEmail bool      `json:"valid_email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	Status     string    `json:"status"`
	ID         uuid.UUID `json:"user_id"`
}

func NewClient(firstname, lastname, username string) *Client {
	client := Client{
		ID:        uuid.New(),
		Status:    DISABLED,
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
	}
	return &client
}

func (c *Client) GetID() uuid.UUID {
	return c.ID
}

func (c *Client) GetUsername() string {
	return c.Username
}

func (c *Client) GetName() string {
	return c.Username
}

func (c *Client) Disable() error {
	if c.ValidEmail == true {
		c.Status = DISABLED
		c.ValidEmail = false
		return nil
	}
	return errors.New("email must be valid to disabled")
}

func (c *Client) Enable() error {
	if c.ValidEmail == false {
		c.Status = ENABLED
		c.ValidEmail = true
		return nil
	}
	return errors.New("email must be valid to enable")
}
