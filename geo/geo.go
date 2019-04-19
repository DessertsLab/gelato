package geo

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

type MapConfiger struct {
	Url       string
	Token     string
	Keyname   string
	ApiStruct ApiStructor
}

type ApiStructor interface {
	isSuccess() bool
	getInfo() string
	getLng() float64
	getLat() float64
}

func (q *QQAPI) isSuccess() bool {
	if q.Status == 0 {
		return true
	} else {
		return false
	}
}

func (b *BaiduAPI) isSuccess() bool {
	if b.Status == 0 {
		return true
	} else {
		return false
	}
}

func (q *QQAPI) getInfo() string {
	return "QQAPI"
}

func (b *BaiduAPI) getInfo() string {
	return "BaiduAPI"
}

func (q *QQAPI) getLng() float64 {
	return q.Result.Location.Lng
}

func (q *QQAPI) getLat() float64 {
	return q.Result.Location.Lat
}

func (b *BaiduAPI) getLng() float64 {
	return b.Result.Location.Lng
}

func (b *BaiduAPI) getLat() float64 {
	return b.Result.Location.Lat
}

func GetGeo(address string) (ApiStructor, error) {
	const DEFAULTLAT = 31.40527 // Shanghai
	const DEFAULTLNG = 121.48941
	var maps = []MapConfiger{
		MapConfiger{
			Url:       "https://apis.map.qq.com/ws/geocoder/v1/",
			Token:     getConf().Token.Qq,
			Keyname:   "key",
			ApiStruct: &QQAPI{},
		},
		MapConfiger{
			Url:       "http://api.map.baidu.com/geocoder/v2/",
			Token:     getConf().Token.Baidu,
			Keyname:   "ak",
			ApiStruct: &BaiduAPI{},
		},
	}

	for _, m := range maps {
		if b, err := apiCall(address, m.Url, m.Token, m.Keyname); err == nil {
			json.Unmarshal(b, &m.ApiStruct)
			if m.ApiStruct.isSuccess() {
				log.Printf("Success get data from %v \n", m.ApiStruct.getInfo())
				return m.ApiStruct, nil
			}
		} else {
			log.Println(err)
		}
	}

	log.Println("no success returned from api return default lat lng...")
	v := QQAPI{}
	v.Result.Location.Lat = DEFAULTLAT
	v.Result.Location.Lng = DEFAULTLNG
	return &v, nil
}

func apiCall(address string, apiurl string, token string, keyname string) ([]byte, error) {
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
	res, err := http.Get(urlPath)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

//func main() {
//	res, _ := GetGeo("南京环亚医疗美容门诊部有限公司")
//	fmt.Println(res.getLat(), res.getLng())
//}
