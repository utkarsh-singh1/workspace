# while read to read from a file 
while read expression
do
		case $expression in
				hello) echo "English" ;;
				howdy) echo "America" ;;
				gday)  echo "Australia" ;;
				bonjour) echo "France" ;;
				"guten tag") echo "German" ;;
				*) echo "Can't detect from your Input($expression) mate" ;;
		esac
done < mytext.txt
