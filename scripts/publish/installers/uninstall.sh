# To uninstall:
# curl -o- https://dl.stackstate.com/stackstate-cli/uninstall.sh | bash
#-----------------------------------
# Parameters to script
#-----------------------------------
# STS_CLI_LOCATION - Path you want to install CLI (empty means `/usr/local/bin`)
#-----------------------------------

#!/usr/bin/env bash

FILE=/usr/local/bin/sts
# Check if sts is in default location
if [[ -f "$FILE" ]]; then
    sudo rm -r "$FILE" ~/.config/stackstate-cli
    echo "Uninstalled the StackState CLI. See you later!"
# check if custome location was provided in the script
elif [[ ! -z "$STS_CLI_LOCATION" ]]; then
    CUSTOM_CLI_LOCATION="$STS_CLI_LOCATION"/sts 
    # check if there is a sts the location provided
    if [[ -f "$CUSTOM_CLI_LOCATION" ]]; then
        # sts found at the custom location
        rm -r "$CUSTOM_CLI_LOCATION" ~/.config/stackstate-cli
        echo "Uninstalled the StackState CLI. See you later!"
    else
        # sts not found at the custom location
        echo "CLI not found in the location provided $CUSTOM_CLI_LOCATION ."
    fi 
else
    echo "Couldn't find StackState CLI in the default location."
    echo "You can provide a custom location with STS_CLI_LOCATION="
fi