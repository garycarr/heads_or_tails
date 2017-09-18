# This would need to do coverage, migrations etc
mkdir -p $BUILD_DIR
rm -f $BUILD_DIR/$APP_NAME
ln -s $PWD $BUILD_DIR/$APP_NAME
export GOPATH=$PWD/_build
export PATH=$GOPATH/bin:$PATH
export PATH=$PATH:/home/gcarr/go/bin:/usr/local/go/bin:
cd $BUILD_DIR/$APP_NAME
go build
go test
