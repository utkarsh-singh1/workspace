#!/bin/sh

echo "Provide your name ="
read USER_NAME
echo "	"
echo "********************"
echo "* Hi" $USER_NAME "*"
echo "********************"


# {} is to insure that from variable is started and where it has ended.
echo "I will create a file named ${USER_NAME}_file"

# Difference between using touch command to create file using qoutes around ${USER_NAME}_FILE can be seen by execution try it by yourself
touch "${USER_NAME}_FILE"
