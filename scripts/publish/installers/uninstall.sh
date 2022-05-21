# To uninstall:
# curl -o- https://dl.stackstate.com/stackstate-cli/install.ps1 |STS_API_URL="url" STS_API_TOKEN="token" bash

#!/usr/bin/env bash
rm -r /usr/local/bin/sts ~/.config/stackstate-cli
echo "Uninstalled the StackState CLI. See you later!"