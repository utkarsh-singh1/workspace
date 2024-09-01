#!/bin/sh

X=0
while [ -n "$X" ] ; do
		echo "Enter your Input (Press return to exit)"
		read X
		if [ -n "$X" ] ; then
				echo `"You have entered - $X"
		fi
done
