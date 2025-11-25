#!/bin/bash

# Exit on error
set -e

# 1. OS Check
if [[ "$(uname -s)" != "Linux" ]]; then
    echo "Error: This script is intended for Linux only."
    exit 1
fi

REPO="OctupusPrime/linux-docker-web-gui"
DEST_DIR="$HOME/linux-docker-web-gui"
SERVICE_NAME="linux-docker-web-gui"

echo "Starting installation for $REPO..."
echo "Target directory: $DEST_DIR"

# 2. Install dependencies
echo "Installing necessary tools..."
if [ "$EUID" -ne 0 ]; then 
    SUDO="sudo"
else
    SUDO=""
fi

$SUDO apt-get update -qq
$SUDO apt-get install -y -qq curl unzip jq openssl

# 3. Get latest release URL
echo "Fetching latest release information..."
LATEST_RELEASE_URL=$(curl -s https://api.github.com/repos/$REPO/releases/latest | jq -r '.assets[0].browser_download_url')

if [ "$LATEST_RELEASE_URL" == "null" ] || [ -z "$LATEST_RELEASE_URL" ]; then
    echo "Error: No release found or the release has no assets."
    exit 1
fi

# 4. Download and Unzip
echo "Downloading latest release..."
curl -L -o release.zip "$LATEST_RELEASE_URL"

echo "Unzipping to '$DEST_DIR'..."
unzip -o release.zip -d "$DEST_DIR"
rm release.zip

# 5. Handle .env file
ENV_PATH="$DEST_DIR/.env"

if [ ! -f "$ENV_PATH" ]; then
    echo ".env not found. Creating one..."
    JWT_SECRET=$(openssl rand -hex 32)
    PASSWORD_SECRET=$(openssl rand -hex 32)
    
    cat <<EOT >> "$ENV_PATH"
JWT_SECRET=$JWT_SECRET
PASSWORD_SECRET=$PASSWORD_SECRET
DATABASE_PATH=$DEST_DIR/database.db
FRONTEND_PATH=$DEST_DIR/dist
EOT
    echo ".env created."
else
    echo ".env already exists. Skipping creation."
fi

# 6. Setup Autorun (Systemd)
APP_SCRIPT="$DEST_DIR/app.sh"

if [ -f "$APP_SCRIPT" ]; then
    echo "Setting up systemd service..."
    chmod +x "$APP_SCRIPT"

    # Determine current user and group for the service
    CURRENT_USER=$(whoami)
    
    # Create service file
    # We use sudo bash -c to handle the write permission to /etc/systemd/system
    $SUDO bash -c "cat <<EOF > /etc/systemd/system/$SERVICE_NAME.service
[Unit]
Description=Linux Docker Web GUI
After=network.target

[Service]
Type=simple
User=$CURRENT_USER
WorkingDirectory=$DEST_DIR
ExecStart=$APP_SCRIPT
Restart=on-failure
EnvironmentFile=$ENV_PATH

[Install]
WantedBy=multi-user.target
EOF"

    # Reload, Enable, and Start
    $SUDO systemctl daemon-reload
    $SUDO systemctl enable $SERVICE_NAME
    $SUDO systemctl start $SERVICE_NAME
    
    echo "Service '$SERVICE_NAME' has been installed and started."
    echo "You can check status with: sudo systemctl status $SERVICE_NAME"
else
    echo "Warning: app.sh not found. Service could not be setup."
fi

echo "Installation complete!"