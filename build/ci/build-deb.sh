#!/bin/bash

PKGDIR=LinuxExecBot-$1-0_all

mkdir -p $PKGDIR/usr/bin
chmod 755 LinuxExecBot
cp LinuxExecBot $PKGDIR/usr/bin/

mkdir -p $PKGDIR/DEBIAN

echo "Package: linuxexecbot
Version: $1
Section: utils
Priority: optional
Architecture: all
Maintainer: aceberg <aceberg_a@proton.me>
Description: Telegram bot to execute a command from a configurable list on your Linux machine
" > $PKGDIR/DEBIAN/control

dpkg-deb --build --root-owner-group $PKGDIR

rm -rf $PKGDIR