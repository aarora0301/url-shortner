package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poc/url-shortner/models"
	"github.com/poc/url-shortner/repository"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Status       string `json:"status"`
	Code         int    `json:"code"`
	Hash         string `json:"hash, omitEmpty"`
	ErrorMessage string `json:"error_message, omitEmpty"`
	OriginalURL  string `json:"original_url, omitEmpty"`
}

/**
Generates shortened url
*/
func GetHash(w http.ResponseWriter, r *http.Request) {
	result, err := repository.GetAvailableKey()
	if err != nil {
		log.Print("Error occurred while generating hash", err)
		json.NewEncoder(w).Encode(Response{Status: "Failed", Code: http.StatusBadGateway, ErrorMessage: error.Error(err)})
		return
	}

	if len(result) <= 0 {
		message := "Error occurred while fetching hash"
		log.Println(message)
		json.NewEncoder(w).Encode(Response{Status: "Failed", Code: http.StatusBadGateway, ErrorMessage: message})
		return
	}
	var url models.Url
	var hash string
	hash = result[0].Key
	url.Hash = hash
	now := time.Now()
	url.CreationDate = now
	/**
	Set Expiry time to 10 hours after creation date
	*/
	url.ExpirationDate = now.Add(10 * time.Hour)
	body, _ := ioutil.ReadAll(r.Body)
	url.OriginalUrl = string(body)
	repository.CreatUrl(url)
	str := "http://" + r.Host + "/tokopedia/" + hash
	json.NewEncoder(w).Encode(Response{Status: "Success", Code: 200, Hash: str})
	return
}

/**
Redirects to  url mapped to key {fetched from request}
*/
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	params = mux.Vars(r)
	hash := params["pattern"]
	var url []models.Url
	var err error
	url, err = repository.GetOriginalURL(hash)
	if err != nil {
		message := "Error occurred while redirecting " + error.Error(err)
		log.Println(message)
		json.NewEncoder(w).Encode(Response{Status: "Failed", Code: http.StatusBadGateway, ErrorMessage: message})
		return
	}
	var str string
	if len(url) <= 0 {
		message := "Error occurred while parsing from database  " + error.Error(err)
		log.Println(message)
		json.NewEncoder(w).Encode(Response{Status: "Failed", Code: http.StatusBadGateway, ErrorMessage: message})
		return
	}
	str = url[0].OriginalUrl
	http.Redirect(w, r, str, http.StatusMovedPermanently)

	var client *http.Client
	client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // or maybe the error from the request
		},
	}
	log.Println("Redirecting client ", client)
}

func GetURL(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	params = mux.Vars(r)
	hash := params["pattern"]
	var url []models.Url
	var err error
	url, err = repository.GetOriginalURL(hash)
	if err != nil {
		message := "Error occurred while fetching URL for hash " + hash + error.Error(err)
		log.Println(message)
		json.NewEncoder(w).Encode(Response{Status: "Failed", Code: http.StatusBadGateway, ErrorMessage: message})
		return
	}

	processedUrl := processURL(url[0].OriginalUrl, r.UserAgent())
	encode := json.NewEncoder(w)
	encode.SetEscapeHTML(false)
	encode.Encode(Response{Status: "Success", Code: http.StatusCreated, OriginalURL: processedUrl})

	return
}

func getOS(userAgent string) (os string) {
	userAgent = strings.ToLower(userAgent)
	if strings.Contains(userAgent, "android") {
		os = "android"
	} else if strings.Contains(userAgent, "ios") {
		os = "ios"
	}
	return
}

func processURL(url, userAgent string) (result string) {
	var urlArray [] string
	urlArray = strings.Split(url, "?")
	os := getOS(userAgent)
	result = urlArray[0] + "?" + generateURL(urlArray, os)
	return
}

func generateURL(urlArray []string, os string) (result string) {
	urlArray = strings.Split(urlArray[1], "&")
	var general []string
	var ios []string
	var android []string
	for _, keys := range urlArray {
		if strings.Contains(keys, "android") {
			android = append(android, keys)
		} else if strings.Contains(keys, "ios") {
			ios = append(ios, keys)
		} else {
			general = append(general, keys)
		}
	}

	if os == "android" {
		android = append(android, general...)
		result += strings.Join(android[:], "&")
	} else if os == "ios" {
		ios = append(ios, general...)
		result += strings.Join(ios, "&")
	} else {
		result += strings.Join(general, "&")
	}
	return
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/branch.html")
}
