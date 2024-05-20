//😜😜😜😜😝🤤😒😓

package helpers 

import (
	"os"
	"strings"
)

// Enforce HTTP ... 

func EnforceHTTP(url string) string {
	// making every url https 
	if url[:4] != "http"{
		return "http://" + url
	}

	return url
}

// remove Domain Error

func RemoveDomainError(url string) bool {
	// basically this functions remove all the commanly found 
	// prefixes from URL such as http, https, www
	// then check of the remaining string is the DOMAIN itself
	
	if url == os.Getenv("DOMAIN"){
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://","", 1)
	newURL = strings.Replace(newURL, "www." ,"", 1)
	newURL = strings.Split(newURL , "/") [0]

	if newURL == os.Getenv("DOMAIN"){
		return false
	}
	return true
}