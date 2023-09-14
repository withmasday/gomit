go build gomit.go
mkdir ~/.gomit
cp -f gomit ~/.gomit
eval sudo chmod u+x ~/.gomit/gomit && sudo ln ~/.gomit/gomit /usr/bin/gomit