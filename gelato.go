package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Configuration struct {
	Token struct {
		Baidu  string `json:"baidu"`
		Mapbox string `json:"mapbox"`
		Qq     string `json:"qq"`
	} `json:"token"`
}

type BaiduAPI struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		Precise       int    `json:"precise"`
		Confidence    int    `json:"confidence"`
		Comprehension int    `json:"comprehension"`
		Level         string `json:"level"`
	} `json:"result"`
}

type QQAPI struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Title    string `json:"title"`
		Location struct {
			Lng float64 `json:"lng"`
			Lat float64 `json:"lat"`
		} `json:"location"`
		AdInfo struct {
			Adcode string `json:"adcode"`
		} `json:"ad_info"`
		AddressComponents struct {
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"address_components"`
		Similarity  float64 `json:"similarity"`
		Deviation   int     `json:"deviation"`
		Reliability int     `json:"reliability"`
		Level       int     `json:"level"`
	} `json:"result"`
}

func getConf() Configuration {
	file, openfileerr := os.Open("conf/config.json")
	defer file.Close()
	if openfileerr != nil {
		log.Println(openfileerr)
	}
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func (b BaiduAPI) GetGeo(address string) (BaiduAPI, error) {
	var URL string = "http://api.map.baidu.com/geocoder/v2/"
	var TOKEN string = getConf().Token.Baidu
	var KEYNAME string = "ak"
	err := apiCall(address, URL, TOKEN, KEYNAME, &b)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (b QQAPI) GetGeo(address string) (QQAPI, error) {
	var URL string = "https://apis.map.qq.com/ws/geocoder/v1/"
	var TOKEN string = getConf().Token.Qq
	var KEYNAME string = "key"
	err := apiCall(address, URL, TOKEN, KEYNAME, &b)
	if err != nil {
		return b, err
	}
	return b, nil
}

func apiCall(address string, apiurl string, token string, keyname string, b interface{}) error {
	params := url.Values{}

	Url, err := url.Parse(apiurl)
	if err != nil {
		panic(err.Error())
	}

	params.Set("address", address)
	params.Set("ret_coordtype", "")
	params.Set(keyname, token)
	params.Set("sn", "")
	params.Set("precise", "1")
	params.Set("output", "json")
	params.Set("callback", "")
	params.Set("region", "")

	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	res, err := http.Get(urlPath)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer res.Body.Close()
		s, err := ioutil.ReadAll(res.Body)
		log.Println(string(s))
		if err != nil {
			panic(err.Error())
		}
		e := json.Unmarshal(s, &b)
		if e != nil {
			panic(e.Error())
		}
	}
	return err
}

func main() {
	a := BaiduAPI{}
	res, err := a.GetGeo("上海市黄浦区")
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("维度：%f, 经度：%f\n", res.Result.Location.Lat, res.Result.Location.Lng)
	}

}
