package xeno_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

type Query struct {
	Search     string
	Parameters map[string]string
}

func (q *Query) GetPage(page int) Response {

	// Get Url for page
	newUrl, err := q.formUrl()
	if err != nil {
		log.Fatal("err")
	}

	// Get page content
	log.Printf("Processing page %v", page)
	recordings, err := q.getResponse(newUrl, page)
	if err != nil {
		log.Fatal("err")
	}

	return recordings
}

func (q *Query) GetRecordings() []Recording {

	page1 := q.GetPage(1)
	recs := append([]Recording{}, page1.Recordings...)

	// Get existing uploaded files
	var w bytes.Buffer
	err := listFilesWithPrefix(&w, "alexandria101", "xenocanto/index/", "")
	if err != nil {
		log.Fatal("error")
	}
	fileListString := w.String()

	// Create list of existing IDs
	fileList := strings.Split(fileListString, "\n")
	var idList []int
	for _, v := range fileList[1 : len(fileList)-1] {
		vIdSuffix := strings.Split(v, "/")
		vId, err := strconv.Atoi(strings.Split(vIdSuffix[2], ".")[0])
		if err != nil {
			log.Fatal()
		}
		idList = append(idList, vId)
	}

	// Get page content and load
	for i := 2; i <= page1.NumPages; i++ {
		page := q.GetPage(i)
		recs = append(recs, page.Recordings...)

		// Load page recordings
		err := q.loadRecording(page.Recordings, idList)
		if err != nil {
			log.Fatal("err")
		}
	}
	return recs
}

func (q *Query) loadRecording(recs []Recording, idList []int) error {

	// Loop over recordings
	for _, v := range recs {

		b, err := json.Marshal(v)
		if err != nil {
			return err
		}

		// Read each recording
		r := strings.NewReader(string(b))

		// Check for existing upload
		vId, err := strconv.Atoi(v.Id)
		if err != nil {
			log.Fatal("err")
		}

		// Upload file
		i := sort.SearchInts(idList, vId)
		if i == 0 {
			var w bytes.Buffer
			fmt.Printf("\tUpload id: %v\n", vId)
			err = uploadFile(&w, r, "alexandria101", fmt.Sprintf("xenocanto/index/%v.json", v.Id))
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (q *Query) formUrl() (string, error) {
	baseUrl, err := url.Parse("https://www.xeno-canto.org/api/2/recordings")
	if err != nil {
		log.Fatal("err")
	}

	// Form parameters
	queryValues := ""
	for key, value := range q.Parameters {
		queryValues += fmt.Sprintf("%v:%v", key, value)
		queryValues += "+"
	}
	queryValues = strings.TrimSuffix(queryValues, "+")
	if len(q.Search) > 0 {
		queryValues = q.Search + "+" + queryValues
	}

	// Load parameters into query
	var query = url.Values{}
	query.Add("query", queryValues)

	// Load url with params
	baseUrl.RawQuery = query.Encode()
	newUrl, err := url.QueryUnescape(baseUrl.String())
	if err != nil {
		log.Fatal("err")
	}
	return newUrl, nil
}

func (q *Query) getResponse(url string, page int) (Response, error) {

	// Append page number
	if page > 1 {
		url += fmt.Sprintf("&page=%v", page)
	}

	// Get response from url
	r, err := http.Get(url)
	if err != nil {
		log.Fatal("get url failed")
	}

	// Read body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("err")
	}

	// Insert into struct
	resp := Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatal(err.Error())
	}

	return resp, nil
}
