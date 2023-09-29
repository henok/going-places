#!/bin/bash

# Check if the argument is provided
if [ "$#" -ne 1 ]; then
    echo "Error: Project directory not provided."
    echo "Usage: $0 <project_directory_name>"
    exit 1
fi

# Use the provided argument as the project directory name
PROJECT_DIR="$1"

# Create the main project directory
mkdir $PROJECT_DIR

# Check if directory creation was successful
if [ $? -ne 0 ]; then
    echo "Error: Failed to create the project directory. It might already exist or there could be another issue."
    exit 2
fi

# Navigate into the project directory
cd $PROJECT_DIR

# Create the main directories
mkdir api db models utils config

# Create sub-directories for the api directory
mkdir api/handlers api/middlewares api/routes

# Create files (optional, just as a starting point)
touch main.go
touch api/handlers/word.go
touch api/middlewares/some_middleware.go
touch api/routes/routes.go
touch db/redis.go
touch models/some_model.go
touch utils/some_util.go
touch config/config.go

echo "Current directory: $(pwd)"
# Print out the structure (optional)
/opt/homebrew/bin/tree . 

echo "Project structure created successfully!"

