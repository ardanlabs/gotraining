# https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

PROJECT_NAME=project
PROJECT_ID=got-project-c

config:
	@echo Setting environment for $(PROJECT_ID)
	gcloud config set project $(PROJECT_ID)
	gcloud config set compute/zone us-east1-b
	@echo ======================================================================

project:
	gcloud projects create $(PROJECT_ID)
	gcloud service enable container.googleapis.com
	gcloud beta billing projects link $(PROJECT_ID)
	@echo ======================================================================
	@echo INFO: You must enable the GKE API and set billing
	@echo INFO: https://console.cloud.google.com/apis/library/container.googleapis.com?q=Kubernetes%20Engine&project=$(PROJECT_ID)&folder&organizationId

cluster:
	gcloud container clusters create $(PROJECT_ID)-cluster
	gcloud compute instances list
	@echo ======================================================================

build:
	docker build \
		-t gcr.io/$(PROJECT_ID)/$(PROJECT_ID)-amd64:1.0 \
		--build-arg TITLE=$(PROJECT_ID) \
		--build-arg VCS_REF=`git rev-parse HEAD` \
		--build-arg BUILD_DATE=`date -u +”%Y-%m-%dT%H:%M:%SZ”` \
		.
	docker system prune -f
	docker images
	@echo ======================================================================

upload:
	gcloud auth configure-docker
	docker push gcr.io/$(PROJECT_ID)/$(PROJECT_ID)-amd64:1.0
	@echo ======================================================================

kompose:
	kompose convert
	mv $(PROJECT_NAME)-deployment.yaml $(PROJECT_ID)-deployment.yaml
	mv $(PROJECT_NAME)-service.yaml $(PROJECT_ID)-service.yaml
	sed -i '' 's/: $(PROJECT_NAME)/: $(PROJECT_ID)/g' $(PROJECT_ID)-deployment.yaml
	sed -i '' 's/: $(PROJECT_NAME)/: $(PROJECT_ID)/g' $(PROJECT_ID)-service.yaml
	@echo ======================================================================

deploy:
	kubectl create -f $(PROJECT_ID)-deployment.yaml
	kubectl get pods
	@echo ======================================================================

getpods:
	kubectl get pods
	@echo ======================================================================

expose:
	kubectl expose -f $(PROJECT_ID)-service.yaml --type=LoadBalancer
	kubectl get services $(PROJECT_ID)
	kubectl describe services $(PROJECT_ID)
	@echo ======================================================================

getservices:
	kubectl get services $(PROJECT_ID)
	kubectl describe services $(PROJECT_ID)
	@echo ======================================================================

up:
	docker-compose up
	@echo ======================================================================

down:
	docker-compose down
	@echo ======================================================================

delete:
	kubectl delete services $(PROJECT_ID)
	kubectl delete deployment $(PROJECT_ID)	
	gcloud container clusters delete $(PROJECT_ID)-cluster
	gcloud projects delete $(PROJECT_ID)
	@echo ======================================================================


