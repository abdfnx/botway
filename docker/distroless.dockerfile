FROM botwayorg/botway-cli:core AS core
FROM gcr.io/distroless/cc

ENV BOTWAY_DIR /botway-dir/

COPY --from=core /botway /bin/botway

ENTRYPOINT ["/bin/botway"]
CMD [ "help" ]
