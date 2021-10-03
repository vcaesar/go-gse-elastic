# go-gse-elastic

## Put mapping
```json
"analyzer": {
    "gse_analyzer": {
        "type": "custom",
        "tokenizer": "whitespace",
        "filter": ["es_synonym"]
    }
}
```

```json
"title_index": {
    "type": "text",
    "analyzer": "gse_analyzer"
}
```

## Index

Create a new [client](/es/client.go) by [elastic](github.com/olivere/elastic) and config your elasticsearch servers.

```go
func (c *Client) data(m map[string]interface{}) map[string]interface{} {
	m["title_index"] = c.seg.CutHtml(m["title"].(string))
	m["describe_index"] = c.seg.CutHtml(m["describe"].(string))
	m["addr_index"] = c.seg.CutHtml(m["addr"].(string))
	return m
}

func (c *Client) Index(m map[string]interface{}) error {
	_, err := c.Client.Index().Index(House).
		Routing(fmt.Sprintf("%d", m["city_id"])).
		Id(m["id"].(string)).
		BodyJson(m).Do(c.ctx)

	return err
}
```

## Search

```go
// github.com/olivere/elastic/v7

func (c *Client) query(title string, cityId, districtId int64,
	page, count int) (int64, []map[string]interface{}, error) {
    qry := elastic.NewBoolQuery().Filter(elastic.NewTermQuery("city_id", cityId))
    qry.Must(
            elastic.NewTermQuery("title_index", c.seg.CutHtml(title)),
        )
    
    res, err := c.Client.Search(House).
		Query(qry).
		IgnoreUnavailable(true).
		From(page * count).Size(count).
		TrackTotalHits(true).
		Do(c.ctx)
	if err != nil {
		return 0, nil, err
	}

	results := make([]map[string]interface{}, 0)
	for _, r := range res.Hits.Hits {
		var data map[string]interface{}
		if err := json.Unmarshal(r.Source, &data); err != nil {
			return 0, nil, err
		}

		results = append(results, data)
	}

	return res.TotalHits(), results, nil
}
```