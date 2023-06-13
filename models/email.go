package models

type VerificationRequest struct {
	Emails []string `json:"emails"`
}

type VerificatopnResponse struct {
	HasMX    bool `json:"hasMX"`
	HasDMARC bool `json:"hasDMARC"`
	HasSPF   bool `json:"hasSpf"`
	Email string `json:"email"`
	
}

type ResponseData struct {
	Data []VerificatopnResponse `json:"data"`
}