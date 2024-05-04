#!/bin/bash

if [ "$(id -u)" -ne 0 ]; then
    echo "This installtion process need sudo priviliges in order to install IntelliPath."
    echo "Please run the installation script as sudo."
fi

mkdir -p /usr/local/bin/intellipath
chmod 777 /usr/local/bin/intellipath

# tar -xzf intellipath.tar.gz -C /usr/local/bin/intellipath

echo "Installation complete."
echo "Run: 'Intellipath init' in order to initialize Intallipath on your machine"
