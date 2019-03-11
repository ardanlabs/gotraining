# https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app

.EXPORT_ALL_VARIABLES:
PROJECT_NAME=project
PROJECT_ID=got-project-d
ACCOUNT_ID=xxxxxxxxxxxxxxxxx

#=======================================================================================================================
# Local

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

up:
	export PROJECT_ID=$(PROJECT_ID)
	docker-compose up
	@echo ======================================================================

down:
	docker-compose down
	@echo ======================================================================

#=======================================================================================================================
# Cloud

config:
	@echo Setting environment for $(PROJECT_ID)
	gcloud config set project $(PROJECT_ID)
	gcloud config set compute/zone us-east1-b
	@echo ======================================================================

project:
	gcloud projects create $(PROJECT_ID)
	gcloud beta billing projects link $(PROJECT_ID) --billing-account=$(ACCOUNT_ID)
	gcloud services enable container.googleapis.com
	@echo ======================================================================

cluster:
	gcloud container clusters create $(PROJECT_ID)-cluster
	gcloud compute instances list
	@echo ======================================================================

upload:
	gcloud auth configure-docker
	docker push gcr.io/$(PROJECT_ID)/$(PROJECT_ID)-amd64:1.0
	@echo ======================================================================

kompose:
	rm *deployment.yaml
	rm *service.yaml
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

dstatus:
	kubectl get pods
	@echo ======================================================================

expose:
	kubectl expose -f $(PROJECT_ID)-service.yaml --type=LoadBalancer
	kubectl get services $(PROJECT_ID)
	@echo ======================================================================

estatus:
	kubectl get services $(PROJECT_ID)
	@echo ======================================================================

delete:
	kubectl delete services $(PROJECT_ID)
	kubectl delete deployment $(PROJECT_ID)	
	gcloud container clusters delete $(PROJECT_ID)-cluster
	gcloud projects delete $(PROJECT_ID)
	docker image remove gcr.io/$(PROJECT_ID)/$(PROJECT_ID)-amd64:1.0
	@echo ======================================================================