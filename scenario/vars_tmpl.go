package scenario

var VarsTmpl = `---
# - <pkg_name> # Description
(( .Name ))_pacman_pkgs: []

(( .Name ))_dirs_user: []

(( .Name ))_dirs_root: []

# { src: <src>, dest: <dest> }
(( .Name ))_dotfiles_user: []

# { src: <src>, dest: <dest> }
(( .Name ))_dotfiles_root: []

# Can include both files and directories
(( .Name ))_files_to_remove: []

# Download and extract binaries from url archive
# use '%%; token inside urlfmt, will be replaced with item.version
# { name: <binary_name>, version: <version>, urlfmt: <url> }
(( .Name ))_remote_bin_archives: []

# Download binary that is not an archive
# use '%%; token inside urlfmt, will be replaced with item.version
# { name: <binary_name>, version: <version>, urlfmt: <url> }
(( .Name ))_remote_bin: []

# Clone git repos
# { repo: <repo>, dest: <dest>, version: <branch/tag> }
(( .Name ))_git_repos: []

# Output zsh completion commands to /usr/share/zsh/site-functions/
# { cmd: <repo>, dest: <dest> }
(( .Name ))_zsh_completions: []
`
