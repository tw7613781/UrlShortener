# Study note of this project   

## makefile  
- Makefile is a configure file of *make utility*, which has been pre-installed by almost every Linux OS or MacOS
- Simple syntax of makefile is:   
    ~~~~
    target: prerequisites      
    <TAB>recipe
    ~~~~
    For example
    ~~~~
    hello: main.o
        g++ -g -o hello main.c
    main.o: main.cpp
        g++ -c -g main.cpp
- Special keywords     
all: the default targets that need to make   
.PHONY: the targets that no need to be created physically

