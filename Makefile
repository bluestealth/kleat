DOCKER ?= docker
DOCKER_EXTRA_ARGS ?= ""

protobuilder:
	$(DOCKER) build -q -t kleat-protobuilder:latest build/protoc/

proto_command = $(DOCKER) run \
  --rm \
  --user $(shell id -u):$(shell id -g) \
  --mount type=bind,source=$(abspath api/proto/),target=/mnt/proto,readonly \
  --mount type=bind,source=$(abspath api/client/),target=/mnt/output/go \
  --mount type=bind,source=$(abspath docs),target=/mnt/output/doc \
  $(DOCKER_EXTRA_ARGS) \
  kleat-protobuilder:latest

proto: protobuilder
	$(proto_command) update

checkproto: protobuilder
	$(proto_command) check
