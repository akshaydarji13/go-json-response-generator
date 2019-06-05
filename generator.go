package generator

import (
	"encoding/json"
	"fmt"
	"log"
)

// Response is the struct for response generator data
type response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (res response) jsonMarsheler() ([]byte, error) {
	type t response
	temp := t(res)

	t1, err := json.Marshal(temp)
	if err != nil {
		return nil, err
	}

	fmt.Println("t1: ", string(t1))

	t2, err := json.Marshal(temp.Data)
	if err != nil {
		return nil, err
	}

	fmt.Println("t2: ", string(t2))

	s1 := string(t1[:len(t1)-1])

	fmt.Println("s1: ", s1)
	s2 := string(t2[1:])

	return []byte(s1 + ", " + s2), nil
}

// ResponseGenerator generates json string for a response
func ResponseGenerator(status bool, message string, data ...interface{}) ([]byte, error) {
	fmt.Println("data is: ", data)
	newResponse := response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	// generatedResponse, err := newResponse.jsonMarsheler()
	generatedResponse, err := json.Marshal(newResponse)
	if err != nil {
		log.Fatal("response generator json conversion error: ", err)
		return nil, err
	}

	fmt.Println("generated response: ", string(generatedResponse))
	return generatedResponse, nil
}
