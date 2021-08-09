package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Path("/welcome").Methods("GET").
		HandlerFunc(Welcome)

	r.Path("/account/{sns_account_id}/post/draft").Methods("GET").
		HandlerFunc(Draft)

	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// corsHandler wraps a HTTP handler and applies the appropriate responses for Cross-Origin Resource Sharing.
type corsHandler http.HandlerFunc

func (h corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		return
	}
	h(w, r)
}

// List is common and represents list of post and it's info
type List struct {
	Data []Post `json:"data"`
}

// Post represents single post data
type Post struct {
	ID               int      `json:"id"`
	ContributorsName string   `json:"contributors_name"`
	ExecutorName     string   `json:"executor_name"`
	ReserveAt        string   `json:"reserve_at"`
	PostedAt         string   `json:"posted_at"`
	MediaIdList      []int    `json:"media_id_list"`
	Body             string   `json:"body"`
	Objective        []string `json:"objective"`
	Category         []string `json:"category"`
	Status           string   `json:"status"`
	Workflow         string   `json:"workflow"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}

// TODO present sample sns account id list when data fromdatabase is not avaiable
var snsAccountIdList = []int{1, 2, 3, 4, 5, 6, 7}

// GetPostResponse get post data and response json
func GetPostResponse(snsAccountId int, path string) (json []byte, err error) {
	// if organization_id not found
	if !IsIdExist(snsAccountIdList, snsAccountId) {
		return nil, errors.New("sns account id not found")
	}

	// get json data
	json, err = ReadJsonFile(path)
	return
}

// IsIdExist check if modified id exist or not
func IsIdExist(idList []int, id int) bool {
	for _, v := range idList {
		if v == id {
			return true
		}
	}
	return false
}

// ReadJsonFile read json inout file and return byte value
func ReadJsonFile(filePath string) (json []byte, err error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer jsonFile.Close()

	json, err = ioutil.ReadAll(jsonFile)
	return
}

// Draft is an HTTP api: GET /account/{sns_account_id}/post/draft
func Draft(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	snsAccountId, err := strconv.Atoi(vars["sns_account_id"])
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	// TODO modify path to get draft data
	var draftDataPath = "./json/post/post_draft.json"

	// get data
	json, err := GetPostResponse(snsAccountId, draftDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	// TODO modify path to get draft data
	var draftDataPath = "./json/post/post_draft.json"

	// get data
	json, err := GetPostResponse(1, draftDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}
