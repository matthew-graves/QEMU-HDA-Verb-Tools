# QEMU HDA Verb Tools
Utilities for Debugging Intel High Definition Audio (HDA) Verbs

Meant to be used with the QEMU Instructions and Patch Found Here: https://jcs.org/2018/11/12/vfio ; https://github.com/jcs/qemu/compare/e22f675b..jcs-hda-dma

Logs for use with these utilities can be captured with `./startvm.sh 2>&1 | grep "CORB caddr" >> commands.txt`
Scripts to Start Windows 11 VMs can be found at https://github.com/matthew-graves/Windows11-QEMU-Scripts