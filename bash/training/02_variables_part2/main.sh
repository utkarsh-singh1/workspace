#!/bin/sh

[ $# -le 1 ] && echo "I was called by $# parameter" || echo "I was called by $# parameters" 

echo "there is basename of the program called as \"$0\""

echo "first argument associated with \"$0\" is $1"

echo "second argument associated with \"$0\" is $2"

echo "All argument associated with \"$0\" is $@"

echo "All argument associated with \"$0\" is $*"

echo "My name is `basename $0`"
