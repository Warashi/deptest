.DELETE_ON_ERROR:

self := $(lastword $(MAKEFILE_LIST))
GO ?= go
PROFILE_NAME ?= cover.out
TEST_COMMAND ?= $(GO) test
DEPTEST ?= $(GO) run github.com/Warashi/deptest/cmd/deptest@latest
GOCOVMERGE ?= $(GO) run github.com/wadey/gocovmerge@latest

MODULE := $(shell GOWORK=off $(GO) list -m)
PACKAGES := $(sort $(shell $(GO) list ./...))
PROFILES := $(patsubst $(MODULE)/%,%/$(PROFILE_NAME),$(PACKAGES))

.PHONY: all
all: $(PROFILE_NAME)

.PHONY: update-testdeps
update-testdeps:
	$(DEPTEST) -module $(MODULE) $(PROFILES) > $(self).tmp
	mv $(self).tmp $(self)

.PHONY: clean
clean:
	rm -f $(PROFILES)

$(PROFILE_NAME): $(PROFILES)
	$(GOCOVMERGE) $(PROFILES) > $(PROFILE_NAME)

$(PROFILES): %/$(PROFILE_NAME):
	$(TEST_COMMAND) -covermode=atomic -coverpkg=./... -coverprofile=./$@ ./$(dir $@)


a/cover.out: a/a.go 
b/cover.out: b/b.go 
c/cover.out: b/b.go c/c.go c/c_test.go 
