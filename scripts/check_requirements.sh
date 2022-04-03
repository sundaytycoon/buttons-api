#!/bin/sh

# 이 프로젝트를 운영하는데에 필요한 모든 프로그램을 검진하는 script입니다.

if [[ ! -f `which docker` ]]; then
  echo "Required 'docker' cli. https://www.docker.com/products/docker-desktop"
fi

if [[ ! -f `which mockgen` ]]; then
  echo "Required 'mockgen' cli. https://github.com/golang/mock"
fi

if [[ ! -f `which buf` ]]; then
  echo "Required 'buf' cli https://docs.buf.build/installation"
  echo "godepgraph is static code analytics tool for golang project it's for generate dependency graph"
  echo "[Quick installation]"
  echo "brew install bufbuild/buf/buf"
fi

if [[ ! -f `which jq` ]]; then
  echo "Required 'godepgraph' cli https://stedolan.github.io/jq/"
  echo "jq is utility at terminal environment for manupulating JSON objects"
  echo "[Quick installation]"
  echo "brew install jq"
fi

if [[ ! -f `which godepgraph` ]]; then
  echo "Required 'godepgraph' cli https://github.com/kisielk/godepgraph"
  echo "godepgraph is static code analytics tool for golang project it's for generate dependency graph"
  echo "[Quick installation]"
  echo "1. install graphviz"
  echo "brew install graphviz" # for godepgraph

  echo "2. install godepgraph"
  echo "go install github.com/kisielk/godepgraph@latest"
fi

if [[ ! -f `which direnv` ]]; then
  echo "Required 'direnv' cli. https://github.com/direnv/direnv"
  echo "direnv inject environments at current terminal. commonly useful tool"
fi

if [[ ! -f `which oapi-codegen` ]]; then
  echo "Required 'oapi-codegen' cli."
  echo "go get -u github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
fi

if [[ ! -f `which mockgen` ]]; then
  echo "Required 'mockgen' cli."
  echo "go get -u github.com/golang/mock/mockgen"
fi


