## Search Engine Application

To see Go in action we are going to build a complete Go program. The program implements functionality that can be found in many Go programs being developed today. The program provides a sample to the html package to create a simple search engine. The engine supports Google, Bing and Blekko searches. You can request results for all three engines or ask for just the first result. Searches are performed concurrently. Use the GOMAXPROCS environment variables to run the searches in parallel.

## Program Architecture

![Image of App.](client_image.png)

The program is broken into several distinct steps that run across different goroutines. We will explore the code as it flows from the main goroutine into the searching goroutines and then back to the main goroutine. To start, let's review the structure of the project:

*cd $GOPATH/src/github.com/ArdanStudios/web_app/sample*

* **sample**
	* **search**
		* [bing.go](sample/search/bing.go) -- Performs searches against Bing
		* [blekko.go](sample/search/blekko.go) -- Performs searches against Blekko
		* [google.go](sample/search/google.go)-- Performs searches against Google
		* [rss.go](sample/search/rss.go)-- Boilerplate code for handling RSS feeds
		* [search.go](sample/search/search.go)-- Searching framework code
	* **service**
		* [index.go](sample/service/index.go)-- Handles the rendering of the index page
		* [service.go](sample/service/service.go)-- Initializes and runs the web app
		* [templates.go](sample/service/templates.go)-- Support for handling html templates
	* **static/css**
		* [main.css](sample/static/css/main.css)-- Stylesheet for web app
	* **tests**
		* [endpoint_test.go](sample/tests/endpoint_test.go)-- Tests for endpoint testing
	* **views**
		* [basic-layout.html](sample/views/basic-layout.html)-- Layout HTML for the web app
		* [index.html](sample/views/index.html)-- HTML for the index page
		* [results.html](sample/views/resuls.html)-- HTML for rendering the search results
	* [main.go](sample/main.go) -- Programs entry point

The code is organized within two packages. The service package handles the processing of HTTP requests and responses. HTML templates are used to render the views. The search package handles the processing of searches agains the different search engines. An interface called Searcher is declared to support the implementation of new Searchers.

___
[![GoingGo Training](../00-slides/images/ggt_logo.png)](http://www.goinggotraining.net)
[![Ardan Studios](../00-slides/images/ardan_logo.png)](http://www.ardanstudios.com)
[![GoingGo Blog](../00-slides/images/ggb_logo.png)](http://www.goinggo.net)
___
All material is licensed under the [GNU Free Documentation License](https://github.com/ArdanStudios/gotraining/blob/master/LICENSE).