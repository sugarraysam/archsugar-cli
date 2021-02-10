# Archsugar ![Pipeline](https://github.com/sugarraysam/archsugar-cli/workflows/ci/badge.svg?branch=dev) [![Go Report Card](https://goreportcard.com/badge/github.com/sugarraysam/archsugar-cli)](https://goreportcard.com/report/github.com/sugarraysam/archsugar-cli)

## Table of contents

1. [Overview](#overview)
2. [TLDR](#tldr)
3. [Installation from Live ISO](#installation-from-live-iso-recommended)
4. [Packer VM image](#packer-vm-image)
5. [Bootstrap](#bootstrap)
6. [Available Scenarios](#available-scenarios)
7. [Caching Credentials](#caching-credentials)
8. [Roadmap](#roadmap)
9. [Increase size of Vagrant Disk](#increase-size-of-vagrant-disk)

## Overview

This project helps you bootstrap, provision and fully maintain an high-end modular archlinux system

It is highly customizable and relies on [Ansible](https://github.com/ansible/ansible). It is basically managing your localhost with ansible playbooks.

The CLI wrapper around ansible is written in Golang using [Cobra](https://github.com/spf13/cobra).

**This software only supports ArchLinux**

## TLDR

Here is an example workflow:

```bash
# List all scenarios
$ archsugar list

# Run all enabled scenarios
$ archsugar run

# Create a new scenario to manage `lua` programming language - edit the generated templates
$ archsugar create lua --desc "Configure Lua programming language"
$ vim ~/.archsugar/vars/master/lua.yml
> ...edit ansible vars...
$ vim ~/.archsugar/tasks/master/lua.yml
> ...edit ansible tasks...

# View and Run your new scenario
$ archsugar list
$ archsugar run lua

# Enable your new scenario because you like it
$ archsugar enable lua
$ archsugar run # lua scenario is ran

# There is a bug in my scenario, let's disable and investigate
$ archsugar disable lua

# After all I don't like lua :(
$ archsugar rm lua
```

## Installation from Live ISO (Recommended)

**Features**

- UEFI bootloader /w systemd-boot
- Encrypted root partition
- /swapfile
- btrfs filesystem
- Microcode update files for Intel or AMD CPU (auto-detecting)
- Follows [ArchLinux Installation Guide](https://wiki.archlinux.org/index.php/Installation_guide).
- Configures an unpriviledged user named `sugar`

**Prerequisites**

- A computer!! duhh
- Burn the latest [ArchLinux iso](https://www.archlinux.org/download/) on a flash drive
- Boot the computer using the flash drive
- Configure networking (recommended to use Ethernet). Here is one one to do it:

```bash
# Configure your interface
$ export IFACE=ethXX
$ cat >"/etc/systemd/network/${IFACE}.network" <<EOF
[Match]
Name=${IFACE}

[Network]
DHCP=ipv4
EOF
$ systemctl restart systemd-networkd

# Configure DNS
$ sed -i 's,#DNSSEC=.*$,DNSSEC=false,g' /etc/systemd/resolved.conf
$ ln -sf /run/systemd/resolve/resolv.conf /etc/resolv.conf
$ systemctl restart systemd-resolved
```

**Steps**

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

## Packer VM image

TODO

## Bootstrap

Here are the tasks ran by `$ archsugar bootstrap`. They are split in two stages, which are ran sequentially:

**Bootstrap stage**

| Filename         | Description                                                                         |
| ---------------- | ----------------------------------------------------------------------------------- |
| `partition.yml`  | Creates EFI boot partition and root partition.                                      |
| `cryptsetup.yml` | Encrypt and open root partition.                                                    |
| `mkfs.yml`       | Create and mount filesystems, vfat for boot partition and btrfs for root partition. |
| `pacstrap.yml`   | Base archlinux install in chroot and generate fstab.                                |
| `initramfs.yml`  | Install CPU manufacturer microcode, modify mkinitcpio hooks and generate initramfs. |
| `bootloader.yml` | Verify EFI vars are present, install and configure systemd-boot.                    |

**Chroot stage**

| Filename        | Description                                                          |
| --------------- | -------------------------------------------------------------------- |
| `clock.yml`     | Set timezone and hardware clock.                                     |
| `languages.yml` | Configure languages.                                                 |
| `user.yml`      | Set root password, create unpriviledged user /w sudo config.         |
| `pacman.yml`    | Refresh keys, generate mirrors and refresh package database.         |
| `system.yml`    | Create swapfile, configure NetworkManager and other system settings. |
| `archsugar.yml` | Clone archsugar on your new system.                                  |

## Available Scenarios

| NAME      | DESCRIPTION                                                         |
| :-------- | :------------------------------------------------------------------ |
| X         | Install and configure Xorg /w startx                                |
| base      | Configure systemd, firefox, iptables and much more essential sub... |
| blackarch | Add blackarch penetration testing and security research reposito... |
| go        | Configure & clean golang environment                                |
| i3        | Install and configure i3 window manager /w lock screen and statu... |
| k8s       | Install k8s tooling                                                 |
| nodejs    | Configure nodejs and install global packages                        |
| nvim      | Configure neovim with couple of plugins managed by vim-plug         |
| python    | Configure python                                                    |
| qemu      | Configure QEMU/KVM and other vm tools                               |
| ruby      | Configure ruby                                                      |
| rust      | Configure rust                                                      |
| tmux      | Install and configure tmux with dotfile and tmux-plugin-manager     |
| vscode    | install and configure vscode IDE                                    |
| zsh       | Configure zsh and oh-my-zsh /w autosuggestions and custom theme     |

## Caching Credentials

The `$ archsugar run` command will prompt you for your `BECOME password: ...` on every run.
This can be annoying so here is how to cache those credentials using your favorite password manager:

1. **Add your ansible vault password to your favorite password manager. I will be using `gopass`:**

```bash
$ gopass generate -s vault/archsugar
```

2. **Write `vault_password_file.sh` to be a script invoking your password manager:**

```bash
# This files needs to be in the root of your archsugar directory
# !! THE FILENAME NEEDS TO MATCH EXACTLY !!
$ cat <<EOF >> ~/.archsugar/vault_password_file.sh
#!/bin/bash
$ gopass show -o vault/archsugar
EOF
# Make it executable
$ chmod +x vault_password_file.sh
```

3. **Store the `ansible_become_pass` in `group_vars/all.yml` using ansible vault:**

```bash
# Save sudo password
$ mkdir group_vars
$ vim group_vars/all.yml
> ---
> ansible_become_pass: <SUDO_PASSWORD>

# Encrypt - copy generated password to clipboard
$ gopass show -c vault/archsugar
$ ansible-vault encrypt group_vars/all.yml
> New Vault password: <CTRL+V>
> Confirm New Vault password: <CTRL+V>
```

Now, when you `$ archsugar run`, you will be prompted for your password manager's master password, in order to retrieve your ansible-vault password, which will decrypt the file containing your sudo password.

The cool part is that this mechanism will use your password manager's underlying credential cache and stop asking for your sudo password everytime!

## Roadmap

- Create packer VM images for easier testing (QEMU, virtualbox, vmware)
- Provide more explanations on how to fork `archsugar` so can have your own `dotfiles` and personalize your system to a maximum
- Refactor and guidelines
  -- Use `get_url` with `checksum` to detect new version
  -- Download linux binaries from github releases where they are unavailable as Archlinux packages

## Vagrant

TODO

- Set SUGAR_VAGRANT=true in packer/vagrant (no flag for this, not a user feature I want to expose)
- Run install.sh ++ $ archsugar bootstrap inside vagrant
  - archlinux/archlinux box ==> not enough space error (fix this?)
  - terrywang/archlinux box ==> too much stuff in there...
  - Build box from virtualbox-iso then use vagrant cloud post-processor
    - Use this template /w virtualbox-iso ++ archlinux iso
      - https://github.com/elasticdog/packer-arch
    - review bootstrap + chroot stage, add back my old stuff
    - remove vagrant install functionality
- Automate build /w packer
- Push to vagrantcloud
- Add to CI/CD
- Documentation
  - provide sample Vagrantfile (https://gist.github.com/terrywang/6506216)
  - Provide increase disk size explanation

### Increase size of Vagrant disk

- Make sure the vm is stopped `$ vagrant halt`
- Uncomment the `VAGRANT_EXPERIMENTAL=disks` env var and increase the disk size in the `Vagrantfile`:

```ruby
# ...
ENV["VAGRANT_EXPERIMENTAL"] = "disks"
# ...
config.vm.disk :disk, size: "2GB", primary: true
```

- Run `$ vagrant up` and `$ vagrant ssh` to resize the root partition:

```bash
# Delete and recreate root partition - this wont delete the data, only change partition metadata
$ fdisk /dev/sda
> Command (m for help): d
> Partition number (1,2,default 2): 2
> Command (m for help): n
> Partition number (2-128, default 2): 2
> First sector (4096-XXXXXXXX, default 4096): <ENTER>
> Last sector, +/-sectors or +/-size{K,M,G,T,P} (4096-20971486, default XXXXXXX): <ENTER>
# Keed your filesystem signature, preview and write changes
> Do you want to remove the signature? [Y]es/[N]o: N
> Command (m for help): p
> Command (m for help): w
# Your partition should be using the new space now
$ fdisk -l
```

- Still from inside the vm, resize the filesystem and reboot:

```bash
# Filesystem needs to be resized
$ df -hT /dev/sda2
$ btrfs filesystem resize max /
$ reboot
```
