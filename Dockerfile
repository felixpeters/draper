# Get latest official Go image
FROM golang:latest
# Create app folder
RUN mkdir /app
# Add project files to folder
ADD . /app
# Set working directory to app folder
WORKDIR /app
# Build the app binary
RUN go build -o draper .
# Run the app binary
CMD ["/app/draper"]

