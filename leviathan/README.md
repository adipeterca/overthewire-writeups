Because this series is not very programming oriented, I'll simply put my walkthroughs here.

# Level 0

Just inspect the `.backup` file.

# Level 1

I used [Ghidra online](https://dogbolt.org/) to disassemble the file.  
I tried with `xxd`, `gdb` and `strings` but couldn't figure out the whole password, just bits of it.

# Level 2

Needed to create a temp directory and a file called `test;bash`.  
This is because `access()` reads the file permissions using the RUID, while `system(cat file_path)` will use EUID.  
This mismatch between them is the problem: you'll need a file that can be read by both users.

# Level 3

Inspected the binary in the same online decompiler.  
Saw some strings that had nothing to do with the actual password verification - the function `do_stuff()` has the actual comparison.  
In there you'll find the password string.

# Level 4

This one was dead easy - just had to look in the home directory and convert from binary to ASCII.

# Level 5

Just needed to create a **symbolic** link - apparently, you can just create a symlink to a file you do not own.

# Level 6

Simple `for` loop in bash, nothing more. Code is around `7000` - `8000`.