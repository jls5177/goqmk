# QMK CLI Toolbox
This repo houses a command line utility to interact with QMK Keyboards with the Via interface
enabled.

## Features
1. Via lighting commands to set the following (without saving to EEPROM):
    1. Color (Hue and Saturation)
    1. Lighting Effect
    1. Lighting Effect Speed
    1. Brightness
    1. Save current settings to EEPROM
1. Flash command that wraps the [wally-cli](https://github.com/zsa/wally-cli) package (supports all ZSA boards)
