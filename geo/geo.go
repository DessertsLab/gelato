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

type configuration struct {
	Token struct {
		Baidu  string `json:"baidu"`
		Mapbox string `json:"mapbox"`
		Qq     string `json:"qq"`
	} `json:"token"`
}

// BaiduAPI is return struct of baidu geo api
type baiduAPI struct {
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

// QQAPI is reaturn struct of QQ geo api
type qqAPI struct {
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

func getConf(args ...string) configuration {
	var path string
	if len(args) != 0 {
		path = args[0]
	} else {
		path = "conf/config.json"
	}

	file, openfileerr := os.Open(path)
	defer file.Close()
	if openfileerr != nil {
		log.Println(openfileerr)
	}
	decoder := json.NewDecoder(file)
	configuration := configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

type mapConfiger struct {
	URL       string
	Token     string
	Keyname   string
	APIStruct apiStructor
}

type apiStructor interface {
	isSuccess() bool
	GetInfo() string
	GetLng() float64
	GetLat() float64
}

func (q *qqAPI) isSuccess() bool {
	if q.Status == 0 {
		return true
	}
	return false
}

func (b *baiduAPI) isSuccess() bool {
	if b.Status == 0 {
		return true
	}
	return false
}

func (q *qqAPI) GetInfo() string {
	return "QQAPI"
}

func (b *baiduAPI) GetInfo() string {
	return "BaiduAPI"
}

func (q *qqAPI) GetLng() float64 {
	return q.Result.Location.Lng
}

func (q *qqAPI) GetLat() float64 {
	return q.Result.Location.Lat
}

func (b *baiduAPI) GetLng() float64 {
	return b.Result.Location.Lng
}

func (b *baiduAPI) GetLat() float64 {
	return b.Result.Location.Lat
}

// GetGeo can get geo info from internet it return interface apiStructor
func GetGeo(address string) (apiStructor, error) {
	const DEFAULTLAT = 31.40527 // Shanghai
	const DEFAULTLNG = 121.48941
	var maps = []mapConfiger{
		mapConfiger{
			URL:       "https://apis.map.qq.com/ws/geocoder/v1/",
			Token:     getConf().Token.Qq,
			Keyname:   "key",
			APIStruct: &qqAPI{},
		},
		mapConfiger{
			URL:       "http://api.map.baidu.com/geocoder/v2/",
			Token:     getConf().Token.Baidu,
			Keyname:   "ak",
			APIStruct: &baiduAPI{},
		},
	}

	for _, m := range maps {
		if b, err := apiCall(address, m.URL, m.Token, m.Keyname); err == nil {
			json.Unmarshal(b, &m.APIStruct)
			if m.APIStruct.isSuccess() {
				log.Printf("Success get data from %v \n", m.APIStruct.GetInfo())
				return m.APIStruct, nil
			}
		} else {
			log.Println(err)
		}
	}

	log.Println("no success returned from api return default lat lng...")
	v := qqAPI{}
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
