package ecs

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"go-uuid/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"
)

func BuildUrl(m map[string]string) string {
	m["Version"] = Version
	m["SignatureMethod"] = SignatureMethod
	m["Timestamp"] = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	m["SignatureVersion"] = SignatureVersion
	m["AccessKeyId"] = C.AccessKeyId
	m["SignatureNonce"] = uuid.NewRandom().String()
	m["Format"] = "json"
	urlstr := ""
	for k, v := range m {
		if v == "" {
			delete(m, k)
			continue
		}
		if urlstr != "" {
			urlstr += "&"
		}
		urlstr += k + "=" + url.QueryEscape(v)
	}
	return Url + urlstr + "&Signature=" + buildSignature(m)
}

func buildSignature(m map[string]string) (signature string) {
	ms := NewMapSort(m)
	sort.Sort(ms)
	encodeStr := "GET&%2F&"
	cqs := ""
	for _, v := range ms {
		if cqs != "" {
			cqs += "%26"
		}
		cqs += url.QueryEscape(v.key) + "%3D" + url.QueryEscape(url.QueryEscape(v.value))
	}
	mac := hmac.New(sha1.New, []byte(C.AccessSecret+"&"))
	mac.Write([]byte(encodeStr + cqs))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}

func doInvoke2(m map[string]string, r interface{}) ErrorResult {
	url := BuildUrl(m)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return ErrorResult{Message: err.Error()}
	}
	text, _ := ioutil.ReadAll(resp.Body)
	println(string(text))
	if resp.StatusCode != 200 {
		errR := ErrorResult{}
		if err := json.Unmarshal(text, &errR); err == nil {
			return errR
		}
	}
	if err := json.Unmarshal(text, r); err == nil {
		return ErrorResult{}
	} else {
		return ErrorResult{Message: err.Error()}
	}
}

func doInvoke(m map[string]string, r interface{}) error {
	url := BuildUrl(m)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return ErrorResult{Message: err.Error()}
	}
	text, _ := ioutil.ReadAll(resp.Body)
	errR := ErrorResult{}
	if resp.StatusCode != 200 {
		json.Unmarshal(text, &errR);
		return errR;
	}
	if err := json.Unmarshal(text, &r); err == nil {
		return nil
	} else {
		errR.Message = err.Error()
		return errR;
	}
}

func NewMapSort(m map[string]string) MapSort {
	ms := make(MapSort, len(m))
	i := 0
	for k, v := range m {
		ms[i] = MapItem{key: k, value: v}
		i++
	}
	return ms
}

type MapItem struct {
	key   string
	value string
}

type MapSort []MapItem

func (m MapSort) Len() int {
	return len(m)
}

func (m MapSort) Less(i, j int) bool {
	return m[i].key < m[j].key
}

func (m MapSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
