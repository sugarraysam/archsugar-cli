#!/bin/bash

set -x
set -e

ARCHSUGAR_VERSION=0.2.1
ARCHSUGAR_URL="https://github.com/sugarraysam/archsugar-cli/releases/download/v${ARCHSUGAR_VERSION}/archsugar-cli_${ARCHSUGAR_VERSION}_linux_amd64.tar.gz"

# Make sure there will be enough space on the live ISO
function resizeRootfs() {
    if [ -d /run/archiso ]; then
        mount -o remount,size=5G /run/archiso/cowspace
    fi
}

function installPacmanDependencies() {
    pacman --noconfirm -q -Sy
    pacman --noconfirm -q -S git ansible
}

function loadKernelModules() {
    # required by cryptsetup
    modprobe dm_mod
}

function setupArchsugar() {
    curl -Lsk -o archsugar.tar.gz "${ARCHSUGAR_URL}"
    tar -xvzf archsugar.tar.gz \
        --exclude="*LICENSE*" \
        --exclude="*README*" \
        --strip-components 1 >/dev/null 2>&1
    chmod 755 archsugar
    cp archsugar /usr/local/bin/archsugar
    rm archsugar.tar.gz archsugar
    archsugar init --branch dev
}

function printInstructions() {
    echo "==> Find your root disk:"
    printf "\tfdisk -l"
    printf "\n\n"

    echo "==> (Optional) Zero your disk. It will take a couple of hours:"
    printf "\tdd if=/dev/zero of=/dev/<disk> bs=4MB status=progress"
    printf "\n\n"

    echo "==> Install archsugar:"
    printf "\tarchsugar bootstrap --disk /dev/<disk> --luksPasswd <...> --rootPasswd <root> --userPasswd <...>"
}

# Execute
resizeRootfs
installPacmanDependencies
loadKernelModules
setupArchsugar
printInstructions
