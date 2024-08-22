#!/bin/sh

echo "Provide your name ="
read USER_NAME
echo "	"
echo "********************"
echo "* Hi" $USER_NAME "*"
echo "********************"

echo "I will create a file named ${USER_NAME}_file"

touch "${USER_NAME}_FILE"
