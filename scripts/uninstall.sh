#!/bin/bash

DEST_DIR="$HOME/linux-docker-web-gui"
SERVICE_NAME="linux-docker-web-gui"

echo "This will remove the installation directory: $DEST_DIR"
echo "And the systemd service: $SERVICE_NAME"
echo "WARNING: All data will be deleted."

read -p "Are you sure you want to continue? (y/N): " -n 1 -r
echo

if [[ $REPLY =~ ^[Yy]$ ]]; then
    # 1. Stop and Remove Service
    if systemctl list-units --full -all | grep -Fq "$SERVICE_NAME.service"; then
        echo "Stopping and disabling service..."
        if [ "$EUID" -ne 0 ]; then SUDO="sudo"; else SUDO=""; fi
        
        $SUDO systemctl stop $SERVICE_NAME
        $SUDO systemctl disable $SERVICE_NAME
        $SUDO rm /etc/systemd/system/$SERVICE_NAME.service
        $SUDO systemctl daemon-reload
        echo "Service removed."
    else
        echo "Service not found or already removed."
    fi

    # 2. Remove Files
    if [ -d "$DEST_DIR" ]; then
        echo "Removing files in $DEST_DIR..."
        rm -rf "$DEST_DIR"
        echo "Files removed."
    else
        echo "Directory not found."
    fi
    
    echo "Uninstallation complete."
else
    echo "Uninstallation cancelled."
fi