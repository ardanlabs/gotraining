tag = app/httpd

all:
	$(error please pick a target)

build:
	docker build -t $(tag) .

run:
	docker run -p 8080:8080 $(tag)

