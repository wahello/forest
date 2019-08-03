// Copyright 2019 tree xie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"fmt"
	"net/http"

	"github.com/vicanso/cod"
	"github.com/vicanso/forest/service"
	"github.com/vicanso/hes"
)

var (
	errIPNotAllow = &hes.Error{
		StatusCode: http.StatusBadRequest,
		Message:    "request is forbidden",
		Category:   "IB",
	}
)

// NewIPBlock create a new block ip middleware
func NewIPBlock() cod.Handler {
	return func(c *cod.Context) (err error) {
		fmt.Println(c.RealIP())
		if service.IsBlockIP(c.RealIP()) {
			err = errIPNotAllow
			return
		}
		return c.Next()
	}
}
