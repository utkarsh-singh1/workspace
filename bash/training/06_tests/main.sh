#!/bin/sh

# test is often denoted by [, to increase readability, it is generally used in case of comparision, used by if and while.

# General syntax -> Ex -> if [ $i = 5 ] 
# Remember space around "[" and "=", these are operators and they need space around them to make them work.
# Sometimes some shell also accepts "==" as comparision operator but "=" is general in use for strings and '-eq' for intergers.

# If syntax ->
# if [ ... ] 
# then
# 	statement
# fi
#
# or 
# 
# if [ ... ] ; then
# 	statement
# fi

# If - else syntax ->
# if [ ... ]
# then
# 	statement
# else
# 	statement
# fi


# If-elseif-Else syntax ->
# if [ ... ] ; then
# 	statement
# elif [ ... ] ; then
# 	statment'
# else 
# 	statement
# fi
#
