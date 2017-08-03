# Version
VERSION = `date +%y.%m`

# If unable to grab the version, default to N/A
ifndef VERSION
    VERSION = "n/a"
endif

#
# Makefile options
#


# State the "phony" targets
.PHONY: all clean build


all: build

build:
	@echo 'Building goplex...'
	@go build

clean:
	@echo 'Cleaning...'
	@go clean
