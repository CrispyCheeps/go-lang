package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var baseURL = "http://localhost:8080"

type student struct {
	ID string
	Name string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

//A.55.2. HTTP Request Dengan Form Data
func fetchUser(ID string) (student, error) {
	// var err error
	// var client = &http.Client{}
	// var data student

	// var param = url.Values{}
	// param.Set("id", ID)
	// //ini ngubah dr data param mjd btk bytes.Buffer.
	// var payload = bytes.NewBufferString(param.Encode())

	// request, err := http.NewRequest("POST", baseURL+"/user", payload)
	// if err != nil {
	// 	return data, err
	// }
	// request.Header.Set("Content-Type", "applicatioin/x-www-form-urlencoded")

	// response, err := client.Do(request)
	// if err != nil {
	// 	return data, err
	// }
	// defer response.Body.Close()

	// err = json.NewDecoder(response.Body).Decode(&data)
	// if err != nil {
	// 	return data, err
	// }

	// return data, nil
	
	var err error
    var client = &http.Client{}
    var data student

    request, err := http.NewRequest("GET", baseURL+"/user?id="+ID, nil)
    if err != nil {
        return data, err
    }

    response, err := client.Do(request)
    if err != nil {
        return data, err
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        return data, err
    }

    return data, nil
}

func main () {
	// var users, err = fetchUsers()
	// if err != nil {
	// 	fmt.Println("Error!", err.Error())
	// 	return
	// }

	// for _, each := range users {
	// 	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	// }

	//fetchUser
	var user1, err2 = fetchUser("E001")
	if err2 != nil {
		fmt.Println("Error!", err2.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade:%d\n", user1.ID, user1.Name, user1.Grade)
}