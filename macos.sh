go build gomit.go
mkdir /usr/local/gomit
cp -f gomit /usr/local/gomit/gomit
sudo echo "export PATH=$PATH:/usr/local/gomit" >> /etc/profile
chmod +x /usr/local/gomit/gomit