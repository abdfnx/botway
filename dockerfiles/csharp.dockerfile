FROM botwayorg/botway:latest AS bw

COPY . .

RUN botway init --docker

FROM mcr.microsoft.com/dotnet/sdk:6.0-alpine AS build

WORKDIR /source

COPY *.csproj .

RUN dotnet restore -r linux-musl-x64

COPY . .

RUN dotnet publish -c release -o /app -r linux-musl-x64 --self-contained false --no-restore

FROM mcr.microsoft.com/dotnet/runtime:6.0-alpine-amd64

ENV \
    DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=false \
    LC_ALL=en_US.UTF-8 \
    LANG=en_US.UTF-8 \
    PACKAGES="build-dependencies build-base gcc git libsodium opus ffmpeg icu-libs binutils openssl-dev zlib-dev boost boost-dev"

RUN apk update && \
	apk add --no-cache --virtual ${PACKAGES}

COPY --from=build /app /app

COPY --from=bw /root/.botway /root/.botway

COPY . .

ENTRYPOINT ["./app/{{.BotName}}"]
