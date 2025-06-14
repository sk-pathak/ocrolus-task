#!/bin/bash

usage() {
    echo "Usage: $0 <migration_name>"
    echo
    echo "Example:"
    echo "  $0 add_users_table"
    exit 1
}

create_migration() {
    local name="$1"

    if [ -z "$name" ]; then
        echo "Error: Migration name is required."
        usage
    fi

    echo "Creating new migration: $name"
    goose create "$name" sql

    if [ $? -eq 0 ]; then
        echo "Migration created successfully."
    else
        echo "Error creating migration."
        exit 1
    fi
}

if [ $# -ne 1 ]; then
    usage
else
    create_migration "$1"
fi
