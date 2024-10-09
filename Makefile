# Makefile to create a module folder structure

MODULE_NAME := $(word 2, $(MAKECMDGOALS))

module:
	@[ "$(MODULE_NAME)" ] || (echo "Error: No module name provided"; exit 1)
	@mkdir -p internal/module/$(MODULE_NAME)/entity
	@mkdir -p internal/module/$(MODULE_NAME)/handler/rest
	@mkdir -p internal/module/$(MODULE_NAME)/ports
	@mkdir -p internal/module/$(MODULE_NAME)/repository
	@mkdir -p internal/module/$(MODULE_NAME)/service
	@touch internal/module/$(MODULE_NAME)/entity/entity.go
	@touch internal/module/$(MODULE_NAME)/entity/$(MODULE_NAME).go
	@touch internal/module/$(MODULE_NAME)/handler/rest/handler.go
	@touch internal/module/$(MODULE_NAME)/ports/ports.go
	@touch internal/module/$(MODULE_NAME)/repository/repo.go
	@touch internal/module/$(MODULE_NAME)/service/service.go
	@touch internal/module/$(MODULE_NAME)/$(MODULE_NAME).go
	@echo "Module '$(MODULE_NAME)' created successfully."

clean:
	@[ "$(MODULE_NAME)" ] || (echo "Error: No module name provided"; exit 1)
	@rm -rf internal/module/$(MODULE_NAME)
	@echo "Module '$(MODULE_NAME)' removed successfully."

.PHONY: module clean
