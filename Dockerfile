# Get latest official Go image
FROM golang
# Add project files to folder
ADD . /go/src/github.com/felixpeters/draper 
# Install the app
RUN go install github.com/felixpeters/draper
# Expose port 8080
EXPOSE 8080
# Run the binary
CMD ["/go/bin/draper"]
