package functionaltesting

import (
	"encoding/json"
	"io"

	"github.com/oseducation/knowledge-graph/model"
	"github.com/pkg/errors"
)

func (c *Client) usersRoute() string {
	return "/users"
}

func (c *Client) RegisterUser(user *model.User) (*model.User, *Response, error) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can't marshal user")
	}

	r, err := c.DoAPIPost(c.usersRoute()+"/register", string(userJSON))
	if err != nil {
		return nil, BuildResponse(r), err
	}

	u, err := decodeUser(r.Body)
	if err != nil {
		return nil, BuildResponse(r), err
	}
	return u, BuildResponse(r), nil
}

// CreateUser creates a user in the system based on the provided user struct.
func (c *Client) CreateUser(user *model.User) (*model.User, *Response, error) {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can't marshal user")
	}

	r, err := c.DoAPIPost(c.usersRoute(), string(userJSON))
	if err != nil {
		return nil, BuildResponse(r), err
	}
	defer closeBody(r)
	u, err := decodeUser(r.Body)
	if err != nil {
		return nil, BuildResponse(r), err
	}
	return u, BuildResponse(r), nil
}

// LoginByEmail authenticates a user by user email and password.
func (c *Client) LoginByEmail(email string, password string) (*model.User, *Response, error) {
	m := make(map[string]string)
	m["email"] = email
	m["password"] = password
	return c.login(m)
}

func (c *Client) login(m map[string]string) (*model.User, *Response, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, nil, errors.Wrap(err, "can't marshal map")
	}

	r, err := c.DoAPIPost(c.usersRoute()+"/login", string(b))
	if err != nil {
		return nil, BuildResponse(r), err
	}
	defer closeBody(r)
	c.AuthToken = r.Header.Get(HeaderToken)
	c.AuthType = HeaderBearer

	u, err := decodeUser(r.Body)
	if err != nil {
		return nil, BuildResponse(r), err
	}
	return u, BuildResponse(r), nil
}

func decodeUser(reader io.ReadCloser) (*model.User, error) {
	var user model.User
	if err := json.NewDecoder(reader).Decode(&user); err != nil {
		return nil, errors.Wrap(err, "can't decode user")
	}
	return &user, nil
}