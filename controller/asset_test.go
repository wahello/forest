// Copyright 2020 tree xie
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

package controller

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vicanso/elton"
)

const testHTMLContent = `<!DOCTYPE html>
<html>
    <head>
        <title>仅用于测试的html</title>
    </head>
    <body></body>
</html>`

func TestStaticFile(t *testing.T) {
	assert := assert.New(t)
	sf := &staticFile{
		box: assetBox,
	}
	assert.True(sf.Exists("index.html"))
	assert.False(sf.Exists("test.html"))

	buf, err := sf.Get("index.html")
	assert.Nil(err)
	assert.Equal(testHTMLContent, string(buf))

	assert.Nil(sf.Stat("index.html"))

	r, err := sf.NewReader("index.html")
	assert.Nil(err)
	assert.NotNil(r)
	buf, err = ioutil.ReadAll(r)
	assert.Nil(err)
	assert.Equal(testHTMLContent, string(buf))
}

func TestSendFile(t *testing.T) {
	assert := assert.New(t)
	c := elton.NewContext(httptest.NewRecorder(), nil)
	err := sendFile(c, "index.html")
	assert.Nil(err)
	assert.Equal(testHTMLContent, c.BodyBuffer.String())
	assert.Equal("text/html; charset=utf-8", c.GetHeader(elton.HeaderContentType))

	err = sendFile(c, "favicon.png")
	assert.Nil(err)
	assert.Equal(958, c.BodyBuffer.Len())
	assert.Equal("image/png", c.GetHeader(elton.HeaderContentType))
}

func TestGetIndex(t *testing.T) {
	assert := assert.New(t)

	ctrl := assetCtrl{}
	c := elton.NewContext(httptest.NewRecorder(), nil)
	err := ctrl.getIndex(c)
	assert.Nil(err)
	assert.Equal(testHTMLContent, c.BodyBuffer.String())
	assert.Equal("text/html; charset=utf-8", c.GetHeader(elton.HeaderContentType))
	assert.Equal("public, max-age=10", c.GetHeader(elton.HeaderCacheControl))
}

func TestGetFavIcon(t *testing.T) {
	assert := assert.New(t)
	ctrl := assetCtrl{}
	c := elton.NewContext(httptest.NewRecorder(), nil)
	err := ctrl.getFavIcon(c)
	assert.Nil(err)
	assert.Equal(958, c.BodyBuffer.Len())
	assert.Equal("image/png", c.GetHeader(elton.HeaderContentType))
	assert.Equal("public, max-age=3600, s-maxage=600", c.GetHeader(elton.HeaderCacheControl))
}
