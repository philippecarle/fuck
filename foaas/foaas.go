package foaas

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type EndPoint struct {
	Fields []Field
	Name   string
	Url    string
}

type Field struct {
	Field string
	Name  string
}

type Params struct {
	Name string
	From string
}

type FuckOff struct {
	Message  string
	Subtitle string
}

const BaseUrl = "http://foaas.com"

func getRandomEndPoint() (EndPoint, error) {
	r, err := http.Get(BaseUrl + "/operations")

	defer r.Body.Close()

	if err == nil && r.StatusCode == http.StatusOK {
		var EndPoints []EndPoint
		var e []EndPoint

		decoder := json.NewDecoder(r.Body)

		decoder.Decode(&EndPoints)

		t := time.Now()
		rand.Seed(int64(t.Nanosecond()))

		for _, endpoint := range EndPoints {
			if strings.Contains(endpoint.Url, ":name") && !strings.Contains(endpoint.Url, ":company") && !strings.Contains(endpoint.Url, ":reference") {
				e = append(e, endpoint)
			}
		}

		return e[rand.Int()%len(e)], nil
	}

	return EndPoint{}, err
}

func GetTheFuck(who string, me string) (*FuckOff, error) {
	e, _ := getRandomEndPoint()

	p := Params{who, me}

	url := e.Url

	for _, f := range e.Fields {
		url = strings.Replace(url, ":"+f.Field, getField(p, f.Name), 1)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", BaseUrl+url, nil)

	if err != nil {
		panic(err.Error())
	}

	req.Header.Add("Accept: ", "application/json")
	res, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := getMessage([]byte(body))

	return s, err
}

func getField(p Params, field string) string {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}

func getMessage(body []byte) (*FuckOff, error) {
	var s = new(FuckOff)
	err := json.Unmarshal(body, &s)
	if err != nil {
		panic(err.Error())
	}
	return s, err
}
