# Kubernetes The Hard Way

Learning Kubernetes based on the tutorial [Kubernetes The Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way).

- [Ubuntu 24.04 LTS](https://ubuntu.com/) or [Pop!_OS 22.04 LTS](https://pop.system76.com/)
- [VMware Workstation Pro](https://support.broadcom.com/group/ecx/productdownloads?subfamily=VMware+Workstation+Pro)
- [VMware Host Modules](https://github.com/mkubecek/vmware-host-modules)
- [Vagrant](https://www.vagrantup.com/)
- [Vagrant VMware Utility](https://developer.hashicorp.com/vagrant/docs/providers/vmware/vagrant-vmware-utility)
- [Ansible](https://www.ansible.com/)

## Install VMware Workstation Pro

```bash
sudo chmod +x VMware-Workstation-Full-17.5.2-23775571.x86_64.bundle
sudo ./VMware-Workstation-Full-17.5.2-23775571.x86_64.bundle
```

## Install VMware Host Modules

```bash
wget https://github.com/mkubecek/vmware-host-modules/archive/workstation-17.5.1.tar.gz
tar -xzf workstation-17.5.1.tar.gz
cd vmware-host-modules-workstation-17.5.1
make
make install
```

## Install Vagrant

```bash
sudo apt update
sudo apt install vagrant
vagrant plugin install vagrant-vbguest
vagrant plugin install --plugin-version=3.0.1 vagrant-vmware-desktop
```

## Install Vagrant VMware Utility

```bash
curl -O https://releases.hashicorp.com/vagrant-vmware-utility/1.0.22/vagrant-vmware-utility_1.0.22-1_amd64.deb
sudo dpkg -i vagrant-vmware-utility_1.0.22-1_amd64.deb
```

## Install Ansible

```bash
sudo apt update
sudo apt install ansible
```

## Vagrant commands

```bash
vagrant up
vagrant halt
vagrant destroy
vagrant status
```
