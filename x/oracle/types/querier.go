package types

import "github.com/cosmos/cosmos-sdk/types"

const (
	QueryListSources = "listsources"
	QuerySearchVotes = "searchvotes"
	QuerySearchVoteKeys = "searchvotekeys"
	QueryCalculateProofSig = "calculateVoteProofSig"
	QueryGetValcons = "getValcons"
	QuerySearchSourceValues = "searchSourceValues"
)

type QueryResultListSources = []struct{
	Name string
	Url string
	Property string
	Owner types.AccAddress
}


type SearchVotesQueryRequest = struct{
	Prefix string
}

type SearchSourceValuesQueryRequest = struct{
	Prefix string
}


type CalculateProofSigQueryRequest = struct {
	Valcons string
	Value string
}