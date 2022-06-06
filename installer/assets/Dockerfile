FROM rust:alpine

ENV PACKAGES "build-dependencies xz libxml2-dev build-base openssl cmake openssl-dev musl-dev libressl-dev bsd-compat-headers gcc g++ git lld clang clang-dev patch fts-dev gmp-dev mpfr-dev mpc1-dev perl libsodium ffmpeg opus autoconf automake libtool"

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

RUN rustup target add x86_64-apple-darwin
COPY --from=crazymax/osxcross:latest-alpine /osxcross /osxcross
ENV PATH="/osxcross/target/bin:$PATH"
RUN echo $PATH
ENV CC="o64-clang"
ENV CXX="o64-clang++"
ENV LD_LIBRARY_PATH="/osxcross/target/lib:$LD_LIBRARY_PATH"
RUN cd ./botway/installer && \
	cargo build --release --target x86_64-apple-darwin && \
	cp ${DARWIN_TARGET_PATH}/botway-installer ./installer-macos && \
 	cp ./installer-macos ./assets/public/installers

RUN ls -a ./botway/installer/assets/public/installers
