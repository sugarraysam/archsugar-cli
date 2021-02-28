package scenario

var TasksTmpl = `# (( .Desc ))
---
- include_vars: master/(( .Name )).yml

- include_role:
    name: scenario
  vars:
    pacman_pkgs: "{{ (( .Name ))_pacman_pkgs }}"
    dirs_user: "{{ (( .Name ))_dirs_user }}"
    dirs_root: "{{ (( .Name ))_dirs_root }}"
    dotfiles_user: "{{ (( .Name ))_dotfiles_user }}"
    dotfiles_root: "{{ (( .Name ))_dotfiles_root }}"
    files_to_remove: "{{ (( .Name ))_files_to_remove }}"
    remote_bin: "{{ (( .Name ))_remote_bin }}"
    remote_bin_archives: "{{ (( .Name ))_remote_bin_archives }}"
    git_repos: "{{ (( .Name ))_git_repos }}"
	zsh_completions: "{{ (( .Name ))_zsh_completions }}"

- name: User commands
  block:
    - name: Smoke test
      debug:
        msg: Smoke test

  become: true
  become_user: "{{ user }}"
`
