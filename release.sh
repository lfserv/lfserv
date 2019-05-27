#!/bin/bash -e
#
# This script will generate a release on github/lfsserv.
# Ensure that you've bumped version in main.go, then run the script.
# The script does the following
#   * Ensure the build succeeds (and pulls the version from the build)
#   * Ensure the tests pass
#   * Cross compiles for supported platforms
#   * Creates the release on github/lfserv
#   * Uploads binary assets to the release.

name="lfserv"

function validate() {
    go build -o lfserv ${name}.go
    rc=$?; if [[ ${rc} != 0 ]]; then echo "Build failed."; exit ${rc}; fi
    version=$(./lfserv -v)

    while true
    do
      read -p "Release version ${version}? [y/n] " yn
      case ${yn} in
        [Yy]* ) break;;
        [Nn]* ) exit;;
        * ) echo "Please answer yes or no.";;
      esac
    done
# Make sure tests pass
    echo "Running tests..."
    go test
    rc=$?
    if [[ ${rc} != 0 ]]
    then
        echo "Tests failed, cannot release."
        exit ${rc}
    fi
}

function build() {
# Build all files
    rm -rf dist
    mkdir dist
    for os in "darwin" "linux" "freebsd"
    do
        echo "Building ${os} amd64"
        mkdir -p dist/${os}-amd64
        GOOS=${os} GOARCH=amd64 go build -o dist/${os}-amd64/${name} ${name}.go
        cp README.md dist/${os}-amd64
        cp LICENSE dist/${os}-amd64
        cd dist
        tar zcf ${name}-${os}-amd64-${version}.tar.gz ${os}-amd64
        cd ..
    done
    echo "Building windows amd64"
    mkdir -p dist/windows-amd64
    GOOS=windows GOARCH=amd64 go build -o dist/windows-amd64/${name}.exe ${name}.go
    cp README.md dist/windows-amd64
    cp LICENSE dist/windows-amd64
    cd dist
    zip -q -j ${name}-windows-amd64-${version}.zip windows-amd64/*
    cd ..
}

function create_release() {
    payload=$(cat <<EOF
{
  "tag_name": "v${version}",
  "target_commitish": "master",
  "name": "Release ${version}",
  "draft": false,
  "prerelease": false,
  "body": ""
}
EOF
)
    echo "${payload}" > ${tmpl}
    ${EDITOR:-vim} ${tmpl}
    curl -n -X POST -d @${tmpl} -o ${out} https://api.github.com/repos/lfserv/lfserv/releases
    id=$(cat ${out} | jq -r ".id")
    if [[ ${id} == "null" ]]
    then
        echo "Failed creating release."
        cat ${out}
        exit 1
    fi
    echo "Created release id: ${id}"
}

function upload() {
# Upload each file to the release
    upload=$(cat ${out} | jq -r ".upload_url" | sed s/"{?name}"//)
    for os in "darwin" "linux" "freebsd"
        do
            echo "Uploading ${name}-${os}-amd64-${version}.tar.gz"
            curl -n -o ${out} -H "Content-Type: application/octet-stream" -X POST --data-binary @dist/${name}-darwin-amd64-${version}.tar.gz "${upload}?name=${name}-${os}-amd64-${version}.tar.gz&label=${os}%20AMD64"
        done
    echo "Uploading ${name}-windows-amd64-${version}"
    curl -n -o ${out} -H "Content-Type: application/octet-stream" -X POST --data-binary @dist/${name}-windows-amd64-${version}.zip "${upload}?name=${name}-windows-amd64-${version}.zip&label=windows%20AMD64"
}

# Create the release
tmpl=$(mktemp ${name}-release.XXXXXXXXX)
out=$(mktemp ${name}-out.XXXXXXXX)
version=

validate
build
#create_release
#upload

rm -f ${tmpl} ${out}
rm -rf dist
