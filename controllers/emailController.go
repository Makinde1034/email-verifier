package controllers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/Makinde1034/email-verifier/models"
) 

func VerifyEmail(w http.ResponseWriter, r *http.Request){
	var hasMX, hasSPF, hasDMARC  bool = false, false, false
	var spfRecord, dmarcRecord string

	request := models.VerificationRequest{}

	json.NewDecoder(r.Body).Decode(&request)
	domain := strings.Split(request.Email,"@")[1]

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		fmt.Printf("An error occured while check MX%v", err)
		hasMX = false
	}

	if (len(mxRecord)>0){
		hasMX =  true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		fmt.Printf("An error occured while checking TXT %v", err)
		
	}

	for _, record := range txtRecords {
		if (strings.HasPrefix(record,"v=spf1")){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, _err := net.LookupTXT("_dmarc." + domain)

	if _err != nil {
		fmt.Printf("An error occured occured while checking dmarc %v",_err)
	}

	for _,record := range dmarcRecords {
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	response := models.VerificatopnResponse{
		HasMX:  hasMX,
		HasDMARC: hasDMARC,
		HasSPF: hasSPF,
	}

	responseData := models.ResponseData{
		Data : response,
	}

	json.NewEncoder(w).Encode(responseData)
	w.WriteHeader(http.StatusOK)

	fmt.Printf("%v, %v, %v, %v, %v,",hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord)


}