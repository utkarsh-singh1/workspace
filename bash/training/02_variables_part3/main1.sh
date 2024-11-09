
echo -en " What is your name [ `whoami` ] "
read myname
echo "My name is -  ${myname:-`whoami`}"

# in bash ":-" is a special character which enables variblales to have default values in case there is no input provided by end-user. 
