FROM rust:alpine as rs

ENV PACKAGES "build-dependencies xz libxml2-dev libpq-dev build-base openssl cmake openssl-dev musl-dev libressl-dev bsd-compat-headers gcc g++ git lld clang clang-dev patch fts-dev gmp-dev mpfr-dev mpc1-dev perl libsodium ffmpeg opus autoconf automake libtool"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

RUN git clone https://github.com/abdfnx/botway

# Build linux target
ENV LINUX_TARGET_PATH="./target/release"

RUN cd ./botway/installer && \
	cargo build --release && \
	cp ${LINUX_TARGET_PATH}/botway-installer ./installer-linux && \
 	cp ./installer-linux ./assets/public/installers

# Build macos target
ENV DARWIN_TARGET_PATH="./target/x86_64-apple-darwin/release"
ENV DARWIN_FILE_NAME="rust-std-1.61.0-x86_64-apple-darwin"

RUN wget https://static.rust-lang.org/dist/${DARWIN_FILE_NAME}.tar.gz && \
	tar -xvxf ${DARWIN_FILE_NAME}.tar.gz && \
	cd ${DARWIN_FILE_NAME} && \
	./install.sh

RUN git clone https://github.com/abdfnx/osxcross-alpine osxcross && \
	cd osxcross && \
	wget -nc https://s3.dockerproject.org/darwin/v2/MacOSX10.10.sdk.tar.xz && \
	mv MacOSX10.10.sdk.tar.xz tarballs/ && \
	UNATTENDED=yes OSX_VERSION_MIN=10.7 ./build.sh

ENV PATH="/osxcross/target/bin:$PATH"

RUN rustup target add x86_64-apple-darwin

RUN cd ./botway/installer && \
	PATH="$(pwd)/osxcross/target/bin:$PATH" \
	CC=o64-clang \
	CXX=o64-clang++ \
	LIBZ_SYS_STATIC=1 \
	cargo build --release --target x86_64-apple-darwin && \
	cp ${DARWIN_TARGET_PATH}/botway-installer ./installer-macos && \
 	cp ./installer-macos ./assets/public/installers

FROM node:alpine

ENV INSTALLERS_DIR_PATH="./public/installers"

COPY . .

RUN rm -rf ${INSTALLERS_DIR_PATH}

COPY --from=rs ./botway/installer/assets/${INSTALLERS_DIR_PATH} ${INSTALLERS_DIR_PATH}

RUN yarn
