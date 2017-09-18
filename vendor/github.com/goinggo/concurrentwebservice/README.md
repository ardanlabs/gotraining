# Concurrent Web Service

Copyright 2014 Ardan Studios. All rights reserved.  
Use of this source code is governed by a BSD-style license that can be found in the LICENSE handle.

This application provides a sample to the html package to create a simple search engine. The engine supports Google, Bing and Blekko searches. You can request results for all three engines or ask for just the first result. Searches are performed concurrently. Use the GOMAXPROCS environment variables to run the searches in parallel.

Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanstudios.com

### Installation

	-- Get, build and install the code
	go get github.com/goinggo/concurrentwebservice
		
	-- Run
	cd $GOPATH/src/github.com/goinggo/concurrentwebservice
	go build
	./concurrentwebservice
	http://localhost:9999/search

### Notes About Architecture

The code is built within two packages. The service package handles the processing of HTTP requests and responses. HTML templates are used to render the views. The search package handles the processing of searches agains the different search engines. An interface called Searcher is declared to support the implementation of new Searchers.