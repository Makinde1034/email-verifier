package utils

import (
	"fmt"
	"net"
	"strings"
) 


func VerifyEmail(email string) (bool, bool, bool){
	var hasMX, hasSPF, hasDMARC  bool = false, false, false
	var spfRecord, dmarcRecord string

	domain := strings.Split(email,"@")[1]

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


	fmt.Printf("%v, %v, %v, %v, %v,",hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord)

	return hasMX, hasSPF, hasDMARC
}

func ValidateEmail(){
	
}