run:
	@gcc -std=c23 -Wall -Wextra -Wpedantic -Werror -lm -o hello hello.c
	@./hello

broken1:
	@gcc -std=c23 -Wall -Wextra -Wpedantic -Werror -lm -o hello_broken_1 hello_broken_1.c

broken2:
	@gcc -std=c23 -Wall -Wextra -Wpedantic -Werror -lm -o hello_broken_2 hello_broken_2.c

broken3:
	@gcc -std=c23 -Wall -Wextra -Wpedantic -Werror -lm -o hello_broken_3 hello_broken_3.c
	@./hello_broken_3
	@rm ./hello_broken_3

pipeline:
	@gcc -std=c23 -E hello.c -o hello.i
	@gcc -std=c23 -S hello.c -o hello.s
	@gcc -std=c23 -c hello.c -o hello.o
