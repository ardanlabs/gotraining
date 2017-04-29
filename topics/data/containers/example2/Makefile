all: compile docker push clean

compile:
	CGO_ENABLED=0 go build -o goregpredict

docker:
	sudo docker build --no-cache=true --force-rm=true -t dwhitena/goregpredict:ult .

push:
	sudo docker push dwhitena/goregpredict:ult

clean:
	rm goregpredict
