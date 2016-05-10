help:
	@echo 'make fuck    - build the fuck executable'
	@echo 'make tag     - tag the current HEAD with VERSION'
	@echo 'make archive - create an archive of the current HEAD for VERSION'
	@echo 'make all     - build, tag and archive VERSION'

version:
	@if [ "$$VERSION" = "" ]; then echo "VERSION not set"; exit 1; fi

fuck: main.go
	go build $^

build: fuck

tag: version
	git tag -s $$VERSION -m "$$VERSION release"

archive: fuck-$$VERSION.zip

fuck-$$VERSION.zip: version fuck
	git archive -o $@ HEAD
	zip $@ fuck

all: version build tag archive

.PHONY: clean
clean:
	rm -f fuck fuck-*.zip