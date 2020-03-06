#!/bin/bash

# run GCE instance command
#  $ GOLANG_VERSION='1.12.6 or whatever' GOOGLE_CLOUD_PROJECT='foo-project' ZONE='asia-northeast1-a' MACHINE_TYPE='n1-standard-{n}' BENCHSTAT_BUCKET_NAME='gs://foo-bucket/'; gcloud --project="$GOOGLE_CLOUD_PROJECT" compute instances create --zone="$ZONE" --machine-type="$MACHINE_TYPE" --image-project='debian-cloud' --image-family='debian-9' --boot-disk-type='pd-ssd' --preemptible --scopes='https://www.googleapis.com/auth/cloud-platform' --service-account='benchstat@foo-project.iam.gserviceaccount.com' --metadata="golang_version=${GOLANG_VERSION},benchstat_bucket_name=${BENCHSTAT_BUCKET_NAME}" --metadata-from-file="startup-script=$(go env GOPATH)/src/github.com/zchee/color/hack/gce-benchmark.bash" --async --verbosity='debug' 'benchstat'

set -x

start_time="$(date -u +%s)"

echo 'Install google-fluentd ...'
curl -sSL https://dl.google.com/cloudagents/install-logging-agent.sh | sudo bash

echo 'Restart google-fluentd ...'
service google-fluentd restart

# ----------------------------------------------------------------------------
# Start tasks

sudo apt-get update
sudo apt-get install -yqq --no-install-recommends --no-install-suggests \
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

PACKAGE='github.com/zchee/color'

GOLANG_VERSION="$(curl -s -H 'Metadata-Flavor:Google' http://metadata.google.internal/computeMetadata/v1/instance/attributes/golang_version)"

wget -qq -O ~/go.tgz "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"
sudo tar -C /usr/local -xzf ~/go.tgz
rm ~/go.tgz

export GOPATH=~/go
export GOCACHE=~/.cache/go-build
export PATH="$GOPATH/bin:/usr/local/go/bin:$PATH"
mkdir -p "$GOPATH/src" "$GOPATH/bin" "$GOCACHE" && chmod -R 777 "$GOPATH" "$GOCACHE"

go env
go version

export GO111MODULE=off

go get -u -v golang.org/x/perf/cmd/benchstat
go get -u -d -v github.com/zchee/color

export GO111MODULE=on

cd "$GOPATH/src/$PACKAGE/benchmarks" || return

go mod download
go mod tidy -v
go mod vendor -v

REVISION="$(git rev-parse --short -q HEAD)"
BENCH_OUT_BASE="$PWD/$(echo "$PACKAGE" | tr '/' '-')@$REVISION"

#  Clear PageCache, dentries and inodes
sync; echo 3 > /proc/sys/vm/drop_caches
go test -v -mod=vendor -tags=benchmark -run='^$' -count 8 -bench=. . -fatih | tee "${BENCH_OUT_BASE}.old.txt" 2>&1 | go tool test2json

#  Clear PageCache, dentries and inodes
sync; echo 3 > /proc/sys/vm/drop_caches
go test -v -mod=vendor -tags=benchmark -run='^$' -count 8 -bench=. . | tee "${BENCH_OUT_BASE}.new.txt" 2>&1 | go tool test2json

benchstat "${BENCH_OUT_BASE}.old.txt" "${BENCH_OUT_BASE}.new.txt" | tee "${BENCH_OUT_BASE}.benchstat.txt"

CPUINFO_OUT="${BENCH_OUT_BASE}.cpuinfo.txt"
cat /proc/cpuinfo > "$CPUINFO_OUT"

gsutil cp "${BENCH_OUT_BASE}.old.txt" "${BENCH_OUT_BASE}.new.txt" "${BENCH_OUT_BASE}.benchstat.txt" "$CPUINFO_OUT" "$(curl -s -H 'Metadata-Flavor:Google' http://metadata.google.internal/computeMetadata/v1/instance/attributes/benchstat_bucket_name)" || true

# End tasks
# ----------------------------------------------------------------------------

end_time="$(date -u +%s)"
elapsed="$(("$end_time-$start_time"))"
echo "Total of $elapsed seconds elapsed for tasks"

echo 'Delete own GCE instance ...'
yes | gcloud compute instances delete "$(hostname)" --zone "$(curl -s -H 'Metadata-Flavor:Google' http://metadata.google.internal/computeMetadata/v1/instance/zone | cut -d/ -f4)"
