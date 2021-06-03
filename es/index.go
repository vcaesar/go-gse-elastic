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

func (c *Client) data(m map[string]interface{}) map[string]interface{} {
	m["title_index"] = c.seg.CutHtml(m["title"].(string))
	m["describe_index"] = c.seg.CutHtml(m["describe"].(string))
	m["addr_index"] = c.seg.CutHtml(m["addr"].(string))
	return m
}
