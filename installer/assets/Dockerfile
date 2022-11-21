FROM rust:alpine AS rs

ENV PACKAGES "build-dependencies xz libxml2-dev libpq-dev build-base openssl cmake openssl-dev musl-dev libressl-dev bsd-compat-headers gcc g++ git lld clang clang-dev patch fts-dev gmp-dev mpfr-dev mpc1-dev perl libsodium ffmpeg opus autoconf automake libtool bash"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

RUN git clone https://github.com/abdfnx/botway

# Build linux target
ENV LINUX_TARGET_PATH="./target/release"

RUN cd ./botway/installer && \
	cargo build --release && \
	cp ${LINUX_TARGET_PATH}/botway-installer ./installer-linux && \
 	cp ./installer-linux ./assets/static/installers

# Build macos target
ENV DARWIN_TARGET_PATH="./target/x86_64-apple-darwin/release"

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
 	cp ./installer-macos ./assets/static/installers

FROM denoland/deno:alpine

ENV INSTALLERS_DIR_PATH="./static/installers"

COPY . .

RUN rm -rf ${INSTALLERS_DIR_PATH}

COPY --from=rs ./botway/installer/assets/${INSTALLERS_DIR_PATH} ${INSTALLERS_DIR_PATH}

ENTRYPOINT ["deno", "task", "start"]

EXPOSE 8000
