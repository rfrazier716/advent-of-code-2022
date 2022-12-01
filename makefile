SHELL := /bin/bash

include .env
AOC_URL = https://adventofcode.com

day_%:
	$(eval DAY := $(shell echo $$((10#$$(echo "$@" | tr -dc '0-9')))))
	$(info Creating Directory Structure for Day ${DAY})
	cp -r ./templates $@
	curl --cookie "session=${SESSION}" "${AOC_URL}/${YEAR}/day/${DAY}/input" -o $@/input.txt
	$(info Files Created. Find Today's Puzzle at: ${AOC_URL}/${YEAR}/day/${DAY})