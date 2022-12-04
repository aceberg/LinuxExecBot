#!/bin/bash

PKGDIR=LinuxExecBot-$1-0_all

umask 0022

mkdir -p $PKGDIR/usr/bin
mkdir -p $PKGDIR/lib/systemd/system
mkdir -p $PKGDIR/etc/LinuxExecBot

cp configs/linuxexecbot.service $PKGDIR/lib/systemd/system/
cp configs/linuxexecbot@.service $PKGDIR/lib/systemd/system/
cp configs/config.yaml $PKGDIR/etc/LinuxExecBot/

cp LinuxExecBot $PKGDIR/usr/bin/

mkdir -p $PKGDIR/DEBIAN

echo "Package: linuxexecbot
Version: $1
Section: utils
Priority: optional
Architecture: all
Maintainer: aceberg <aceberg_a@proton.me>
Description: Telegram bot to execute a command from a configurable list
" > $PKGDIR/DEBIAN/control

echo "/etc/LinuxExecBot/config.yaml
" > $PKGDIR/DEBIAN/conffiles

dpkg-deb --build --root-owner-group $PKGDIR

rm -rf $PKGDIR