PROJECTPATH=$GOPATH/src/bitbucket.org/mundipagg/

rm -rf PROJECTPATH -v

mkdir $PROJECTPATH

mv ~/myagent/_work/1/s -t $PROJECTPATH -v 

mv $PROJECTPATH/s $PROJECTPATH/boletoapi -v

cd $PROJECTPATH

glide install

go build -v

go test $(go list ./... | grep -v /vendor/) -v