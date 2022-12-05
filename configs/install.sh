#!/bin/sh

umask 0022

mkdir -p /etc/LinuxExecBot/

cp LinuxExecBot /usr/bin/
cp config.yaml /etc/LinuxExecBot/
cp linuxexecbot.service /lib/systemd/system/
cp linuxexecbot@.service /lib/systemd/system/