.PHONY: all create_dirs

all: create_dirs

create_dirs:
	@mkdir -p cmd
	@mkdir -p db/migrations
	@mkdir -p internal/adapter
	@mkdir -p internal/application
	@mkdir -p internal/grpc
	@mkdir -p internal/port
	@mkdir -p ssl