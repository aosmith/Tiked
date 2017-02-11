package main

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"strings"
)

func PastebinPut(text, title string) (string, err error) {
	data := url.Values{}
	// Required values.
	data.Set("api_dev_key", "")
	data.Set("api_option", "paste") // Create a paste.
	data.Set("api_paste_code", text)
	// Optional values.
	data.Set("api_paste_name", title)      // The paste should have title "title".
	data.Set("api_paste_private", "2")     // Create a private paste.
	data.Set("api_paste_expire_date", "N") // The paste should never expire.

	resp, err := http.PostForm("http://pastebin.com/api/api_post.php", data)
	if err != nil || resp.StatusCode != 200 {
		return "", err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return StripURL(string(respBody)), nil
}

func StripURL(url string) string {
	return strings.Replace(url, "http://pastebin.com/", "", -1)
}

func PastebinRemove(id string) {
	data := url.Values{}
	data.Set("api_dev_key", "84484aa46e56da51609c7ffe8a5afc76")
	data.Set("api_option", "delete") // Create a paste.
	data.Set("api_paste_key", id)
	http.PostForm("http://pastebin.com/api/api_post.php", data)
}

