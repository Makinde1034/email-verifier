package models

type VerificationRequest struct {
	Emails []string `json:"emails"`
}

type VerificatopnResponse struct {
	HasMX    bool `json:"hasMX"`
	HasDMARC bool `json:"hasDMARC"`
	HasSPF   bool `json:"hasSpf"`
	Email string `json:"email"`
	IsValid bool `json:"isValid"`
	
}

type ResponseData struct {
	Data []VerificatopnResponse `json:"data"`
}