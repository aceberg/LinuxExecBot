#!/bin/bash

cd cmd/LinuxExecBot/ && CGO_ENABLED=0 go build -o ../../LinuxExecBot .
cd ../../

PKGDIR=/opt/LinuxExecBot

umask 0022

mkdir -p $PKGDIR
cp LinuxExecBot $PKGDIR/
cp configs/config.yaml $PKGDIR/
cp configs/linuxexecbot.service $PKGDIR/
cp configs/linuxexecbot@.service $PKGDIR/
cp configs/install.sh $PKGDIR/

cd /opt
tar cvzf LinuxExecBot-$1.tar.gz LinuxExecBot
mv LinuxExecBot-$1.tar.gz ../