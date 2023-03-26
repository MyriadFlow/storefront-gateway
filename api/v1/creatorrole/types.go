package creatorrole

type CreatorRoleRequest struct {
	RoleId string `json:"roleid"`
}

type CreatorRolePayload struct {
	TransactionHash string `json:"transactionHash"`
}
