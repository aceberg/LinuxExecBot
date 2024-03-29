[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/LinuxExecBot)](https://goreportcard.com/report/github.com/aceberg/LinuxExecBot)
[![Maintainability](https://api.codeclimate.com/v1/badges/37a8c60c47f41a54c493/maintainability)](https://codeclimate.com/github/aceberg/LinuxExecBot/maintainability)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/099e3be49e80472bb702b7887e620ab1)](https://app.codacy.com/gh/aceberg/LinuxExecBot/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Binary-release](https://github.com/aceberg/LinuxExecBot/actions/workflows/release.yml/badge.svg)](https://github.com/aceberg/LinuxExecBot/actions/workflows/release.yml)
![GitHub](https://img.shields.io/github/license/aceberg/LinuxExecBot)

# LinuxExecBot
Telegram bot to execute a command from a configurable list on your Linux machine

[1. Installation](https://github.com/aceberg/LinuxExecBot#1-installation)   
[2. Usage](https://github.com/aceberg/LinuxExecBot#2-usage)   
[3. Options](https://github.com/aceberg/LinuxExecBot#3-options)   

## 1. Installation

### From .deb repository (recommended)
```sh
curl -s --compressed "https://aceberg.github.io/ppa/KEY.gpg" | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/aceberg.gpg
```
```sh
sudo curl -s --compressed -o /etc/apt/sources.list.d/aceberg.list "https://aceberg.github.io/ppa/aceberg.list"
```
```sh
sudo apt update && sudo apt install linuxexecbot
```
### From .deb file
Download [latest](https://github.com/aceberg/LinuxExecBot/releases/latest) release, install with your package maneger

### From .tar.gz
Download [latest](https://github.com/aceberg/LinuxExecBot/releases/latest) release, then
```sh
tar xvzf LinuxExecBot-*.tar.gz
cd LinuxExecBot
sudo ./install.sh
```
### With go
```sh
go install github.com/aceberg/LinuxExecBot/cmd/LinuxExecBot@latest
```

## 2. Usage
### As root
1. Edit `/etc/LinuxExecBot/config.yaml`, put your Telegram bot's Token and ChatID there. Add your commands. Example:

```yaml
config:
  token: 5XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  #id: 111111111 
  ids:     # Use either id (for one user) or ids (for several bot users)
    - 111111111
    - 222222222
  args: no # If set to "yes", allows to send arguments with bot commands
           # WARNING! Could be dangerous, if someone gets access to your bot
commands:
- name: la
  exec: "ls -al"
  desc: "List all files"

- name: ping
  exec: "ping -c 3 8.8.8.8"
  desc: "Ping Google DNS"
```
> Only commands with not empty `desc` field will be shown in Bot Menu.   
> `/help` command lists all available commands

2. Enable and start service
```sh
systemctl enable linuxexecbot.service
systemctl start linuxexecbot.service
```
3. Go to your bot in Telegram and try to execute command. Example: `/la`

### As user
1. Copy config example
```sh
mkdir -p ~/.config/LinuxExecBot
cp /etc/LinuxExecBot/config.yaml ~/.config/LinuxExecBot/config.yaml
```
2. Edit `config.yaml`,  put your bot's Token and ChatID there
3. Enable and start service, replace `MYUSER` with your username
```sh
systemctl enable linuxexecbot@MYUSER.service
systemctl start linuxexecbot@MYUSER.service
```
4. Go to your bot in Telegram and try to execute command.

## 3. Options
| Key  | Description | Default |
| --------  | ----------- | ------- |
| -c | Path to config yaml file |config.yaml|