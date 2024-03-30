package models

type DelegateAsset struct {
	Id                       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	MetaDataHash             string `json:"metaDataHash"`
	RoyaltyPercentBasisPoint int    `json:"royaltyPercentBasisPoint"`
	StorefrontId             string `json:"storefrontId"`
	ContractAddress          string `json:"contractAddress"`
}
