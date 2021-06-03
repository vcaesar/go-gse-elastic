package es

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	"github.com/vcaesar/go-gse-elastic/nlp"
)

type Conf struct {
	URL     string `toml:"url" yaml:"url"`
	Sniff   *bool  `toml:"sniff" yaml:"sniff"`
	SetGzip *bool  `toml:"gzip" yaml:"gzip"`
}

type Client struct {
	Client *elastic.Client
	Conf   *Conf

	seg nlp.Segmenter
	ctx context.Context
}

func (c *Conf) Init() error {
	s := true
	if c.Sniff == nil {
		c.Sniff = &s
	}

	if c.SetGzip == nil {
		c.SetGzip = &s
	}

	return nil
}

func (c *Conf) New(ctx1 context.Context) (*Client, error) {
	err := c.Init()
	if err != nil {
		return nil, err
	}

	client, err := elastic.NewClient(
		elastic.SetURL(c.URL),
		elastic.SetSniff(*c.Sniff),
		elastic.SetGzip(*c.SetGzip),
	)

	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
		Conf:   c,
		ctx:    ctx1,
	}, nil
}

func Source(q elastic.Query) (string, error) {
	src, err := q.Source()

	data, _ := json.Marshal(src)
	return string(data), err
}

func Println(q elastic.Query) {
	fmt.Println()

	s, err := Source(q)
	fmt.Println(s, err)
}
