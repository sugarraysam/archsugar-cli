<!-- no toc -->
# Stages

- [Stages](#stages)
  - [Bootstrap and Chroot](#bootstrap-and-chroot)
  - [Master](#master)

There are three playbook stages provided with the default dotfiles. The `bootstrap` and `chroot` stages are only used during the initial installation. The `chroot` stage is ran under the cover and you do not have control over it. The `master` stage is the one you will use in your day-to-day operations.

## Bootstrap and Chroot

When first installing ArchLinux with archsugar, you need to run `$ archsugar bootstrap`. This command runs both the `bootstrap` and `chroot` stages sequentially:

**Bootstrap tasks:**

| Filename         | Description                                                                         |
| ---------------- | ----------------------------------------------------------------------------------- |
| `partition.yml`  | Creates EFI boot partition and root partition.                                      |
| `cryptsetup.yml` | Encrypt and open root partition.                                                    |
| `mkfs.yml`       | Create and mount filesystems, vfat for boot partition and btrfs for root partition. |
| `pacstrap.yml`   | Base archlinux install in chroot and generate fstab.                                |
| `initramfs.yml`  | Install CPU manufacturer microcode, modify mkinitcpio hooks and generate initramfs. |
| `bootloader.yml` | Verify EFI vars are present, install and configure systemd-boot.                    |


**Chroot tasks:**

| Filename        | Description                                                          |
| --------------- | -------------------------------------------------------------------- |
| `clock.yml`     | Set timezone and hardware clock.                                     |
| `languages.yml` | Configure languages.                                                 |
| `user.yml`      | Set root password, create unpriviledged user /w sudo config.         |
| `pacman.yml`    | Refresh keys, generate mirrors and refresh package database.         |
| `system.yml`    | Create swapfile, configure NetworkManager and other system settings. |
| `archsugar.yml` | Clone archsugar on your new system.                                  |

## Master

This is the stage for day-to-day operations. Whenever you create/rm, enable/disable or edit a scenario, you are modifying the `master` stage.

Here are the available scenarios from the [default archsugar dotfiles](https://github.com/sugarraysam/archsugar):

| NAME      | DESCRIPTION                                                           |
| :-------- | :-------------------------------------------------------------------- |
| X         | Install and configure Xorg /w startx                                  |
| backup    | __Not implemented, no-op at the moment__                              |
| base      | Configure systemd, autologin, git, gpg, etc.                          |
| blackarch | Add blackarch penetration testing and security research repositories  |
| cpp       | Configure and manage c++ coding ecosystem                             |
| docker    | Manage local docker daemon                                            |
| firefox   | Setup firefox profile                                                 |
| go        | Configure & clean golang environment                                  |
| i3        | Install and configure i3 window manager /w lock screen and status bar |
| iptables  | Configure iptables ipv4 and ipv6                                      |
| k8s       | Install k8s tooling                                                   |
| nodejs    | Configure nodejs and install global packages                          |
| nvim      | Configure neovim with couple of plugins managed by vim-plug           |
| pacman    | Manage pacman configs and packages                                    |
| python    | Configure python                                                      |
| ruby      | Configure ruby                                                        |
| rust      | Configure rust                                                        |
| tmux      | Install and configure tmux with dotfile and tmux-plugin-manager       |
| vbox      | Configure virtualbox and related tools                                |
| vscode    | install and configure vscode IDE                                      |
| zsh       | Configure zsh and oh-my-zsh /w autosuggestions and custom theme       |
