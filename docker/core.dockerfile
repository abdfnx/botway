FROM buildpack-deps:20.04-curl AS download

RUN export DEBIAN_FRONTEND=noninteractive \
  && apt-get update \
  && apt-get install -y unzip \
  && rm -rf /var/lib/apt/lists/*

RUN curl -s https://get-latest.herokuapp.com/abdfnx/botway >> tag.txt

RUN curl -fsSL "https://github.com/abdfnx/botway/releases/download/$(cat tag.txt)/botway_linux_$(cat tag.txt)_amd64.zip" \
  --output botway.zip \
  && unzip botway.zip \
  && rm botway.zip \
  && mv "botway_linux_$(cat tag.txt)_amd64/bin/botway" . \
  && chmod 755 botway

FROM scratch

COPY --from=download /botway /botway
