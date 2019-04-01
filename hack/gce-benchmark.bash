#!/bin/bash

start_time="$(date -u +%s)"

echo 'Install google-fluentd ...'
curl -sSL https://dl.google.com/cloudagents/install-logging-agent.sh | sudo bash

echo 'Restart google-fluentd ...'
service google-fluentd restart

# ----------------------------------------------------------------------------
# Start tasks

sudo apt-get update
sudo apt-get install -yqq --no-install-recommends \
  ca-certificates \
  curl \
  netbase \
  wget \
  \
  bzr \
  git \
  mercurial \
  openssh-client \
  subversion \
  \
  procps \
  \
  g++ \
  gcc \
  libc6-dev \
  make \
  pkg-config

export GOLANG_VERSION=1.12.1

wget -qq -O ~/go.tgz "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"
sudo tar -C /usr/local -xzf ~/go.tgz
rm ~/go.tgz

export GOPATH=~/go
export GOCACHE=~/go-build
export PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"
mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

go version

go get -u -v golang.org/x/perf/cmd/benchstat
go get -u -v github.com/zchee/color

export GO111MODULE=on

cd "$GOPATH/src/github.com/zchee/color/benchmarks" || return
go mod download
go mod tidy -v
go mod vendor -v

REVISION="$(git rev-parse --short -q HEAD)"
go test -v -mod=vendor -tags=benchmark -count 10 -run='^$' -bench=. | tee "new.$REVISION.txt" || true
benchstat "old.$REVISION.txt" "new.$REVISION.txt" | tee "benchstat.$REVISION.txt" || true

gsutil cp "old.$REVISION.txt" "new.$REVISION.txt" "benchstat.$REVISION.txt" "$(curl -s -H 'Metadata-Flavor:Google' http://metadata.google.internal/computeMetadata/v1/instance/attributes/gsutil_benchstat_bucket_name)" || true

# End tasks
# ----------------------------------------------------------------------------

end_time="$(date -u +%s)"
elapsed="$(("$end_time-$start_time"))"
echo "Total of $elapsed seconds elapsed for tasks"

echo 'Delete own GCE instance ...'
yes | gcloud compute instances delete "$(hostname)" --zone "$(curl -s -H 'Metadata-Flavor:Google' http://metadata.google.internal/computeMetadata/v1/instance/zone | cut -d/ -f4)"
