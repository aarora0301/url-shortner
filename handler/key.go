package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poc/url-shortner/models"
	"github.com/poc/url-shortner/repository"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Hash   string `json:"hash"`
}

func GetHash(w http.ResponseWriter, r *http.Request) {
	result := repository.GetAvailableKey()
	var url models.Url
	var hash string
	hash = result[0].Key
	url.Hash = hash
	now := time.Now()
	url.CreationDate = now
	url.ExpirationDate = now.Add(10 * time.Hour)
	body, _ := ioutil.ReadAll(r.Body)
	url.OriginalUrl = string(body)
	repository.CreatUrl(url)
	str := "http://toko/" + hash
	json.NewEncoder(w).Encode(Response{Status: "success", Code: 200, Hash: str})
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	var params map[string]string
	params = mux.Vars(r)
	hash := params["pattern"]
	var url []models.Url
	url = repository.GetOriginalURL(hash)
	var str string
	str = url[0].OriginalUrl
	http.Redirect(w, r, str, 301)

//	var client *http.Client
//client= &http.Client{
//	CheckRedirect: func(req *http.Request, via []*http.Request) error {
//
//		return errors.New("something bad happened") // or maybe the error from the request
//	},
//}
}
