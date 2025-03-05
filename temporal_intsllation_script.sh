#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" &> /dev/null
}

# Check if Homebrew is installed
if ! command_exists brew; then
    echo "Homebrew not found. Installing Homebrew..."
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
else
    echo "Homebrew is already installed. Updating..."
    brew update
fi

# Check if Go is installed
if ! command_exists go; then
    echo "Go not found. Installing Go..."
    brew install go
else
    echo "Go is already installed."
fi

# Install Temporal via Homebrew
echo "Installing Temporal..."
brew install temporal
# temporal server start-dev

# Define the Temporal project directory
TEMPORAL_DIR="$HOME/Documents/temporal_poc"

# Create the directory if it doesn't exist
if [ ! -d "$TEMPORAL_DIR" ]; then
    echo "Creating Temporal project directory at $TEMPORAL_DIR..."
    mkdir -p "$TEMPORAL_DIR"
fi

# Navigate to the Temporal directory
cd "$TEMPORAL_DIR"

# Initialize Go module
if [ ! -f "go.mod" ]; then
    echo "Initializing Go module..."
    go mod init temporalpoc
else
    echo "Go module already initialized."
fi

# Install Temporal Go SDK
echo "Installing Temporal Go SDK..."
go get go.temporal.io/sdk

echo "Setup complete! ✅"
