# made for GNU make

.ONESHELL:
REPO := test

all-images: web_server stream_worker batch_worker

web_server: command-check
		docker build . -f docker/web_server.dockerfile -t ${REPO}/samples/web_server:v1.0

stream_worker: command-check
		docker build . -f docker/stream_worker.dockerfile -t ${REPO}/samples/stream_worker:v1.0

batch_worker: command-check
		docker build . -f docker/batch_worker.dockerfile -t ${REPO}/samples/batch_worker:v1.0

define COMMAND_ERROR = 
Error: 'docker' command not found. 
			 Please install Docker command interface.
endef
command-check:
		@command -v docker 2>&1 >/dev/null || { echo "$(COMMAND_ERROR)"; exit 1; }

define REPOSITORY_ERROR =
ERROR: REPO is unknown.
			 Please run 'make REPO=xxx target'.
			 ex. 'make REPO=REPO=asia-northeast1-docker.pkg.dev/{PROJECT_ID} target'
endef
release-check:
		@test ${REPO} = "test" && { echo "${REPOSITORY_ERROR}"; exit 1; } || exit 0

push: release-check all-images
		docker push ${REPO}/samples/web_server:v1.0
		docker push ${REPO}/samples/stream_worker:v1.0
		docker push ${REPO}/samples/batch_worker:v1.0
