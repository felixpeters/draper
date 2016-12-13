# Get latest official Go image
FROM golang
# Add project files to folder
ADD . /go/src/github.com/felixpeters/draper 
# Install the app
RUN go install github.com/felixpeters/draper
# Run the binary
ENTRYPOINT /go/bin/draper
# Expose port 8080
EXPOSE 8080
