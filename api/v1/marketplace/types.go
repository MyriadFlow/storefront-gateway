package marketplace

type MarketPlaceCounts struct {	
	Actions		int `json:"action,omitempty"`
	Artwork			int `json:"artwork,omitempty"`
	Artists			int `json:"artists,omitempty"`
}

type MarketPlaceInfo struct {	
	Name		string `json:"name,omitempty"`
	Description			string `json:"description,omitempty"`
	Contact			string `json:"contact,omitempty"`
}


type LikesQueryPayload struct {	
	NFT_Contract_Address     string   `json:"nft-contract-address"`
	TokenId            string   `json:"tokenId,omitempty"`
}