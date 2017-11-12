#!/usr/bin/env bash

exe_name="webapi-server"
platforms=("linux/amd64" "windows/amd64" "darwin/amd64")
gopath="/home/leonardo/Documents/Go/webapi-dav/"

export GOPATH=${gopath}

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=${exe_name}'-'${GOOS}'_'${GOARCH}
    echo -e "Building $output_name\n"
    if [ ${GOOS} = "windows" ]; then
        output_name+='.exe'
    fi  

    env GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${output_name}
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

mv -f webapi-* ../build