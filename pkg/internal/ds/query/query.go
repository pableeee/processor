package query

type Type string

const (
	And           Type = "and"
	Or            Type = "or"
	Not           Type = "not"
	Eq            Type = "eq"
	Like          Type = "like"
	Regexp        Type = "regexp"
	Match         Type = "match"
	Exists        Type = "exists"
	In            Type = "in"
	Range         Type = "range"
	DateRange     Type = "date_range"
	GeoDistance   Type = "geo_distance"
	Script        Type = "script"
	Ids           Type = "ids"
	FunctionScore Type = "function_score"
	Nested        Type = "nested"
	IP            Type = "ip"
)
