all: compile docker push clean

compile:
	CGO_ENABLED=0 go build -o goregtrain

docker:
	sudo docker build --no-cache=true --force-rm=true -t dwhitena/goregtrain:singleult .

push:
	sudo docker push dwhitena/goregtrain:singleult

clean:
	rm goregtrain
