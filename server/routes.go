///////////////////////////////////////////////////////////////////////////////
//
// The MIT License (MIT)
// Copyright (c) 2018 Jivan Amara
// Copyright (c) 2018 Tom Kralidis
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.
//
///////////////////////////////////////////////////////////////////////////////

package server

// Provides mappings for URL routes to handler functions.

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func setUpRoutes() http.Handler {
	r := httprouter.New()

	r.HandlerFunc("GET", "/", root)
	r.HandlerFunc("GET", "/conformance", conformance)
	r.HandlerFunc("GET", "/api", openapi)

	r.HandlerFunc("GET", "/collections", collectionsMetaData)
	r.HandlerFunc("GET", "/collections/:name", collectionMetaData)
	r.HandlerFunc("GET", "/collections/:name/items", collectionData)
	r.HandlerFunc("GET", "/collections/:name/items/:feature_id", collectionData)

	return r
}
