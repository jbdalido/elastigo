// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http:www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elastigo

import (
	"encoding/json"
	"fmt"
)

type JsonAliases struct {
	Actions []JsonAliasAction `json:"actions,omitempty"`
}

type JsonAliasAction struct {
	Remove *JsonAlias `json:"remove,omitempty"`
	Add    *JsonAlias `json:"add,omitempty"`
}

type JsonAlias struct {
	Index string `json:"index,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type GetAliasesResponse map[string]interface{}

// The API allows you to create an index alias through an API.
func (c *Conn) AddAlias(index string, alias string) (BaseResponse, error) {
	var url string
	var retval BaseResponse

	if len(index) > 0 {
		url = "/_aliases"
	} else {
		return retval, fmt.Errorf("You must specify an index to create the alias on")
	}

	jsonAliases := JsonAliases{
		Actions: []JsonAliasAction{
			{
				Add: &JsonAlias{
					Alias: alias,
					Index: index,
				},
			},
		},
	}

	requestBody, err := json.Marshal(jsonAliases)
	if err != nil {
		return retval, err
	}

	body, err := c.DoCommand("POST", url, nil, requestBody)
	if err != nil {
		return retval, err
	}

	jsonErr := json.Unmarshal(body, &retval)
	if jsonErr != nil {
		return retval, jsonErr
	}

	return retval, err
}

// The API allows you to remova an index alias through an API.
func (c *Conn) RemoveAlias(index string, alias string) (BaseResponse, error) {
	var url string
	var retval BaseResponse

	if len(index) > 0 {
		url = "/_aliases"
	} else {
		return retval, fmt.Errorf("You must specify an index to create the alias on")
	}

	jsonAliases := JsonAliases{
		Actions: []JsonAliasAction{
			{
				Remove: &JsonAlias{
					Alias: alias,
					Index: index,
				},
			},
		},
	}

	requestBody, err := json.Marshal(jsonAliases)
	if err != nil {
		return retval, err
	}

	body, err := c.DoCommand("POST", url, nil, requestBody)
	if err != nil {
		return retval, err
	}

	jsonErr := json.Unmarshal(body, &retval)
	if jsonErr != nil {
		return retval, jsonErr
	}

	return retval, err
}

func (c *Conn) GetIndexFromAlias(alias string) (GetAliasesResponse, error) {
	var retval GetAliasesResponse
	var url string
	if alias != "" {
		url = fmt.Sprintf("/_alias/%s", alias)

	} else {
		url = "/_alias"
	}
	body, err := c.DoCommand("GET", url, nil, nil)
	if err != nil {
		return retval, err
	}
	if err == nil {
		// marshall into json
		jsonErr := json.Unmarshal(body, &retval)
		if jsonErr != nil {
			return retval, jsonErr
		}
	}
	return retval, err
}
