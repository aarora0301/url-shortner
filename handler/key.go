package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poc/url-shortner/models"
	"github.com/poc/url-shortner/repository"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Status       string `json:"status"`
	Code         int    `json:"code"`
	Hash         string `json:"hash, omitEmpty"`
	ErrorMessage string `json:"error_message, omitEmpty"`
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
	str := "http://" + r.Host + "/toko/" + hash
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
