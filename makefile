SHELL := /bin/bash

include .env
AOC_URL = https://adventofcode.com/

day_%:
	$(eval DAY := $(shell echo $$((10#$$(echo "$@" | tr -dc '0-9')))))
	$(info Creating Directory Structure for Day ${DAY})
	mkdir $@
	echo "package main" > $@/main.go
	echo "package main" > $@/main_test.go
	curl --cookie "session=${SESSION}" "${AOC_URL}/${YEAR}/day/${DAY}/input" -o $@/input.txt