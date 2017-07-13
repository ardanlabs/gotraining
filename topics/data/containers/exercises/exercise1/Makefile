all: compile docker push clean

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goregtrain

docker:
	sudo docker build --no-cache=true --force-rm=true -t dwhitena/goregtrain:multiult .

push:
	sudo docker push dwhitena/goregtrain:multiult

clean:
	rm goregtrain
