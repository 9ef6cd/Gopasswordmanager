package main

import (
	"fmt"
	"time"
	"crypto/sha256"		
)

func Passgen() string {
	cTime := time.Now()
	fTime := cTime.Format("2006-01-02 15:04:05.000000")
		
	h := sha256.New()
	h.Write([]byte(fTime))
	
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash[:16]
}

