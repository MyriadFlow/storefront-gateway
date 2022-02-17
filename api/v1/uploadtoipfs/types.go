package uploadtoipfs

type UploadToIpfsPayload struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}
