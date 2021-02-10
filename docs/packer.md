# Packer VM image

**TODO**


## Vagrant

**TODO**

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
