[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/LinuxExecBot)](https://goreportcard.com/report/github.com/aceberg/LinuxExecBot)
[![Binary-release](https://github.com/aceberg/LinuxExecBot/actions/workflows/release.yml/badge.svg)](https://github.com/aceberg/LinuxExecBot/actions/workflows/release.yml)
![GitHub](https://img.shields.io/github/license/aceberg/LinuxExecBot)

# LinuxExecBot
Telegram bot to execute a command from a configurable list on your Linux machine

## Install

### From .deb repository (recommended)
```sh
curl -s --compressed "https://aceberg.github.io/ppa/KEY.gpg" | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/aceberg.gpg
sudo curl -s --compressed -o /etc/apt/sources.list.d/aceberg.list "https://aceberg.github.io/ppa/aceberg.list"
sudo apt update
sudo apt install linuxexecbot
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

## Usage
### As root
1. Edit `/etc/LinuxExecBot/config.yaml`, put your Telegram bot's Token and ChatID there. Add your commands. Example:

```yaml
config:
  token: 5XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  id: 111111111

commands:
- name: la
  exec: "ls -al"

- name: ping
  exec: "ping -c 3 8.8.8.8"
```
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

## Options
| Key  | Description | Default |
| --------  | ----------- | ------- |
| -c | Path to config yaml file |config.yaml|


## Support

If you like this app, please support me
- BTC: `bc1qj59rxmfvanvqqltq9t73qls4su3xrvwuv3sxhr`
- ETH: `0x276124c218aa8110F96989AA1f6f2Bb960C234B7`
- USDT(ETH Network): `0x276124c218aa8110F96989AA1f6f2Bb960C234B7`