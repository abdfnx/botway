FROM alpine AS download

RUN apk update && apk add unzip curl
RUN curl -s https://get-latest.deno.dev/abdfnx/botway >> tag.txt

RUN curl -fsSL "https://github.com/abdfnx/botway/releases/download/$(cat tag.txt)/botway_linux_$(cat tag.txt)_amd64.zip" \
  --output botway.zip \
  && unzip botway.zip \
  && rm botway.zip \
  && mv "botway_linux_$(cat tag.txt)_amd64/bin/botway" . \
  && chmod 755 botway

FROM scratch

COPY --from=download /botway /botway
