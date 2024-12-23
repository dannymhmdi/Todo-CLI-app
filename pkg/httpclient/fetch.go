package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func FetchData(url string) (*[]ApiResponse,error) {
	resp ,err := http.Get(url)
	if err !=nil {
		return nil , fmt.Errorf("faild to get data %+v",err)
	}
	defer resp.Body.Close()
	body , err := io.ReadAll(resp.Body)
	if err !=nil {
		return nil , fmt.Errorf("faild in reading body %+v",err)
	}
	apiResponse := []ApiResponse{}
	if err := json.Unmarshal(body,&apiResponse) ; err !=nil {
		return nil , fmt.Errorf("faild to parse json data %+v",err)
	}
	return &apiResponse, nil
}




type ApiResponse struct{
	Id int
	Name, Username , Email,Phone,Website string
	Address Address
    Company Company
	}
	
	type Address struct {
		Street , Suite , City,Zipcode string
		Geo Geo
	}
	
	type Geo struct {
	Lat , Lng string
	}

	type Company struct {
		Name , CatchPhrase , Bs string
	}