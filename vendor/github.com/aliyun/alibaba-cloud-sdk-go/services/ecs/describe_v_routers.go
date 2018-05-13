package ecs

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

// DescribeVRouters invokes the ecs.DescribeVRouters API synchronously
// api document: https://help.aliyun.com/api/ecs/describevrouters.html
func (client *Client) DescribeVRouters(request *DescribeVRoutersRequest) (response *DescribeVRoutersResponse, err error) {
	response = CreateDescribeVRoutersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVRoutersWithChan invokes the ecs.DescribeVRouters API asynchronously
// api document: https://help.aliyun.com/api/ecs/describevrouters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVRoutersWithChan(request *DescribeVRoutersRequest) (<-chan *DescribeVRoutersResponse, <-chan error) {
	responseChan := make(chan *DescribeVRoutersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVRouters(request)
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

// DescribeVRoutersWithCallback invokes the ecs.DescribeVRouters API asynchronously
// api document: https://help.aliyun.com/api/ecs/describevrouters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVRoutersWithCallback(request *DescribeVRoutersRequest, callback func(response *DescribeVRoutersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVRoutersResponse
		var err error
		defer close(result)
		response, err = client.DescribeVRouters(request)
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

// DescribeVRoutersRequest is the request struct for api DescribeVRouters
type DescribeVRoutersRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VRouterId            string           `position:"Query" name:"VRouterId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// DescribeVRoutersResponse is the response struct for api DescribeVRouters
type DescribeVRoutersResponse struct {
	*responses.BaseResponse
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	VRouters   VRouters `json:"VRouters" xml:"VRouters"`
}

// CreateDescribeVRoutersRequest creates a request to invoke DescribeVRouters API
func CreateDescribeVRoutersRequest() (request *DescribeVRoutersRequest) {
	request = &DescribeVRoutersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeVRouters", "ecs", "openAPI")
	return
}

// CreateDescribeVRoutersResponse creates a response to parse from DescribeVRouters response
func CreateDescribeVRoutersResponse() (response *DescribeVRoutersResponse) {
	response = &DescribeVRoutersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}