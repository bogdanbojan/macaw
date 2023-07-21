FROM golang:1.20.6-bullseye AS base

RUN git clone https://github.com/bogdanbojan/macaw.git && \ 
    cd ./macaw && \
    git checkout automate-gh-releases

RUN apt-get update && \ 
    apt-get install xz-utils && \
    wget https://ziglang.org/download/0.9.1/zig-linux-x86_64-0.9.1.tar.xz && \
    tar -xf zig-linux-x86_64-0.9.1.tar.xz && \
    cd zig-linux-x86_64-0.9.1 && \
    ln -s $(pwd)/zig /usr/bin/

COPY /SDKs /SDKs

ENV MACOS_MIN_VER=10.14 
ENV MACOS_SDK_PATH="/SDKs/MacOSX11.sdk"
 
WORKDIR /go/macaw/

RUN mkdir -p dist/darwin-amd64 
 
RUN CGO_ENABLED=1 \
    GOOS=darwin \
    GOARCH=amd64 \
    CGO_LDFLAGS="-mmacosx-version-min=${MACOS_MIN_VER} --sysroot ${MACOS_SDK_PATH} -F/System/Library/Frameworks -L/usr/lib" \
    CC="zig cc -mmacosx-version-min=${MACOS_MIN_VER} -target x86_64-macos-gnu -isysroot ${MACOS_SDK_PATH} -iwithsysroot /usr/include -iframeworkwithsysroot /System/Library/Frameworks" \
    CXX="zig c++ -mmacosx-version-min=${MACOS_MIN_VER} -target x86_64-macos-gnu -isysroot ${MACOS_SDK_PATH} -iwithsysroot /usr/include -iframeworkwithsysroot /System/Library/Frameworks" \
    go build -ldflags="-w -buildmode=pie" -trimpath -o dist/darwin-amd64 .

RUN mkdir -p dist/windows-amd64

RUN CGO_ENABLED=1 \
    GOOS=windows \
    GOARCH=amd64 \
    CC="zig cc -target x86_64-windows-gnu" \
    CXX="zig c++ -target x86_64-windows-gnu" \
    go build -trimpath -ldflags='-H=windowsgui' -o dist/windows-amd64 .

FROM scratch AS binary
COPY --from=base /go/macaw/dist/darwin-amd64 /
COPY --from=base /go/macaw/dist/windows-amd64 /

