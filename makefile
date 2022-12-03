SHELL := /bin/bash

include .env
AOC_URL = https://adventofcode.com
ROOT_PATH = AoC_${YEAR}

.PHONY: day_%
day_%: ${ROOT_PATH}/day_%
	
${ROOT_PATH}/day_%:
	$(eval DAY := $(shell echo $$((10#$$(echo "$(@F)" | tr -dc '0-9')))))
	$(info Creating Directory Structure for Day ${DAY})
	mkdir -p ${ROOT_PATH}
	cp -r ./templates $@
	curl --cookie "session=${SESSION}" "${AOC_URL}/${YEAR}/day/${DAY}/input" -o $@/input.txt
	$(info Files Created. Find Today's Puzzle at: ${AOC_URL}/${YEAR}/day/${DAY})