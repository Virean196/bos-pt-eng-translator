package translator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	ResponseData struct {
		TranslatedText string  `json:"translatedText"`
		Match          float64 `json:"match"`
	} `json:"responseData"`
	QuotaFinished   bool        `json:"quotaFinished"`
	MtLangSupported interface{} `json:"mtLangSupported"`
	ResponseDetails string      `json:"responseDetails"`
	ResponseStatus  int         `json:"responseStatus"`
	ResponderID     interface{} `json:"responderId"`
	ExceptionCode   interface{} `json:"exception_code"`
	Matches         []struct {
		ID             any         `json:"id"`
		Segment        string      `json:"segment"`
		Translation    string      `json:"translation"`
		Source         string      `json:"source"`
		Target         string      `json:"target"`
		Quality        any         `json:"quality"`
		Reference      interface{} `json:"reference"`
		UsageCount     int         `json:"usage-count"`
		Subject        any         `json:"subject"`
		CreatedBy      string      `json:"created-by"`
		LastUpdatedBy  string      `json:"last-updated-by"`
		CreateDate     string      `json:"create-date"`
		LastUpdateDate string      `json:"last-update-date"`
		Match          float64     `json:"match"`
		Penalty        int         `json:"penalty"`
	} `json:"matches"`
}

func GetTranslation(input string, langPair string) (string, error) {
	var response Response
	if input == "" {
		return "", fmt.Errorf("No input selected, usage: bos <phrase>")
	}
	request := fmt.Sprintf("https://api.mymemory.translated.net/get?q=%s&langpair=%s", url.QueryEscape(input), langPair)
	resp, err := http.Get(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}
	return response.ResponseData.TranslatedText, nil
}
