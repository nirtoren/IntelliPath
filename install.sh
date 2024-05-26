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

add_line_if_not_exists() {
    local line="$1"
    local file="$2"
    grep -qxF "$line" "$file" || echo "$line" >> "$file"
}

if [ "$(id -u)" -ne 0 ]; then
    print_error "Please run this script as root (sudo)."
    exit 1
fi

USER_HOME=$(eval echo ~${SUDO_USER})
BASHRC="$USER_HOME/.bashrc"

sudo mkdir -p "$INSTALL_DIR" || { print_error "Failed to create installation directory."; exit 1; }
chmod 777 "$INSTALL_DIR"

# Install executable
sudo tar -xzf "$TAR_FILE" -C "$INSTALL_DIR" || { print_error "Failed to install executable."; cleanup; }

# Set permissions
sudo chmod +x "$INSTALL_DIR/$EXECUTABLE_NAME" || { print_error "Failed to set permissions."; cleanup; }

# Prompt user to modify bashrc/zshrc
read -p "Do you want to add $INSTALL_DIR to your PATH in .bashrc? (y/n): " add_to_path
if [ "$add_to_path" == "y" ]; then
    path_line="export PATH=\"$INSTALL_DIR:\$PATH\""
    add_line_if_not_exists "$path_line" $BASHRC || { print_error "Failed to modify .bashrc."; cleanup; }
    source $BASHRC
else
    cleanup
fi

read -p "Do you want to add 'icd' function to your .bashrc? (y/n): " add_function
if [ "$add_function" == "y" ]; then
    icd_function="icd(){
    dir=\$($INSTALL_DIR/$EXECUTABLE_NAME \"icd\" \"\$1\")
    cd \"\$dir\"
}"
    add_line_if_not_exists "$icd_function" $BASHRC || { print_error "Failed to add function to .bashrc."; cleanup; }
    source $BASHRC
else
    cleanup
fi

echo "Installation completed successfully."
exit 0