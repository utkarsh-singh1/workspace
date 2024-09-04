#!/bin/sh

echo -en "What is your name - " 
read myname
if [ -z "$myname" ] ; then
    myname=`whoami`
fi
echo "My name is $myname"

