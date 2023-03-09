package models

type NftLikes struct {
	Cid string `json:"cid"`
	Likes uint64 `json:"likes_count"`
}