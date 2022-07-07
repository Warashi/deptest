GO ?= go
DEPTEST ?= github.com/Warashi/deptest/cmd/deptest@latest
PROFILE_NAME ?= cover.out

CODES := $(shell find . -name *.go)
DIRS := $(sort $(dir $(CODES)))
PROFILES := $(patsubst %/,%/$(PROFILE_NAME),$(DIRS))
GO_MODULE := $(shell $(GO) list -m)

deps = $(foreach d,$(shell $(GO) run $(DEPTEST) $(1)),$(wildcard $(d:$(GO_MODULE)/%=./%)/*.go))

.PHONY: all
all: $(PROFILES)

$(foreach p,$(PROFILES),$(eval $(p): $(call deps,$(p))))
$(PROFILES): %/$(PROFILE_NAME):
	$(GO) test -coverpkg=./... -coverprofile=./$@ ./$(dir $@)
