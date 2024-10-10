#!/bin/bash
sudo apt remove docker-desktop
rm -r $HOME/.docker/desktop
sudo rm /usr/local/bin/com.docker.cli
sudo apt purge  -y docker-desktop
sudo apt install  -y gnome-terminal
sudo curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | apt-key add -
sudo add-apt-repository "deb [arch=amd64] http://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable"
apt -y update
sudo apt install -y docker-ce
sudo echo '{ "registry-mirrors": ["https://ra9q5uam.mirror.aliyuncs.com"] }' > /etc/docker/daemon.json
sudo service docker restart