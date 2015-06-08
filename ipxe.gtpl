#!ipxe
set base-url http://{{.BaseUrl}}
kernel ${base-url}/coreos_production_pxe.vmlinuz sshkey="{{.SSHKey}}"
initrd ${base-url}/coreos_production_pxe_image.cpio.gz
boot
