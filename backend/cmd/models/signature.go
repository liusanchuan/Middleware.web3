package models

type SyncMint struct {
	Address1 string `json:"address_1"`
	Address2 string `json:"address_2"`
	Id       int    `json:"id"`
}

type CrossTransfer struct {
	Sender         string `json:"sender"`
	Receiver       string `json:"receiver"`
	Id             int    `json:"id"`
	ReceiveChainId int    `json:"receiveChainId"`
}

type CrossReceive struct {
	Sender        string `json:"sender"`
	Receiver      string `json:"receiver"`
	Id            int    `json:"id"`
	SenderChainId int    `json:"senderChainId"`
}
