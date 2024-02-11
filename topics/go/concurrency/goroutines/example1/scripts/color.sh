while IFS= read -r line; do
    for i in $(seq 0 $((${#line} - 1))); do
	char="${line:i:1}"
	if [[ "$char" =~ [A-Z] ]]; then
	    printf "\033[32m%c\033[0m" "$char"
	elif [[ "$char" =~ [a-z] ]]; then
	    printf "\033[31m%c\033[0m" "$char"
	else
	    printf "$char"
	fi
    done
    echo ""
done
