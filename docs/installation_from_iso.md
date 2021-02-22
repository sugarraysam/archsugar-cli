# Installation from Live ISO (Recommended)

## Features

- UEFI bootloader /w systemd-boot
- Encrypted root partition
- /swapfile
- btrfs filesystem
- Microcode update files for Intel or AMD CPU (auto-detecting)
- Follows [ArchLinux Installation Guide](https://wiki.archlinux.org/index.php/Installation_guide).
- Configures an unpriviledged user named `sugar`

## Prerequisites

- Burn the latest [ArchLinux iso](https://www.archlinux.org/download/) on a flash drive
- Boot the computer using the flash drive
  a. Disable Secure Boot
  b. Disable corporate rootkits
- Configure networking (recommended to use Ethernet). Here is one one to do it:

### Configure network interface

**1. Ethernet**

```bash
$ export IFACE=ethXX
$ cat >"/etc/systemd/network/00-${IFACE}.network" <<EOF
[Match]
Name=${IFACE}

[Network]
DHCP=ipv4
EOF

$ systemctl restart systemd-networkd
```

**2. Wifi**

```bash
$ export SSID=<ssid>
$ export PASSWD=<passwd>
$ export IFACE=wlanX

$ wpa_passphrase ${SSID} ${PASSWD} | tee /etc/wpa_supplicant/wpa_supplicant-${IFACE}.conf
$ cat >"/etc/systemd/network/00-${IFACE}.network" <<EOF
[Match]
Name=${IFACE}

[Network]
DHCP=ipv4
EOF

$ systemctl restart systemd-networkd
$ systemctl restart wpa_supplicant@${IFACE}.conf
```

**Configure DNS**

```bash
$ sed -i 's,#DNSSEC=.*$,DNSSEC=false,g' /etc/systemd/resolved.conf
$ ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
$ systemctl restart systemd-resolved
```

## Steps

```bash
# As root user
$ curl https://github.com/sugarraysam/archsugar/blob/master/install.sh | bash

# The repo should be present in /root/.archsugar
$ ls /root/.archsugar

# The CLI should be installed
$ archsugar --help

# Bootstrap your machine
$ archsugar bootstrap --disk /dev/sda --luksPasswd luks --rootPasswd root --userPasswd user

# Reboot
$ reboot

# Log back in (provide luksPasswd to decrypt disk, then should autologin as sugar)

# Configure your new system with archsugar (enable all scenarios)
$ archsugar list
$ archsugar enable --all
$ archsugar run
$ reboot
```
