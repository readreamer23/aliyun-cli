package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateUploadPathForSQLServer invokes the rds.CreateUploadPathForSQLServer API synchronously
// api document: https://help.aliyun.com/api/rds/createuploadpathforsqlserver.html
func (client *Client) CreateUploadPathForSQLServer(request *CreateUploadPathForSQLServerRequest) (response *CreateUploadPathForSQLServerResponse, err error) {
	response = CreateCreateUploadPathForSQLServerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUploadPathForSQLServerWithChan invokes the rds.CreateUploadPathForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/createuploadpathforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUploadPathForSQLServerWithChan(request *CreateUploadPathForSQLServerRequest) (<-chan *CreateUploadPathForSQLServerResponse, <-chan error) {
	responseChan := make(chan *CreateUploadPathForSQLServerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUploadPathForSQLServer(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateUploadPathForSQLServerWithCallback invokes the rds.CreateUploadPathForSQLServer API asynchronously
// api document: https://help.aliyun.com/api/rds/createuploadpathforsqlserver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUploadPathForSQLServerWithCallback(request *CreateUploadPathForSQLServerRequest, callback func(response *CreateUploadPathForSQLServerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUploadPathForSQLServerResponse
		var err error
		defer close(result)
		response, err = client.CreateUploadPathForSQLServer(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateUploadPathForSQLServerRequest is the request struct for api CreateUploadPathForSQLServer
type CreateUploadPathForSQLServerRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	DBName               string           `position:"Query" name:"DBName"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// CreateUploadPathForSQLServerResponse is the response struct for api CreateUploadPathForSQLServer
type CreateUploadPathForSQLServerResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	InternetFtpServer string `json:"InternetFtpServer" xml:"InternetFtpServer"`
	InternetPort      int    `json:"InternetPort" xml:"InternetPort"`
	IntranetFtpserver string `json:"IntranetFtpserver" xml:"IntranetFtpserver"`
	Intranetport      int    `json:"Intranetport" xml:"Intranetport"`
	UserName          string `json:"UserName" xml:"UserName"`
	Password          string `json:"Password" xml:"Password"`
	FileName          string `json:"FileName" xml:"FileName"`
}

// CreateCreateUploadPathForSQLServerRequest creates a request to invoke CreateUploadPathForSQLServer API
func CreateCreateUploadPathForSQLServerRequest() (request *CreateUploadPathForSQLServerRequest) {
	request = &CreateUploadPathForSQLServerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateUploadPathForSQLServer", "rds", "openAPI")
	return
}

// CreateCreateUploadPathForSQLServerResponse creates a response to parse from CreateUploadPathForSQLServer response
func CreateCreateUploadPathForSQLServerResponse() (response *CreateUploadPathForSQLServerResponse) {
	response = &CreateUploadPathForSQLServerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}