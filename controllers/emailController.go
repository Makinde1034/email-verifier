package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"github.com/Makinde1034/email-verifier/models"
	"github.com/Makinde1034/email-verifier/utils"
) 

func VerifyEmail(w http.ResponseWriter, r *http.Request){

	request := models.VerificationRequest{}
	validEmails := []models.VerificatopnResponse{}
	json.NewDecoder(r.Body).Decode(&request)

	wg := &sync.WaitGroup{}

	for _, email := range request.Emails {

		wg.Add(1)
		go func(email string) {
		
			hasMx, hasSPF, hasDMARC := utils.VerifyEmail(email)
			fmt.Println(hasMx,hasSPF,hasDMARC,email)

			validMail := models.VerificatopnResponse{
				HasMX: hasMx,
				HasSPF: hasSPF,
				HasDMARC: hasDMARC,
				Email: email,
				IsValid: hasMx,
				
			}
			validEmails = append(validEmails, validMail)
			wg.Done()

		}(email)
		

	}
	wg.Wait()

	responseData := models.ResponseData{
		Data : validEmails,
	}

	json.NewEncoder(w).Encode(responseData)
	w.WriteHeader(http.StatusOK)
}