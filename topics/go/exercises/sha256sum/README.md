# sha256sum exercise

Write a program that is given a list of file names as arguments then prints the
sha256 sum for the contents of each file. Print the hashes as a hex string.

Usage and output may look like this:

```
~ $ sha256sum index.html main.css app.js
07997eeb5e73beb86961eeed4dba6d67dca05bc9fb2b412fd6872999d8cc9dbc  index.html
e50857aa7425df74bd7c9e50d4ece17d85c50292104d28a7f58e0d0a4abe4727  main.css
e122c7ebc7b7a8c4960cc4cf1d4d62aaabfa7c88fc90f15926f9febaaeef8345  app.js
```

Topics covered by this exercise:

- Using the `crypto/sha256` package to calculate hashes.
- Interpreting command line arguments and opening files with the `os` package.
- Formatting a `[]byte` as a hex string using `encoding/hex` or `fmt.Printf`
  with the `%x` verb.

A [template file](template/main.go) is included to get you started.
