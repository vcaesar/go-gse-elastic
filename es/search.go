package es

import (
	"encoding/json"

	"github.com/olivere/elastic/v7"
)

const (
	SearchSize = 100
)

func (c *Client) FilterId(cityId, districtId int64, defId ...int) *elastic.BoolQuery {
	qry := elastic.NewBoolQuery().Filter(elastic.NewTermQuery("city_id", cityId))
	if districtId <= 0 {
		return qry
	}

	return qry.Filter(elastic.NewTermQuery("district_id", districtId))
}

func (c *Client) Query(title string, cityId, districtId int64, page, count int) (
	int64, []map[string]interface{}, error) {
	qry := c.FilterId(cityId, districtId)
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
