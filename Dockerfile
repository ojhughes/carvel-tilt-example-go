FROM gcr.io/paketo-buildpacks/run:tiny-cnb
COPY build/carvel-tilt /app
ENTRYPOINT ["/app"]
