// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `app.js`,
		FileModTime: time.Unix(1477170545, 0),
		Content:     string("h1 = document.getElementsByTagName(\"h1\")[0];\nh1.innerHTML = \"Hello World\";\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `index.html`,
		FileModTime: time.Unix(1477170545, 0),
		Content:     string("<html>\n<head>\n  <meta charset=\"utf-8\">\n  <title>Ultimate Web</title>\n  <link rel=\"stylesheet\" href=\"/assets/styles.css\" type=\"text/css\" media=\"all\" />\n</head>\n<body>\n  <h1></h1>\n  <script src=\"/assets/app.js\" type=\"text/javascript\" charset=\"utf-8\"></script>\n</body>\n</html>\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    `styles.css`,
		FileModTime: time.Unix(1477170545, 0),
		Content:     string("h1 {\n  font-weight: bold;\n  color: blue;\n  font-size: 48px;\n}\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1477170545, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // app.js
			file3, // index.html
			file4, // styles.css

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`./static`, &embedded.EmbeddedBox{
		Name: `./static`,
		Time: time.Unix(1477170545, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"app.js":     file2,
			"index.html": file3,
			"styles.css": file4,
		},
	})
}
