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
	htmlutils "html"
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

func even(number int) bool {
	return number%2 == 0
}

func getFormatedTides(tides [][]string) string {
	var buffer bytes.Buffer
	for _, day := range tides {
		glog.V(4).Infof("Day: %s", day)
		for i := 2; i < len(day); i++ {
			if len(day[i]) > 0 {
				if even(i) {
					buffer.WriteString(fmt.Sprintf("PM: %s ", day[i]))
				} else {
					buffer.WriteString(fmt.Sprintf("BM: %s ", day[i]))
				}
			}

		}
		buffer.WriteString("\n")
	}
	return strings.TrimSpace(buffer.String())
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
						id := strings.TrimPrefix(t.Attr[0].Val, "/")
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

func DescribeHarbor(id string) (map[string]string, error) {
	results := map[string]string{}
	glog.V(2).Infof("Describe harbor: %s", id)

	body, err := fetch(fmt.Sprintf("%s/%s", website, id), url.Values{})
	if err != nil {
		return nil, err
	}

	inTides := false
	tides := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", ""}
	weekTides := [][]string{}
	notFinished := true
	i := 0

	z := html.NewTokenizer(strings.NewReader(string(body)))
	for {

		tokenType := z.Next()
		if tokenType == html.ErrorToken {
			break
		}
		// token := z.Token()
		switch tokenType {
		case html.SelfClosingTagToken: // <tag>
			t := z.Token()
			if t.Data == "meta" {
				glog.V(3).Infof("Meta: %s", t)
				if len(t.Attr) == 2 {
					glog.V(3).Infof("Attributes: %s", t.Attr[0].Val)
					if t.Attr[0].Val == "latitude" || t.Attr[0].Val == "longitude" {
						results[t.Attr[0].Val] = t.Attr[1].Val

					}
				}
			}
		case html.StartTagToken: // <tag>
			t := z.Token()
			if t.Data == "h2" {
				if len(t.Attr) == 3 {
					if t.Attr[0].Val == fmt.Sprintf("Port%s", id) {
						glog.V(3).Infof("Harbor name: %s", t)
						inner := z.Next()
						if inner == html.TextToken {
							text := (string)(z.Text())
							value := strings.TrimSpace(text)
							glog.V(3).Infof("Harbor Name: %s", value)
							results["name"] = value
						}
					}
				}
			}
		case html.EndTagToken: // </tag>
			t := z.Token()
			if t.Data == "td" {

			} else if t.Data == "tr" {
				if i > 0 {
					glog.V(3).Infof(">> Values: %s ||||| %s", weekTides, tides)
					weekTides = append(weekTides, tides)
					if inTides {
						tides = []string{"", "", "", "", "", "", "", "", "", "", "", "", "", ""}
						i = 0
					}
				}
			}
		case html.TextToken:
			t := z.Token()
			if notFinished {
				text := strings.TrimSpace(htmlutils.UnescapeString(t.Data))
				glog.V(3).Infof(">>>>>>>>>>>>>>> Text: [%s]", text)
				if inTides && notFinished {
					glog.V(3).Infof("Ajout de : %s %d", t.Data, i)
					// tides = append(tides, t.Data)
					tides[i] = text
					i++
				}
				if t.Data == "Coeff." {
					glog.V(3).Infof("Start tides -------------------")
					inTides = true
				} else if t.Data == "@maree_info_136" {
					notFinished = false
					tides = []string{"", "", "", "", "", "", "", "", "", "", "", "", "", ""}
					i = 0
				}
			}
		}
	}
	results["tides"] = getFormatedTides(weekTides)
	glog.V(2).Infof("Week Tides: %s", weekTides)
	glog.V(2).Infof("Harbors: %s", results)
	return results, nil
}
