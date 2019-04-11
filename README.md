# Hyperledger Fabric Demo

**<h1>Introduction</h1>**
INTUZ is presenting a working demo for Property blockchain based on Hyperledger Fabric framework.

**<h1>Features</h1>**
* The demo is based on Hyperledger Fabric framework created to keep the record for the transaction of the Property owners.
* The chaincode includes the following fields e.g.: Owner name, House number, Address field 1, Address field 2, City, Pincode, Area (in Sq. feet), Purchase Date, Purchase Type, Price.
* A user can add and edit the information like property owner, Date of Purchase etc.

**<h1>Terminal commands</h1>**
Please follow the below steps, this will execute the results in terminal.

Copy the chaincode property.go in **fabric-samples/chaincode/chaincode_example02/go** directory.

Go to docker dev mode, open the bash and fire the following command. This will create a docker container and put the container up for the execution.
```
docker-compose -f docker-compose-simple.yaml up -d
```

**<h1>The following commands should need to be executed in same terminal window.</h1>**

To go into Docker chaincode.
```
docker exec -it chaincode bash
```

To build the executable from a chaincode.go file.
P.S.: Please don't execute this command If the executable file is already available at the desired path.
```
go build -o property
```

To set the Peer address port.
```
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./chaincode_example02/go/property
```

**<h1>Open another terminal in parallel and follow the commands.</h1>**

To go into Docker CLI
```
docker exec -it cli bash
```

This will install the chaincode.
```
peer chaincode install -p chaincodedev/chaincode/chaincode_example02/go -n mycc -v 0
```

To instantiate the chaincode.
```
peer chaincode instantiate -n mycc -v 0 -c '{"Args":["init"]}' -C myc
```

To insert multiple Property data into Blockchain, kindly follow the below commands but make sure to change the text instead "pay12" to the actual data in below format.
```
peer chaincode invoke -n mycc -c '{"Args":["addData","1","OwnerName","houseNo","addressField1","addressField2","city","pincode","area","purchaseDate","purchaseType","pay12"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["addData","2","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["addData","3","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["addData","4","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["addData","5","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["addData","6","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["addData","7","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12","pay12"]}' -C myc
```

To check whether the data is inserted successfully or not.
```
peer chaincode query -C myc -n mycc -c '{"Args":["readAllData"]}'
```

To fetch the data for single record using ID
```
peer chaincode query -C myc -n mycc -c '{"Args":["readData","1"]}'
```

To update data in the chaincode
```
peer chaincode invoke -C myc -n mycc -c '{"Args":["UpdateData","2","abc"]}'
```

To delete any specific chaincode data
```
peer chaincode invoke -C myc -n mycc -c '{"Args":["deleteData","1"]}'
```

----------------------------------

**<h1>License</h1>**
The MIT License (MIT)
<br/>
Copyright (c) 2019 INTUZ
<br/>
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions: 
<br/>
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
<h1></h1>
<a href="https://www.intuz.com/" target="_blank"><img src="https://d32qh7kc7vdj86.cloudfront.net/2017/logo-z.png"></a>
