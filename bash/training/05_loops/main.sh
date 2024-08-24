#!/bin/sh


# for loops
for i in 1 2 3 4 5
do
		touch $i.txt
done

for i in hello 1 *  2 world
do
		echo "value of i is... $i"
done

# While loop
message=hello
while [ $message != bye ] 
do	
		echo "Current message is $message"
		echo "Please type your message ( bye to quit )"
		read message
		echo "You have typed $message"
done


# : --> always yiels true
while :
do	
		echo "Please type something here (Ctrl+C to quit)"
		read input
		echo "You have some input... $input here"
done


