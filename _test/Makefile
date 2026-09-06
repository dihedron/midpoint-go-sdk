include ../help.mk

default: help

midpoint-quickstart.sh: ## Download the midpoint quickstart script
	@mkdir -p midpoint
	@cd midpoint && curl -O https://raw.githubusercontent.com/Evolveum/midpoint-docker/master/midpoint-quickstart.sh && chmod 744 midpoint-quickstart.sh && cd ..

.phony: midpoint-start
midpoint-start: midpoint-quickstart.sh ## Create or restart the midpoint instance
	@mkdir -p midpoint
	@cd midpoint && ./midpoint-quickstart.sh start --password 'S3cret!!' && cd ..

.phony: midpoint-stop
midpoint-stop: midpoint-quickstart.sh ## Stop the midpoint instance
	@cd midpoint && ./midpoint-quickstart.sh stop && cd ..

.phony: midpoint-info
midpoint-info: midpoint-quickstart.sh ## Print info about the midpoint instance
	@cd midpoint && ./midpoint-quickstart.sh info && cd ..

.phony: midpoint-delete
midpoint-delete: midpoint-quickstart.sh ## Destroy the midpoint instance (including data)
	@cd midpoint && ./midpoint-quickstart.sh delete && cd ..
	@rm -rf midpoint

.phony: user-self ## Show info about the current user
user-self:
	@. ./.env && ../dist/midpoint_linux_amd64_v1/midpoint self

.phony: user-read
user-read: ## Show info about a specific user
	@. ./.env && ../dist/midpoint_linux_amd64_v1/midpoint user read 00000000-0000-0000-0000-000000000002

.phony: user-create
user-create: ## Create a new user
	@. ./.env && ../dist/midpoint_linux_amd64_v1/midpoint user create --user=@user-create.json
