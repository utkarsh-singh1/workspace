#!/bin/sh

echo Hello World   # --> Hello World
echo "Hello		World" # --> Hello		World

echo "Hello 	\"World\"" # --> Hello 		"World"



# *, ' are not taken seriously by interpreter if placed in ""
echo *
echo *txt
echo "*"
echo "*.txt"


# $,`,",\
# 'To get ouput like this -->   A qoute is ", backslach is \, backtick is `. A few spaces are     ,and dollar is $. $x is 5.

# We have to write it like ->

echo "A qoute is \", backslach is \\, backtick is \`. A few spaces are     ,and dollar is \$. \$x is 5."


