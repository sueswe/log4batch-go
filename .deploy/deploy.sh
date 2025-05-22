#!/bin/bash

source ~/.profile

cd "$HOME"/compile/log4batch-go || {
    echo "Status: $?"
    exit 4
}

echo "------------------------------------"
env | grep PATH
env | grep LOADED
echo "------------------------------------"


echo '
### LINUX #########################################################
'

stages="
stp,testta3
lgkk,testta3
"

for UMG in ${stages}
do
    cd /tmp/ || exit 1
    "$HOME"/bin/vicecersa.sh "${UMG}" log4batch \$HOME/bin/ || {
        echo "Status: $?"
        exit 2
    }
done


echo '
### AIX #########################################################
'

stages="
stp,testta2
lgkk,testta2
"

for UMG in ${stages}
do
    cd /tmp/ || exit 1
    "$HOME"/bin/vicecersa.sh "${UMG}" log4batch.aix \$HOME/bin/ log4batch || {
        echo "Status: $?"
        exit 2
    }
done

