#!/bin/sh

EXEC=$1
cd ../ksctl-components/ || exit 1

echo "--------------------------------------------"
echo "|   Testing (ksctl-components/stateimport)"
echo "--------------------------------------------"

cd stateimport
GOTEST_PALETTE="red,yellow,green" $EXEC . -v && cd -

