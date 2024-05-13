#!/bin/bash


EXECUTABLE_NAME="intellipath"
INSTALL_DIR="/usr/local/bin/$EXECUTABLE_NAME"
TAR_FILE="./$EXECUTABLE_NAME.tar.gz"


print_error() {
    echo -e "\033[0;31mERROR: $1\033[0m"
}

# Function to cleanup on installation failure
cleanup() {
    # Remove installed files
    sudo rm -rf "$INSTALL_DIR"
    print_error "Installation aborted. Cleaned up installed files."
    exit 1
}

if [ "$(id -u)" -ne 0 ]; then
    print_error "Please run this script as root (sudo)."
    exit 1
fi

sudo mkdir -p "$INSTALL_DIR" || { print_error "Failed to create installation directory."; exit 1; }
chmod 777 "$INSTALL_DIR"

# Install executable
sudo tar -xzf "$TAR_FILE" -C "$INSTALL_DIR" || { print_error "Failed to install executable."; cleanup; }

# Set permissions
sudo chmod +x "$INSTALL_DIR/$EXECUTABLE_NAME" || { print_error "Failed to set permissions."; cleanup; }

# Prompt user to modify bashrc/zshrc
read -p "Do you want to add $INSTALL_DIR to your PATH in .bashrc/.zshrc? (y/n): " add_to_path
if [ "$add_to_path" == "y" ]; then
    echo "You picked Yes"
    echo "export PATH=\"$INSTALL_DIR:\$PATH\"" >> ~/.bashrc || { print_error "Failed to modify .bashrc."; cleanup; }
    source ~/.bashrc
else
    cleanup
fi

# Prompt user to add alias
read -p "Do you want to add 'icd' alias for the executable? (y/n): " add_alias
if [ "$add_alias" == "y" ]; then
    echo "You picked Yes"
    echo "alias icd=\"$INSTALL_DIR/$EXECUTABLE_NAME icd\"" >> ~/.bashrc || { print_error "Failed to add alias."; cleanup; }
    source ~/.bashrc
else
    cleanup
fi

# Run intellipath init
intellipath init

# Check if database file exists
if [ ! -f "$INSTALL_DIR/ipath.db" ]; then
    print_error "Database file not found."
    cleanup
fi

echo "Installation completed successfully."
exit 0
