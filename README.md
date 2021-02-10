# Archsugar ![Pipeline](https://github.com/sugarraysam/archsugar-cli/workflows/ci/badge.svg?branch=dev) [![Go Report Card](https://goreportcard.com/badge/github.com/sugarraysam/archsugar-cli)](https://goreportcard.com/report/github.com/sugarraysam/archsugar-cli) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=sugarraysam_archsugar-cli&metric=alert_status)](https://sonarcloud.io/dashboard?id=sugarraysam_archsugar-cli)


## Description

This project helps you bootstrap, provision and fully maintain an high-end modular archlinux system. It helps you maintain your `Workstation as Code`.

It is highly customizable and relies on [Ansible](https://github.com/ansible/ansible) playbooks to manage your localhost.

This project is the CLI wrapper which uses [Cobra](https://github.com/spf13/cobra) and by default my [ansible dotfiles](https://github.com/sugarraysam/archsugar).

**This software only supports ArchLinux**

## Getting Started

Here is an example workflow:

```bash
# Initialize ansible dotfiles in ~/.archsugar
$ archsugar init

# Enable, list and run all scenarios
$ archsugar enable --all
$ archsugar list
$ archsugar run
```

## Documentation

- [How to install from live ArchLinux ISO](docs/installation_from_iso.md)
- [Explanation of Bootstrap, Chroot and Master stages](docs/stages.md)
- [HashiCorp Packer vagrant builder documentation](docs/packer.md)
- [Caching Credentials for the Master stage using your local password manager](docs/caching_credentials.md)
