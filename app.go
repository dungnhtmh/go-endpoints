package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Path("/account/{sns_account_id}/post/draft").Methods("GET").
		HandlerFunc(Draft)

	r.Path("/account/{sns_account_id}/post/detail/{post_id}").Methods("GET").
		HandlerFunc(DraftDetail)

	r.Path("/account/{sns_account_id}/post/reserve").Methods("GET").
		HandlerFunc(Reserve)

	r.Path("/account/{sns_account_id}/post/history").Methods("GET").
		HandlerFunc(History)

	r.Path("/organization/{organization_id}/media/list").Methods("GET").
		HandlerFunc(Media)

	r.Path("/organization/{organization_id}/media/detail/{media_id}").Methods("GET").
		HandlerFunc(MediaDetail)

	r.Path("/organization/{organization_id}/media/image/{media_id}").Methods("GET").
		HandlerFunc(GetImage)

	r.Path("/organization/{organization_id}/media/video/{media_id}").Methods("GET").
		HandlerFunc(GetVideo)

	r.Path("/organization/{organization_id}/media/uploader/list").Methods("GET").
		HandlerFunc(MediaUploader)

	r.Path("/organization/{organization_id}/media/uploader/{uploader_token}").Methods("GET").
		HandlerFunc(MediaUploaderDetail)

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
	var draftDataPath = "./json/post/draft.json"

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

// DraftDetail is an HTTP api: GET /account/{sns_account_id}/post/detail/{post_id}
func DraftDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	snsAccountId, err := strconv.Atoi(vars["sns_account_id"])
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	// TODO modify path to get draft data
	var draftDataPath = "./json/post/detail/1.json"

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

// Reserve is an HTTP api: GET /account/{sns_account_id}/post/reserve
func Reserve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	snsAccountId, err := strconv.Atoi(vars["sns_account_id"])
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	// TODO modify path to get draft data
	var reserveDataPath = "./json/post/reserve.json"

	// get data
	json, err := GetPostResponse(snsAccountId, reserveDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

// History is an HTTP api: GET /account/{sns_account_id}/post/history
func History(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	snsAccountId, err := strconv.Atoi(vars["sns_account_id"])
	if err != nil {
		w.WriteHeader(http.StatusCreated)
		return
	}

	// TODO modify path to get draft data
	var historyDataPath = "./json/post/history.json"

	// get data
	json, err := GetPostResponse(snsAccountId, historyDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

// Media is an HTTP api: GET /organization/{organization_id}/media/list
func Media(w http.ResponseWriter, r *http.Request) {

	// TODO modify path to get draft data
	var mediaDataPath = "./json/organization/media/all.json"

	// get data
	json, err := GetPostResponse(1, mediaDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

// MediaDetail is an HTTP api: GET /organization/{organization_id}/media/detail/{media_id}
func MediaDetail(w http.ResponseWriter, r *http.Request) {

	// TODO modify path to get draft data
	var mediaDataPath = "./json/organization/media/detail/1.json"

	// get data
	json, err := GetPostResponse(1, mediaDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

// GetImage to get image file Get /organization/{organization_id}/media/image/{media_id}
func GetImage(w http.ResponseWriter, r *http.Request) {

	// TODO temporary image for getting image function when database is not available
	fileName := "image.jpg"

	//Check if file exists and open
	openFile, err := os.Open("./json/organization/media/image/" + fileName)
	if err != nil {
		//File not found, send 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer openFile.Close()

	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	buffer := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	_, err = openFile.Read(buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get the file size
	fileStat, err := openFile.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Get info from file
	fileSize := strconv.FormatInt(fileStat.Size(), 10) //Get file size as a string

	//Send the headers
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Length", fileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	_, err = openFile.Seek(0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, openFile) //'Copy' the file to the client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetVideo to get video file Get: /organization/{organization_id}/media/video/{media_id}
func GetVideo(w http.ResponseWriter, r *http.Request) {

	// TODO temporary video for getting video function when database is not available
	fileName := "video.mp4"

	//Check if file exists and open
	openFile, err := os.Open("./json/organization/media/video/" + fileName)
	if err != nil {
		//File not found, send 404
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer openFile.Close()
	//File is found, create and send the correct headers

	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	buffer := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	_, err = openFile.Read(buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get the file size
	fileStat, err := openFile.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Get file size as a string
	fileSize := strconv.FormatInt(fileStat.Size(), 10)
	//Send the headers
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Length", fileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	_, err = openFile.Seek(0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(w, openFile) //'Copy' the file to the client
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// MediaUploader is an HTTP api: GET /organization/{organization_id}/media/uploader/list
func MediaUploader(w http.ResponseWriter, r *http.Request) {

	// TODO modify path to get draft data
	var mediaUploaderDataPath = "./json/organization/media/uploader/all.json"

	// get data
	json, err := GetPostResponse(1, mediaUploaderDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}

// MediaUploader is an HTTP api: GET /organization/{organization_id}/media/uploader/{uploader_token}
func MediaUploaderDetail(w http.ResponseWriter, r *http.Request) {

	// TODO modify path to get draft data
	var mediaUploaderDataPath = "./json/organization/media/uploader/detail/1.json"

	// get data
	json, err := GetPostResponse(1, mediaUploaderDataPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// output data
	w.Header().Add("Content-Type", "application/json")
	w.Write(json)
}
