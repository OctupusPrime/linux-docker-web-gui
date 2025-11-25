# Installation Instructions

This project includes automated scripts to install and manage the Linux Docker Web GUI on your Ubuntu/Linux server.

## Prerequisites

- A Linux server (Ubuntu recommended)
- `curl` installed (usually present by default)
- Root or `sudo` access

## Install

You can install the application by running the `install.sh` script.

**Download the installation script:**

```bash
curl -fsSL https://raw.githubusercontent.com/OctupusPrime/linux-docker-web-gui/master/install.sh | sudo bash
```

## Uninstall

You can unistall the applcation by running included `uinstall.sh` scrpt.

```bash
cd ~/linux-docker-web-gui && sudo ./uninstall.sh
```

### What the installer does:

- Installs necessary system dependencies (`unzip`, `jq`, `openssl`).
- Downloads the latest release from GitHub.
- Unzips it to `~/linux-docker-web-gui`.
- Generates a `.env` file with secure secrets (`JWT_SECRET`, `PASSWORD_SECRET`) if one doesn't exist.
- Sets up and starts a systemd service named `linux-docker-web-gui` so the app starts on boot.

---

## Managing the Service

After installation, the application runs automatically in the background.

- **Check Status:**
  ```bash
  sudo systemctl status linux-docker-web-gui
  ```
