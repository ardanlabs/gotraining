# https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

PROJECT_ID=$(PROJECT)-b

init:
	@echo $(PROJECT_ID)
	gcloud config set project $(PROJECT_ID)
	gcloud config set compute/zone us-east1-b

project: init	
	gcloud projects create $(PROJECT_ID) --enable-cloud-apis
	@echo You must enable the GKE API and set billing: https://console.cloud.google.com/apis/library/container.googleapis.com?q=Kubernetes%20Engine&project=$(PROJECT_ID)&folder&organizationId

cluster: init
	gcloud container clusters create $(PROJECT_ID)-cluster
	gcloud compute instances list

build: init
	docker build \
		-t gcr.io/$(PROJECT_ID)/$(PROJECT_ID)-amd64:1.0 \
		--build-arg TITLE=$(PROJECT_ID) \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.
	docker system prune -f
	docker images

upload: init
	# gcloud container clusters get-credentials $(PROJECT)-cluster
	gcloud auth configure-docker
	docker push gcr.io/$(PROJECT_ID)/$(PROJECT_ID)-amd64:1.0

kompose:
	export PROJECT_ID=$(PROJECT_ID)
	kompose convert

deploy:
	kubectl create -f $(PROJECT_ID)-deployment.yaml
	kubectl get pods

expose:
	kubectl expose -f $(PROJECT_ID)-service.yaml --type=LoadBalancer
	kubectl get services $(PROJECT_ID)
	kubectl describe services $(PROJECT_ID)

up:
	export PROJECT_ID=$(PROJECT_ID)
	docker-compose up

down:
	docker-compose down

delete:
	kubectl delete services $(PROJECT_ID)
	kubectl delete deployment $(PROJECT_ID)
	gcloud container clusters delete $(PROJECT_ID)-cluster
	gcloud projects delete $(PROJECT_ID)


