package es

import (
	"fmt"
)

const (
	House = "house"
)

func (c *Client) Index(m map[string]interface{}) error {
	_, err := c.Client.Index().Index(House).
		Routing(fmt.Sprintf("%d", m["city_id"])).
		Id(m["id"].(string)).
		BodyJson(m).Do(c.ctx)

	return err
}
