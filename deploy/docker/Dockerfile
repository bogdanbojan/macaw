FROM golang:1.20.6-bullseye AS base

RUN apt-get update && \ 
    apt-get install xz-utils && \
    wget https://ziglang.org/download/0.9.1/zig-linux-x86_64-0.9.1.tar.xz && \
    tar -xf zig-linux-x86_64-0.9.1.tar.xz && \
    cd zig-linux-x86_64-0.9.1 && \
    ln -s $(pwd)/zig /usr/bin/

COPY . /go/macaw/

WORKDIR /go/macaw/

# BUILD MACOS BINARY
############################################################################### 

ENV OSX_SDK="MacOSX11.3.sdk"
ENV OSX_SDK_URL="https://github.com/joseluisq/macosx-sdks/releases/download/11.3/${OSX_SDK}.tar.xz"

RUN curl -sSL "$OSX_SDK_URL" -o "/$OSX_SDK.tar.xz"
RUN mkdir /osxsdk && tar -xf "/$OSX_SDK.tar.xz" -C "/osxsdk"

ENV MACOS_MIN_VER=10.14 
ENV MACOS_SDK_PATH="/osxsdk/MacOSX11.3.sdk"

RUN mkdir -p dist/darwin-amd64 
 
RUN CGO_ENABLED=1 \
    GOOS=darwin \
    GOARCH=amd64 \
    CGO_LDFLAGS="-mmacosx-version-min=${MACOS_MIN_VER} --sysroot ${MACOS_SDK_PATH} -F/System/Library/Frameworks -L/usr/lib" \
    CC="zig cc -mmacosx-version-min=${MACOS_MIN_VER} -target x86_64-macos-gnu -isysroot ${MACOS_SDK_PATH} -iwithsysroot /usr/include -iframeworkwithsysroot /System/Library/Frameworks" \
    CXX="zig c++ -mmacosx-version-min=${MACOS_MIN_VER} -target x86_64-macos-gnu -isysroot ${MACOS_SDK_PATH} -iwithsysroot /usr/include -iframeworkwithsysroot /System/Library/Frameworks" \
    go build -ldflags="-w -buildmode=pie" -trimpath -o dist/darwin-amd64/macaw-darwin-amd64 ./cmd/macaw/main.go

# BUILD WINDOWS BINARY
############################################################################### 

RUN mkdir -p dist/windows-amd64

RUN CGO_ENABLED=1 \
    GOOS=windows \
    GOARCH=amd64 \
    CC="zig cc -target x86_64-windows-gnu" \
    CXX="zig c++ -target x86_64-windows-gnu" \
    go build -trimpath -ldflags='-H=windowsgui' -o dist/windows-amd64/macaw-windows-amd64.exe ./cmd/macaw/main.go

# BUILD LINUX BINARY
############################################################################### 

RUN apt-get update && \
    apt-get install -y -q --no-install-recommends \
        libgl-dev \
        libx11-dev \
        libxrandr-dev \
        libxxf86vm-dev \
        libxi-dev \
        libxcursor-dev \
        libxinerama-dev

RUN mkdir -p dist/linux-amd64

RUN CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    go build -o dist/linux-amd64/macaw-linux-amd64 ./cmd/macaw/main.go

FROM scratch AS binaries
COPY --from=base /go/macaw/dist/darwin-amd64 /
COPY --from=base /go/macaw/dist/windows-amd64 /
COPY --from=base /go/macaw/dist/linux-amd64 /
