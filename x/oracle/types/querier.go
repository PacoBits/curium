package types

// Query endpoints supported by the oracle querier
const (
	QueryListSources = "listsources"
)

//type QueryResultListSources = []struct{
//	Name string `json:"name"`
//	Url string `json:"url"`
//	Property string `json:"property"`
//	Owner types.AccAddress `json:"owner"`
//}

type QueryResultListSources = []string
