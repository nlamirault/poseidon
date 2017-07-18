// Copyright (C) 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tides

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

const (
	website = "http://maree.info"
)

func fetch(uri string, data url.Values) ([]byte, error) {
	u, _ := url.ParseRequestURI(uri)
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	glog.V(2).Infof("URI: %s %s", urlStr, data)

	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Http request to %s failed: %s", r.URL, err.Error())
	}
	glog.V(2).Infof("HTTP Status: %s", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, fmt.Errorf("Reading the body: %s", err.Error())
	}
	return body, nil
}

func ExtractHarbors() (map[string]string, error) {
	results := map[string]string{}
	glog.V(2).Infof("Search harbors")

	body, err := fetch(website, url.Values{})
	if err != nil {
		return nil, err
	}

	z := html.NewTokenizer(strings.NewReader(string(body)))
	for {
		tokenType := z.Next()
		if tokenType == html.ErrorToken {
			break
		}
		// token := z.Token()
		switch tokenType {
		case html.StartTagToken: // <tag>
			t := z.Token()
			if t.Data == "a" {
				if len(t.Attr) == 3 {
					// glog.V(3).Infof("Attributes: %s", t.Attr)
					if t.Attr[2].Val == "Port PP" || t.Attr[2].Val == "Port PS" {
						glog.V(3).Infof("Attributes: %s", t.Attr)
						id := t.Attr[0].Val
						glog.V(3).Infof("Harbor ID: %s", id)

						inner := z.Next()
						if inner == html.TextToken {
							text := (string)(z.Text())
							value := strings.TrimSpace(text)
							glog.V(3).Infof("Harbor Name: %s", value)
							results[id] = value
						}
					}
				}
			}
		}
	}
	glog.V(2).Infof("Harbors: %s", results)
	return results, nil
}
