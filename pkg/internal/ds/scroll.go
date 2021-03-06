package ds

import (
	"fmt"
	"time"

	"github.com/RediSearch/redisearch-go/redisearch"
)

type redisScrollBuilder struct {
	client *redisearch.Client
	query  *redisearch.Query
}

func (sc *redisScrollBuilder) IsSecondarySearch(enabled bool) godsclient.ScrollBuilder {
	return sc
}

func (sc *redisScrollBuilder) buildRangeQuery(param types.RangeQueryParams) string {
	gt := "-inf"
	lt := "+inf"

	if param.Gt != nil {
		gt = fmt.Sprintf("%d", param.Gt)
	}

	if param.Gte != nil {
		gt = fmt.Sprintf("(%d", param.Gte)
	}

	if param.Lt != nil {
		lt = fmt.Sprintf("%d", param.Lt)
	}

	if param.Gt != nil {
		lt = fmt.Sprintf("(%d", param.Lte)
	}

	query := fmt.Sprintf("@%s:[%s %s]", param.Field, gt, lt)
	return query
}

func (sc *redisScrollBuilder) WithQuery(queryBuilder querybuilders.QueryBuilder) godsclient.ScrollBuilder {
	var queryString string

	//sc.query.AddFilter()

	q := queryBuilder.Build()

	switch v := q.Params.(type) {
	case types.FieldQueryParams:
		queryString = v.Field
	case types.FieldValueQueryParams:
		queryString = fmt.Sprintf("@%s:%s", v.Field, v.Value)
	case types.MatchQueryParams:

	case types.ScriptQueryParams:

	case types.RangeQueryParams:
		queryString = sc.buildRangeQuery(v)
	case types.DateRangeQueryParams:

	case types.GeoDistanceQueryParams:

	case types.FunctionScoreQueryParams:

	case types.NestedQueryParams:

	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	switch q.Type {
	case Type.And:
	case Type.Or:
	case Type.Not:
	case Type.Eq:
	case Type.Range:
	default:
	}

	sc.query = redisearch.NewQuery(queryString)

	return sc
}

func (sc *redisScrollBuilder) WithSize(size int32) godsclient.ScrollBuilder {
	sc.query.Limit(0, int(size))

	return sc
}

func (sc *redisScrollBuilder) WithContextID(contextID string) godsclient.ScrollBuilder {
	return sc
}

func (sc *redisScrollBuilder) WithContextTimeout(timeout time.Duration) godsclient.ScrollBuilder {
	return sc
}

func (sc *redisScrollBuilder) AddSort(sortBuilder sortbuilders.SortBuilder) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) AddSorts(sortBuilders ...sortbuilders.SortBuilder) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) AddAggregation(aggregationBuilder aggbuilders.AggregationBuilder) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) AddAggregations(aggregationBuilders ...aggbuilders.AggregationBuilder) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) AddProjection(field string) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) AddProjections(fields ...string) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) WithScoreQuery(queryBuilder querybuilders.QueryBuilder) godsclient.ScrollBuilder {
	return sc
}
func (sc *redisScrollBuilder) Execute() (*godsclient.ScrollResponse, error) {
	docs, total, err := sc.client.Search(sc.query)
	res := godsclient.ScrollResponse{}

	res.Documents = make(types.RawDocuments, total)

	for i := 0; i < total; i++ {
		res.Documents[i] = docs[i].Payload
	}

	return &res, err
}
func (sc *redisScrollBuilder) Build() *types.ScrollSearch {
	return nil
}
func (sc *redisScrollBuilder) String() string {
	return ""
}
