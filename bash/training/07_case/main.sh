#!/bin/sh

# Case syntax
# case expression in

#     case1)
# 	statement
# 	;;
#     case2)
# 	statment
# 	;;
#     *)
# 	statement
# 	;;
# esac
#

echo "Please tell me something...."


while : ; do

    read INPUT_STRING
    case $INPUT_STRING in

	hello)
	    echo " Hi everynyan "
	    ;;
	bye)
	    echo " Bye! see you again "
	    break
	    # use exit to break from the whole code
	    ;;
	*)
	    echo " Soory! I don't understand what you are saying "
	    ;;
    esac

done

echo
echo "Thanks for talking to me"
