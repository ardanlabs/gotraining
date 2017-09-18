#### VARIABLES
# RUNARGS: arguments for run
# DOCKER_OPTS: docker-compose options for run, test, launch-*
# TESTPKGS: packages for test, default ./src/...
# TESTFLAGS: flags for test
# VENDOR_ALL: do not ignore some vendors when updating vendor directory
# VENDOR_IGNORE_DIRS: ignore vendor dirs
# KUBECTLFLAGS: flags for kubectl
####

ifndef TESTPKGS
	TESTPKGS = ./src/...
endif
ifdef VENDOR_ALL
	VENDOR_IGNORE_DIRS =
endif

COMPILE_RUN_ARGS = -d -v /var/run/docker.sock:/var/run/docker.sock --privileged=true
VERSION_ADDITIONAL = $(shell git log --pretty=format:%H | head -n 1)
LD_FLAGS = -X github.com/pachyderm/pachyderm/src/server/vendor/github.com/pachyderm/pachyderm/src/client/version.AdditionalVersion=$(VERSION_ADDITIONAL)

CLUSTER_NAME?=pachyderm
CLUSTER_MACHINE_TYPE?=n1-standard-4
CLUSTER_SIZE?=4

BENCH_CLOUD_PROVIDER=aws

ifdef TRAVIS_BUILD_NUMBER
	# Upper bound for travis test timeout
	TIMEOUT = 3600s
else
ifndef TIMEOUT
	# You should be able to specify your own timeout, but by default we'll use the same bound as travis
	TIMEOUT = 1000s
endif
endif

echo-timeout:
	echo $(TIMEOUT)

all: build

version:
	@echo 'package main; import "github.com/pachyderm/pachyderm/src/client/version"; func main() { println(version.PrettyPrintVersion(version.Version)) }' > /tmp/pachyderm_version.go
	go run /tmp/pachyderm_version.go

deps:
	GO15VENDOREXPERIMENT=0 go get -d -v ./src/... ./.

update-deps:
	GO15VENDOREXPERIMENT=0 go get -d -v -u -f ./src/... ./.

test-deps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t ./src/... ./.

update-test-deps:
	GO15VENDOREXPERIMENT=0 go get -d -v -t -u -f ./src/... ./.

build-clean-vendored-client:
	rm -rf src/server/vendor/github.com/pachyderm/pachyderm/src/client

build:
	GO15VENDOREXPERIMENT=1 go build $$(go list ./src/client/... | grep -v '/src/client$$')
	GO15VENDOREXPERIMENT=1 go build $$(go list ./src/server/... | grep -v '/src/server/vendor/' | grep -v '/src/server$$')

install:
	# GOPATH/bin must be on your PATH to access these binaries:
	GO15VENDOREXPERIMENT=1 go install -ldflags "$(LD_FLAGS)" ./src/server/cmd/pachctl

install-doc:
	GO15VENDOREXPERIMENT=1 go install ./src/server/cmd/pachctl-doc

check-docker-version:
	# The latest docker client requires server api version >= 1.24.
	# However, minikube uses 1.23, so if you're connected to minikube, releases
	# may break
	@ \
		docker_major="$$(docker version -f "{{.Server.APIVersion}}" | cut -d. -f1)"; \
		docker_minor="$$(docker version -f "{{.Server.APIVersion}}" | cut -d. -f2)"; \
		echo "docker version = $${docker_major}.$${docker_minor}, need at least 1.24"; \
		test \( "$${docker_major}" -gt 1 \) -o \( "$${docker_minor}" -ge 24 \)

point-release:
	@make VERSION_ADDITIONAL= release-custom

# Run via 'make VERSION_ADDITIONAL=RC release-custom' to specify a version string
release-custom: check-docker-version release-version release-pachd release-worker release-pachctl doc-custom
	@rm VERSION
	@echo "Release completed"

release-version:
	@# Need to blow away pachctl binary if its already there
	@rm $(GOPATH)/bin/pachctl || true
	@make install
	@./etc/build/release_version

release-pachd:
	@VERSION="$(shell cat VERSION)" ./etc/build/release_pachd

release-worker:
	@VERSION="$(shell cat VERSION)" ./etc/build/release_worker

release-pachctl:
	@VERSION="$(shell cat VERSION)" ./etc/build/release_pachctl

docker-build-compile:
	docker build -t pachyderm_compile .

docker-clean-worker:
	docker stop worker_compile || true
	docker rm worker_compile || true

docker-build-worker: docker-clean-worker docker-build-compile
	docker run --name worker_compile $(COMPILE_RUN_ARGS) pachyderm_compile sh etc/compile/compile.sh worker "$(LD_FLAGS)"

docker-wait-worker:
	etc/compile/wait.sh worker_compile

docker-clean-pachd:
	docker stop pachd_compile || true
	docker rm pachd_compile || true

docker-build-pachd: docker-clean-pachd docker-build-compile
	docker run --name pachd_compile $(COMPILE_RUN_ARGS) pachyderm_compile sh etc/compile/compile.sh pachd "$(LD_FLAGS)"

docker-clean-test:
	docker stop test_compile || true
	docker rm test_compile || true

docker-build-test: docker-clean-test docker-build-compile
	docker run --name test_compile $(COMPILE_RUN_ARGS) pachyderm_compile sh etc/compile/compile_test.sh
	etc/compile/wait.sh test_compile
	docker tag pachyderm_test:latest pachyderm/test:`git rev-list HEAD --max-count=1`

docker-push-test:
	docker push pachyderm/test:`git rev-list HEAD --max-count=1`

docker-build-microsoft-vhd:
	docker build -t microsoft_vhd etc/microsoft/create-blank-vhd

docker-wait-pachd:
	etc/compile/wait.sh pachd_compile

docker-build: docker-build-worker docker-build-pachd docker-wait-worker docker-wait-pachd

docker-build-proto:
	docker build -t pachyderm_proto etc/proto

docker-build-netcat:
	docker build -t pachyderm_netcat etc/netcat

docker-build-gpu:
	docker build -t pachyderm_nvidia_driver_install etc/deploy/gpu
	docker tag pachyderm_nvidia_driver_install pachyderm/nvidia_driver_install

docker-push-gpu:
	docker push pachyderm/nvidia_driver_install

docker-push-gpu-dev:
	docker tag pachyderm/nvidia_driver_install pachyderm/nvidia_driver_install:`git rev-list HEAD --max-count=1`
	docker push pachyderm/nvidia_driver_install:`git rev-list HEAD --max-count=1`
	echo pushed pachyderm/nvidia_driver_install:`git rev-list HEAD --max-count=1`

docker-gpu: docker-build-gpu docker-push-gpu

docker-gpu-dev: docker-build-gpu docker-push-gpu-dev

check-kubectl:
	# check that kubectl is installed
	which kubectl

check-kubectl-connection:
	kubectl $(KUBECTLFLAGS) get all > /dev/null

launch-dev-bench: docker-build docker-build-test install
	@# Put it here so sudo can see it
	rm /usr/local/bin/pachctl || true
	ln -s $(GOPATH)/bin/pachctl /usr/local/bin/pachctl
	make launch-bench

build-bench-images: docker-build docker-build-test

push-bench-images: install-bench tag-images push-images
	# We need the pachyderm_compile image to be up to date
	docker tag pachyderm_test pachyderm/bench:`git rev-list HEAD --max-count=1`
	docker push pachyderm/bench:`git rev-list HEAD --max-count=1`

tag-images: install
	docker tag pachyderm_pachd pachyderm/pachd:`$(GOPATH)/bin/pachctl version 2>/dev/null | grep pachctl | awk -v N=2 '{print $$N}'`
	docker tag pachyderm_worker pachyderm/worker:`$(GOPATH)/bin/pachctl version 2>/dev/null | grep pachctl | awk -v N=2 '{print $$N}'`

push-images: tag-images
	docker push pachyderm/pachd:`$(GOPATH)/bin/pachctl version 2>/dev/null | grep pachctl | awk -v N=2 '{print $$N}'`
	docker push pachyderm/worker:`$(GOPATH)/bin/pachctl version 2>/dev/null | grep pachctl | awk -v N=2 '{print $$N}'`

launch-bench:
	@# Make launches each process in its own shell process, so we have to structure
	@# these to run these as one command
	ID=$$( etc/testing/deploy/$(BENCH_CLOUD_PROVIDER).sh --create | tail -n 1); \
	@echo To delete this cluster, run etc/testing/deploy/$(BENCH_CLOUD_PROVIDER).sh --delete=$${ID}; \
	echo etc/testing/deploy/$(BENCH_CLOUD_PROVIDER).sh --delete=$${ID} >./clean_current_bench_cluster.sh; \
	until timeout 10s ./etc/kube/check_ready.sh app=pachd; do sleep 1; done; \
	cat ~/.kube/config;

clean-launch-bench:
	./clean_current_bench_cluster.sh || true

install-bench: install
	@# Since bench is run as sudo, pachctl needs to be under
	@# the secure path
	rm /usr/local/bin/pachctl || true
	[ -f /usr/local/bin/pachctl ] || sudo ln -s $(GOPATH)/bin/pachctl /usr/local/bin/pachctl

launch-dev-test: docker-build-test docker-push-test
	sudo kubectl run bench --image=pachyderm/test:`git rev-list HEAD --max-count=1` \
	    --restart=Never \
	    --attach=true \
	    -- \
	    ./test -test.v

aws-test: tag-images push-images
	ZONE=sa-east-1a etc/testing/deploy/aws.sh --create
	$(MAKE) launch-dev-test
	rm $(HOME)/.pachyderm/config.json
	ZONE=sa-east-1a etc/testing/deploy/aws.sh --delete

run-bench:
	kubectl scale --replicas=4 deploy/pachd
	echo "waiting for pachd to scale up" && sleep 15
	kubectl delete --ignore-not-found po/bench && \
	    kubectl run bench \
			    --image=pachyderm/bench:`git rev-list HEAD --max-count=1` \
					--image-pull-policy=Always \
					--restart=Never \
					--attach=true \
					-- \
					PACH_TEST_CLOUD=true ./test -test.v -test.bench=BenchmarkDaily -test.run=`etc/testing/passing_test_regex.sh`

delete-all-launch-bench:
	etc/testing/deploy/$(BENCH_CLOUD_PROVIDER).sh --delete-all

bench: clean-launch-bench build-bench-images push-bench-images launch-bench run-bench clean-launch-bench

launch-kube: check-kubectl
	etc/kube/start-kube-docker.sh

clean-launch-kube:
	docker kill $$(docker ps -q)

launch: install check-kubectl
	$(eval STARTTIME := $(shell date +%s))
	pachctl deploy local --dry-run | kubectl $(KUBECTLFLAGS) create -f -
	# wait for the pachyderm to come up
	until timeout 1s ./etc/kube/check_ready.sh app=pachd; do sleep 1; done
	@echo "pachd launch took $$(($$(date +%s) - $(STARTTIME))) seconds"

launch-dev: check-kubectl check-kubectl-connection install
	$(eval STARTTIME := $(shell date +%s))
	pachctl deploy local -d --dry-run | kubectl $(KUBECTLFLAGS) create -f -
	# wait for the pachyderm to come up
	until timeout 1s ./etc/kube/check_ready.sh app=pachd; do sleep 1; done
	@echo "pachd launch took $$(($$(date +%s) - $(STARTTIME))) seconds"	

clean-launch: check-kubectl
	pachctl deploy local --dry-run | kubectl $(KUBECTLFLAGS) delete --ignore-not-found -f -

clean-launch-dev: check-kubectl
	pachctl deploy local -d --dry-run | kubectl $(KUBECTLFLAGS) delete --ignore-not-found -f -
	kubectl $(KUBECTLFLAGS) delete rc -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete svc -l suite=pachyderm

full-clean-launch: check-kubectl
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found job -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found all -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found serviceaccount -l suite=pachyderm
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found secret -l suite=pachyderm

launch-test-rethinkdb:
	@# Expose port 8081 so you can connect to the rethink dashboard
	@# (You may need to forward port 8081 if you're running docker machine)
	docker run --name pachyderm-test-rethinkdb -d -p 28015:28015 -p 8081:8080 rethinkdb:2.3.3
	sleep 20  # wait for rethinkdb to start up

clean-launch-test-rethinkdb:
	docker stop pachyderm-test-rethinkdb || true
	docker rm pachyderm-test-rethinkdb || true

clean-pps-storage: check-kubectl
	kubectl $(KUBECTLFLAGS) delete pvc rethink-volume-claim
	kubectl $(KUBECTLFLAGS) delete pv rethink-volume

integration-tests:
	CGOENABLED=0 go test -v ./src/server $(TESTFLAGS) -timeout $(TIMEOUT)

example-tests:
	CGOENABLED=0 go test -v ./src/server/examples $(TESTFLAGS) -timeout $(TIMEOUT)

proto: docker-build-proto
	find src -regex ".*\.proto" \
	| grep -v vendor \
	| xargs tar cf - \
	| docker run -i pachyderm_proto \
	| tar xf -

# Use this to grab a binary for profiling purposes
pachd-profiling-binary: docker-clean-pachd docker-build-compile
	docker run -i  pachyderm_compile sh etc/compile/compile.sh pachd "$(LD_FLAGS)" PROFILE \
	| tar xf -
	# Binary emitted to ./pachd

pretest:
	go get -v github.com/kisielk/errcheck
	rm -rf src/server/vendor
	go vet -n ./src/... | while read line; do \
		modified=$$(echo $$line | sed "s/ [a-z0-9_/]*\.pb\.gw\.go//g"); \
		$$modified; \
		if [ -n "$$($$modified)" ]; then \
		exit 1; \
		fi; \
		done
	git checkout src/server/vendor
	#errcheck $$(go list ./src/... | grep -v src/cmd/ppsd | grep -v src/pfs$$ | grep -v src/pps$$)

#test: pretest test-client clean-launch-test-rethinkdb launch-test-rethinkdb test-fuse test-local docker-build docker-build-netcat clean-launch-dev launch-dev integration-tests example-tests

local-test: docker-build launch-dev test-pfs test-hashtree clean-launch-dev 

test: docker-build clean-launch-dev launch-dev test-pfs test-pps test-hashtree test-auth

test-pfs:
	@# don't run this in verbose mode, as it produces a huge amount of logs
	go test ./src/server/pfs/server -timeout $(TIMEOUT)

test-pps:
	go test -v ./src/server -timeout $(TIMEOUT)

test-hashtree:
	go test ./src/server/pkg/hashtree -timeout $(TIMEOUT)

test-client:
	rm -rf src/client/vendor
	rm -rf src/server/vendor/github.com/pachyderm
	cp -R src/server/vendor src/client/
	GO15VENDOREXPERIMENT=1 go test -cover $$(go list ./src/client/... | grep -v vendor)
	rm -rf src/client/vendor
	git checkout src/server/vendor/github.com/pachyderm

test-fuse:
	CGOENABLED=0 GO15VENDOREXPERIMENT=1 go test -cover $$(go list ./src/server/... | grep -v '/src/server/vendor/' | grep '/src/server/pfs/fuse')

test-local:
	CGOENABLED=0 GO15VENDOREXPERIMENT=1 go test -cover -short $$(go list ./src/server/... | grep -v '/src/server/vendor/' | grep -v '/src/server/pfs/fuse') -timeout $(TIMEOUT)

test-auth:
	go test -v ./src/server/auth/server -timeout $(TIMEOUT)

clean: clean-launch clean-launch-kube

doc-custom: install-doc release-version
	# we rename to pachctl because the program name is used in generating docs
	cp $(GOPATH)/bin/pachctl-doc ./pachctl
	# This file isn't autogenerated so we need to keep it around:
	cp doc/pachctl/pachctl.rst .
	rm -rf doc/pachctl && mkdir doc/pachctl
	./pachctl
	rm ./pachctl
	mv pachctl.rst doc/pachctl
	VERSION="$(shell cat VERSION)" ./etc/build/release_doc

doc:
	@make VERSION_ADDITIONAL= doc-custom

clean-launch-monitoring:
	kubectl delete --ignore-not-found -f ./etc/plugin/monitoring

launch-monitoring:
	kubectl create -f ./etc/plugin/monitoring
	@echo "Waiting for services to spin up ..."
	until timeout 5s ./etc/kube/check_ready.sh k8s-app=heapster kube-system; do sleep 5; done
	until timeout 5s ./etc/kube/check_ready.sh k8s-app=influxdb kube-system; do sleep 5; done
	until timeout 5s ./etc/kube/check_ready.sh k8s-app=grafana kube-system; do sleep 5; done
	@echo "All services up. Now port forwarding grafana to localhost:3000"
	kubectl --namespace=kube-system port-forward `kubectl --namespace=kube-system get pods -l k8s-app=grafana -o json | jq '.items[0].metadata.name' -r` 3000:3000 &

clean-launch-logging: check-kubectl check-kubectl-connection
	git submodule update --init
	cd etc/plugin/logging && ./undeploy.sh

launch-logging: check-kubectl check-kubectl-connection
	@# Creates Fluentd / Elasticsearch / Kibana services for logging under --namespace=monitoring
	git submodule update --init
	cd etc/plugin/logging && ./deploy.sh
	kubectl --namespace=monitoring port-forward `kubectl --namespace=monitoring get pods -l k8s-app=kibana-logging -o json | jq '.items[0].metadata.name' -r` 35601:5601 &

grep-data:
	go run examples/grep/generate.go >examples/grep/set1.txt
	go run examples/grep/generate.go >examples/grep/set2.txt

grep-example:
	sh examples/grep/run.sh

logs: check-kubectl
	kubectl $(KUBECTLFLAGS) get pod -l app=pachd | sed '1d' | cut -f1 -d ' ' | xargs -n 1 -I pod sh -c 'echo pod && kubectl $(KUBECTLFLAGS) logs pod'

follow-logs: check-kubectl
	kubectl $(KUBECTLFLAGS) get pod -l app=pachd | sed '1d' | cut -f1 -d ' ' | xargs -n 1 -I pod sh -c 'echo pod && kubectl $(KUBECTLFLAGS) logs -f pod'

kubectl:
	gcloud config set container/cluster $(CLUSTER_NAME)
	gcloud container clusters get-credentials $(CLUSTER_NAME)

google-cluster-manifest:
	@pachctl deploy --rethinkdb-cache-size=5G --dry-run google $(BUCKET_NAME) $(STORAGE_NAME) $(STORAGE_SIZE)

google-cluster:
	gcloud container clusters create $(CLUSTER_NAME) --scopes storage-rw --machine-type $(CLUSTER_MACHINE_TYPE) --num-nodes $(CLUSTER_SIZE)
	gcloud config set container/cluster $(CLUSTER_NAME)
	gcloud container clusters get-credentials $(CLUSTER_NAME)
	gcloud components install kubectl
	-gcloud compute firewall-rules create pachd --allow=tcp:30650
	gsutil mb gs://$(BUCKET_NAME) # for PFS
	gcloud compute disks create --size=$(STORAGE_SIZE)GB $(STORAGE_NAME) # for PPS

clean-google-cluster:
	gcloud container clusters delete $(CLUSTER_NAME)
	gcloud compute firewall-rules delete pachd
	gsutil -m rm -r gs://$(BUCKET_NAME)
	gcloud compute disks delete $(STORAGE_NAME)

amazon-cluster-manifest: install
	@pachctl deploy --dry-run amazon $(BUCKET_NAME) $(AWS_ID) $(AWS_KEY) $(AWS_TOKEN) $(AWS_REGION) $(STORAGE_NAME) $(STORAGE_SIZE)

amazon-cluster:
	aws s3api create-bucket --bucket $(BUCKET_NAME) --region $(AWS_REGION)
	aws ec2 create-volume --size $(STORAGE_SIZE) --region $(AWS_REGION) --availability-zone $(AWS_AVAILABILITY_ZONE) --volume-type gp2

amazon-clean-cluster:
	aws s3api delete-bucket --bucket $(BUCKET_NAME) --region $(AWS_REGION)
	aws ec2 detach-volume --force --volume-id $(STORAGE_NAME)
	sleep 20
	aws ec2 delete-volume --volume-id $(STORAGE_NAME)

amazon-clean-launch: clean-launch
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found secrets amazon-secret
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found persistentvolumes rethink-volume
	kubectl $(KUBECTLFLAGS) delete --ignore-not-found persistentvolumeclaims rethink-volume-claim

amazon-clean:
	@while :; \
        do if echo "The following script will delete your AWS bucket and volume. The action cannot be undone. Do you want to proceed? (Y/n)";read REPLY; then \
        case $$REPLY in Y|y) make amazon-clean-launch;make amazon-clean-cluster;break;; \
	N|n) echo "The amazon clean process has been cancelled by user!";break;; \
	*) echo "input parameter error, please input again ";continue;;esac; \
        fi;done;

microsoft-cluster-manifest:
	@pachctl deploy --dry-run microsoft $(CONTAINER_NAME) $(AZURE_STORAGE_NAME) $(AZURE_STORAGE_KEY) $(VHD_URI) $(STORAGE_SIZE)

microsoft-cluster:
	azure group create --name $(AZURE_RESOURCE_GROUP) --location $(AZURE_LOCATION)
	azure storage account create $(AZURE_STORAGE_NAME) --location $(AZURE_LOCATION) --resource-group $(AZURE_RESOURCE_GROUP) --sku-name LRS --kind Storage
	$(eval _AZURE_STORAGE_KEY :="`azure storage account keys list $(AZURE_STORAGE_NAME) --resource-group $(AZURE_RESOURCE_GROUP) --json | jq .[0].value`")
	docker run -it microsoft_vhd $(AZURE_STORAGE_NAME) $(_AZURE_STORAGE_KEY) vhds $(DISK_NAME)

clean-microsoft-cluster:
	azure group delete $(AZURE_RESOURCE_GROUP) -q

install-go-bindata:
	go get -u github.com/jteeuwen/go-bindata/...

lint:
	@go get -u github.com/golang/lint/golint
	@for file in $$(find "./src" -name '*.go' | grep -v '/vendor/' | grep -v '\.pb\.go'); do \
		golint $$file; \
		if [ -n "$$(golint $$file)" ]; then \
			echo "golint errors!" && echo && exit 1; \
		fi; \
	done;

vet:
	@etc/testing/vet.sh

goxc-generate-local:
	@if [ -z $$GITHUB_OAUTH_TOKEN ]; then \
		echo "Missing token. Please run via: 'make GITHUB_OAUTH_TOKEN=12345 goxc-generate-local'"; \
		exit 1; \
	fi
	goxc -wlc default publish-github -apikey=$(GITHUB_OAUTH_TOKEN)

goxc-release:
	@if [ -z $$VERSION ]; then \
		@echo "Missing version. Please run via: 'make VERSION=v1.2.3-4567 VERSION_ADDITIONAL=4567 goxc-release'"; \
		@exit 1; \
	fi
	sed 's/%%VERSION_ADDITIONAL%%/$(VERSION_ADDITIONAL)/' .goxc.json.template > .goxc.json
	goxc -pv="$(VERSION)" -wd=./src/server/cmd/pachctl

goxc-build:
	sed 's/%%VERSION_ADDITIONAL%%/$(VERSION_ADDITIONAL)/' .goxc.json.template > .goxc.json
	goxc -tasks=xc -wd=./src/server/cmd/pachctl

.PHONY: all \
	version \
	deps \
	deps-client \
	update-deps \
	test-deps \
	update-test-deps \
	build-clean-vendored-client \
	build \
	install \
	install-doc \
	homebrew \
	release \
	release-worker \
	release-manifest \
	release-pachd \
	release-version \
	docker-build \
	docker-build-compile \
	docker-build-worker \
	docker-build-microsoft-vhd \
	docker-build-pachd \
	docker-build-proto \
	docker-push-worker \
	docker-push-pachd \
	docker-push \
	launch-kube \
	clean-launch-kube \
	kube-cluster-assets \
	launch \
	launch-dev \
	clean-launch \
	full-clean-launch \
	clean-pps-storage \
	integration-tests \
	proto \
	pretest \
	test \
	test-client \
	test-fuse \
	test-local \
	clean \
	doc \
	grep-data \
	grep-example \
	logs \
	kubectl \
	google-cluster-manifest \
	google-cluster \
	clean-google-cluster \
	amazon-cluster-manifest \
	amazon-cluster \
	amazon-clean-cluster \
	amazon-clean-launch \
	amazon-clean \
	install-go-bindata \
	assets \
	lint \
	goxc-generate-local \
	goxc-release \
	goxc-build \
	launch-test-rethinkdb \
	clean-launch-test-rethinkdb
