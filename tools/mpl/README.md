# markdown-playground-links

This tool processes markdown files from the training materials to automatically
update "Go Playground" links with the latest code from the linked samples.

It will find markdown like the following and automatically update it:

```
The following link will automatically be updated

[markdown-playground-links](main.go) ([Go Playground](http://play.golang.org/p/X8oPoc-i9L))

The program searches for a link followed by a "([Go Playground](.*))" link.
```

You can use an empty playground link and it will be automatically populated
with the right link.

## Usage

Simply run `mpl` with a list of files you want to process.

```
# Update the links of the constants README.md
mpl topics/constants/README.md

# Using zsh wildcards, update all markdown files under the current directory
mpl **/*.md
```

