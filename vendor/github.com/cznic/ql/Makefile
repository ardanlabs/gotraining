# Copyright (c) 2014 The ql Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean cover cpu editor internalError later mem nuke todo edit

grep=--include=*.go --include=*.l --include=*.y --include=*.yy --exclude=ql.y
ngrep='TODOOK\|parser\.go\|scanner\.go\|.*_string\.go'

all: editor scanner.go parser.go
	go vet 2>&1 | grep -v $(ngrep) || true
	golint 2>&1 | grep -v $(ngrep) || true
	make todo
	unused . || true
	misspell *.go
	gosimple || true
	go install ./...

bench: all
	go test -run NONE -bench .

clean:
	go clean
	rm -f *~ y.go y.tab.c *.out *.test

coerce.go: helper/helper.go
	if [ -f coerce.go ] ; then rm coerce.go ; fi
	go run helper/helper.go | gofmt > $@

cover:
	t=$(shell tempfile) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu: clean
	go test -run @ -bench . -cpuprofile cpu.out
	go tool pprof -lines *.test cpu.out

edit:
	@ 1>/dev/null 2>/dev/null gvim -p Makefile *.l *.y *.go testdata.ql testdata.log

editor: ql.y scanner.go parser.go coerce.go
	gofmt -s -l -w *.go
	go test -i
	go test 2>&1 | tee log

internalError:
	egrep -ho '"internal error.*"' *.go | sort | cat -n

later:
	@grep -n $(grep) LATER * || true
	@grep -n $(grep) MAYBE * || true

mem: clean
	go test -run @ -bench . -memprofile mem.out -memprofilerate 1 -timeout 24h
	go tool pprof -lines -web -alloc_space *.test mem.out

nuke: clean
	go clean -i

parser.go: parser.y
	a=$(shell tempfile) ; \
	  goyacc -o /dev/null -xegen $$a $< ; \
	  goyacc -cr -o $@ -xe $$a $< ; \
	  rm -f $$a
	sed -i -e 's|//line.*||' -e 's/yyEofCode/yyEOFCode/' $@

ql.y: doc.go
	sed -n '1,/^package/ s/^\/\/  //p' < $< \
		| ebnf2y -o $@ -oe $*.ebnf -start StatementList -pkg $* -p _
	goyacc -cr -o /dev/null $@

scanner.go: scanner.l parser.go
	golex -o $@ $<

todo:
	@grep -nr $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * || true
	@grep -nr $(grep) TODO * || true
	@grep -nr $(grep) BUG * || true
	@grep -nr $(grep) [^[:alpha:]]println * || true
