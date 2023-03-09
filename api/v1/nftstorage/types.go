package nftstorage

type NftStorageUploadResponse struct {
	FileName string `json:"name"`
	CID string `json:"cid"`
}

type LikeNft struct {
	CID string `json:"cid"`
}


