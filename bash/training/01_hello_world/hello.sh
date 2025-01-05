#!/bin/sh


# This above line will insure that during execution is done by the sh from /bin/sh (general location of sh in most of the GNU/linux systems which is symbolic linked to bash).


# this is a comment
echo "Hello World"       #Single argument

echo Hello 	World		#Two arguments

# Simple string as * is permissable
echo "Hello * World"

# Here there are 3 arguments, * list all files around this files
echo Hello * World

# 2 different arguments
echo "Hello" World

# 1 single argument
echo "Hello"World

# 3 arguments
echo Hello "	" World

# 3 arguemnts
echo "Hello " * " World"

# `Hello` is treat as command, anything b/t `` will be execute by shell
echo `Hello` World

# 2 arguments
echo 'Hello' World
