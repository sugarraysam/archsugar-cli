# Caching Credentials

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
