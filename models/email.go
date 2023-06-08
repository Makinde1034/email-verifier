package models 

type VerificationRequest struct{
	Email string `json:"email"`
}

type VerificatopnResponse struct{
	HasMX bool `json:"hasMX"`
	HasDMARC bool `json:"hasDMARC"`
	HasSPF bool `josln:"hasSpf"`	
}

type ResponseData struct{
	Data VerificatopnResponse `json:"data"`
}