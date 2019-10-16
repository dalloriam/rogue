NAME := rogue
PKG := github.com/dalloriam/$(NAME)

CGO_ENABLED := 1

BUILDTAGS :=

include root.mk

.PHONY: prebuild
prebuild:

.PHONY: extra_validation
extra_validation:
