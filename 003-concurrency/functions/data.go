package functions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Year int

func WomenRelatedCrimeStats(y Year) {

	fmt.Println("Fetching for year: ", y)
	dataId := map[Year]int{
		2015: 3957761,
		2014: 681321,
		2013: 135789,
	}

	id, ok := dataId[y]
	if !ok {
		fmt.Printf("Data for the year %v not available\n", y)
		return
	}

	url := "https://data.gov.in/node/" + strconv.Itoa(id) + "/datastore/export/json"

	r, _ := http.NewRequest("GET", url, nil)
	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("Error while getting the data:", err.Error())
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	//fmt.Println(string(body))
	fmt.Printf("Got Response for year %v with %v bytes\n", y, len(body))

}
