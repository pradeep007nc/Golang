#!/bin/bash

# Source directory containing files without a specific extension
source_dir="$PWD"

# Destination directory for files without a specific extension
destination_dir="exec_files"

# Ensure the destination directory exists
mkdir -p "$destination_dir"

rm -r "exec_files/test.out"

# Move files without a specific extension to the destination directory
find "$source_dir" -type f -name "*.out" -exec mv {} "$destination_dir" \;
