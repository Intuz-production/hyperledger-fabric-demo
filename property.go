//  The MIT License (MIT)

//  Copyright (c) 2018 Intuz Solutions Pvt Ltd.

//  Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files
//  (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify,
//  merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:

//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
//  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
//  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
//  CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

func main() {
	fmt.Println("Using library with object")
	err := shim.Start(new(PropertData))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

type PropertData struct {
}

type Property struct {
	OwnerName     string `json:"OwnerName"`
	HouseNo       string `json:"houseNo"`
	AddressField1 string `json:"addressField1"`
	AddressField2 string `json:"addressField2"`
	City          string `json:"city"`
	Pincode       string `json:"pincode"`
	Area          string `json:"area"`
	PurchaseDate  string `json:"purchaseDate"`
	PurchaseType  string `json:"purchaseType"`
	Price         string `json:"price"`
}

// Init Function Started
func (s *PropertData) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Init Function End

// Invoke Function Start
func (s *PropertData) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	if function == "addData" {
		return s.addData(APIstub, args)
	}

	if function == "readData" {
		return s.readData(APIstub, args)
	}

	if function == "deleteData" {
		return s.deleteData(APIstub, args)
	}

	if function == "readAllData" {
		return s.readAllData(APIstub, args)
	}

	if function == "UpdateData" {
		return s.UpdateData(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// Invoke Function End

func (s *PropertData) addData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 11 {
		return shim.Error("Incorrect number of arguments. Expecting 11")
	}

	var data = Property{OwnerName: args[1], HouseNo: args[2], AddressField1: args[3], AddressField2: args[4], City: args[5], Pincode: args[6], Area: args[7], PurchaseDate: args[8], PurchaseType: args[9], Price: args[10]}

	dataAsBytes, _ := json.Marshal(data)
	APIstub.PutState(args[0], dataAsBytes)

	return shim.Success(nil)
}
func (s *PropertData) readData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting id of data to query")
	}

	name = args[0]
	valAsbytes, err := APIstub.GetState(name)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Property does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}
func (s *PropertData) deleteData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	var jsonResp string
	var dataJSON Property
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	dataID := args[0]

	// to maintain the Inv~PI index, we need to read the data first and get its data
	valAsbytes, err := APIstub.GetState(dataID) 
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get stateinvoke for " + dataID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Property does not exist: " + dataID + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &dataJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + dataID + "\"}"
		return shim.Error(jsonResp)
	}

	err = APIstub.DelState(dataID) //remove data from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	// maintain the index
	indexName := "Inv~PI"
	DataIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{dataJSON.OwnerName, dataJSON.HouseNo})
	if err != nil {
		return shim.Error(err.Error())
	}

	// Delete index entry to state.
	err = APIstub.DelState(DataIndexKey)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

func (s *PropertData) readAllData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	startKey := "1"
	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllData:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
func (s *PropertData) UpdateData(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	dataAsBytes, _ := APIstub.GetState(args[0])
	if dataAsBytes == nil {
		return shim.Error("Data not recieved")
	}
	data := Property{}

	json.Unmarshal(dataAsBytes, &data)

	data.OwnerName = args[1]

	dataAsBytes, _ = json.Marshal(data)
	err := APIstub.PutState(args[0], dataAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change Inv: %s", args[0]))
	}

	return shim.Success(nil)
}
func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return &buffer, nil
}