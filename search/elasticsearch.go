package search

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"

	"github.com/wsgoggway/library/models"
	"github.com/wsgoggway/library/utils"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/kelseyhightower/envconfig"
)

var (
	c = new(config)
)

// Config read only.
type config struct {
	Addresses []string `envconfig:"SEARCH_ENGINE_ADDRESSES" required:"true"`
	Username  string   `envconfig:"SEARCH_ENGINE_USERNAME" required:"true"`
	Possword  string   `envconfig:"SEARCH_ENGINE_PASSWORD" required:"true"`
}

// Init initialization from environment variables
func init() {
	if err := envconfig.Process("", c); err != nil {
		log.Fatalf("search layer failed to load configuration: %s", err)
	}
}

// Elasticsearch is SearchEngine implementation
type Elasticsearch struct {
	log       utils.Logger
	esClient  *elasticsearch.Client
	fuzziness int64
}

var _ SearchEngine = (*Elasticsearch)(nil) // validate interface

func New(log utils.Logger, fuzziness int64) (es *Elasticsearch, err error) {
	es = &Elasticsearch{log: log, fuzziness: fuzziness}

	log.Info("Connect to elasticsearch...")
	es.esClient, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.Addresses,
		Username:  c.Username,
		Password:  c.Possword,
		Transport: &Transport{},
	})
	if err != nil {
		return nil, err
	}

	log.Info("Connected...")

	return es, nil
}

func (es *Elasticsearch) Search(ctx context.Context, searchString string, limit, offset int, index string, fields []string) (offersIds []int64, err error) {
	query := models.Query{
		Query: &models.QueryClass{
			MultiMatch: &models.MultiMatch{
				Query:     searchString,
				Fuzziness: es.fuzziness,
				Fields:    fields,
			},
		},
		Highlight: &models.Highlight{
			Fields: &models.Fields{
				Title: &models.Title{},
			},
		},
	}

	q, err := query.Marshal()
	if err != nil {
		return nil, err
	}

	read := bytes.NewReader(q)

	res, err := es.esClient.Search(
		es.esClient.Search.WithBody(read),
		es.esClient.Search.WithIndex(index),
		es.esClient.Search.WithSize(limit),
		es.esClient.Search.WithFrom(offset),
		es.esClient.Search.WithTrackTotalHits(true),
		es.esClient.Search.WithPretty(),
		es.esClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	lReponseModel, err := models.UnmarshalResponse(body)
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(lReponseModel.Hits.Hits)) // offer weights
	for _, item := range lReponseModel.Hits.Hits {
		ids = append(ids, item.Source.ID)
	}

	return ids, nil
}
