GOBIN = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN = $(shell go env GOPATH)/bin
endif
ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif
ifeq ($(XDG_CONFIG_HOME),)
    XDG_CONFIG_HOME := $$HOME/.config
endif
.PHONY: help
help:
	@echo "Usage: [variables] make <target>"
	@echo
	@echo "This Makefile makes use of dependency lists."
	@echo "The artifacts are compiled only if any of their dependencies are newer than them."
	@echo
	@echo "Commands:"
	@echo "\tbuild         \tBuilds the binary."
	@echo "\tinstall       \tInstall the binary, a basic config file and a desktop file."
	@echo "\t              \tSet the value of the PREFIX (default: /usr/local) to change the installation location."
	@echo
	@echo "Utilities:"
	@echo "\tclear         \tRemoves all build artifacts."

.PHONY: install
install:
	OS=linux ARCH=amd64 CMD=browser-proxy $(MAKE) compile-and-install

.PHONY: compile-and-install
compile-and-install:
	OS=$(OS) ARCH=$(ARCH) CMD=$(CMD) $(MAKE) compile
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 755 build/$(CMD)/$(CMD)_$(OS)_$(ARCH) $(DESTDIR)$(PREFIX)/bin/$(CMD)
	install -d $(XDG_CONFIG_HOME)/browser-local
	install -b -S .old configs/config.json.template $(XDG_CONFIG_HOME)/browser-local/config.json
	install -d $(PREFIX)/share/applications
	install -m 764 configs/browser-proxy.desktop.template $(PREFIX)/share/applications/browser-proxy.desktop
	sed -i'' "s#Exec\=browser\-proxy \%u#Exec\=\"$$(realpath $(DESTDIR)$(PREFIX)/bin/$(CMD))\" \%u#" $(PREFIX)/share/applications/browser-proxy.desktop
	update-desktop-database $(PREFIX)/share/applications

.PHONY: build
build:
	OS=linux ARCH=amd64 CMD=browser-proxy $(MAKE) compile

.PHONY: compile
compile: build/$(CMD)/$(CMD)_$(OS)_$(ARCH)

build/$(CMD)/$(CMD)_$(OS)_$(ARCH): $(shell find cmd/$(CMD) -type f -name '*.go' -print)
	mkdir -p $(shell dirname $@)
	rm -f $@
	CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -o $@ ./cmd/$(CMD)
	chmod 0700 build/$(CMD)/*

clear:
	rm -rf build/*