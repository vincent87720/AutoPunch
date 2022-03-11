WINDOWS := windows
PACKAGEPIN := github.com/vincent87720/AutoPunch/punchin
PACKAGEPOUT := github.com/vincent87720/AutoPunch/punchout

all: build

build: release

##########BUILD##########
.PHONY: buildwindows
buildwindows:
	GOOS=$(WINDOWS) GOARCH=amd64 go build -o bin/punchin.exe $(PACKAGEPIN)
	GOOS=$(WINDOWS) GOARCH=amd64 go build -o bin/punchout.exe $(PACKAGEPOUT)

.PHONY: release
release: buildwindows