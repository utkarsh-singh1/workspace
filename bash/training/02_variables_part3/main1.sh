
echo -en " What is your name [ `whoami` ] "
read myname
echo "My name is -  ${myname:-`whoami`}"
