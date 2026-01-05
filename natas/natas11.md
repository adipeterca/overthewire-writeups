
# General info

Username: natas11
Password: UJdqkK1pTu6VLt9UHWAgRZz6sVUZ3lEk
URL:      http://natas11.natas.labs.overthewire.org

# Special code section


```py
import requests
import base64


# base64 decode -> xor -> json
# deci eu trebuie sa il fac invers: json, xor, base64 encode

# trebuie sa aiba showpassword si bgcolor

def xor(data, key = "eDWo"):
    new_data = []
    for i, d in enumerate(data):
        new_data.append(chr(ord(data[i]) ^ ord(key[i % len(key)])))

    data = ''.join(new_data)


    return data


data = '{"showpassword":"yes","bgcolor":"#ffffff"}'
data = xor(data)

data = bytes(data, encoding="utf-8")
data = base64.encodebytes(data)

data = bytes.decode(data[:-1])
print(data)

response = requests.get(
    url="http://natas11.natas.labs.overthewire.org/",
    auth=("natas11", "UJdqkK1pTu6VLt9UHWAgRZz6sVUZ3lEk"),
    cookies={
        "data": data
    }
)

print(response.text)
print(response.headers)


# resp = base64.decodebytes(b"HmYkBwozJw4WNyAAFyB1VUcqOE1JZjUIBis7ABdmbU1GIjEJAyIxTRg=")
# resp = bytes.decode(resp)


# other = xor(resp, '{"showpassword":"no","bgcolor":"#ffffff"}')
# print(other)    # prints 'eDWoeDWoeDWoeDWoeDWoeDWoeDWoeDWoeDWoeDWoe'
# print(len(other))

# print(f"response before XOR: >> {resp} <<")
# resp = xor(resp)
# print(f">>>>> {resp} <<<<<")
```


A trebuit sa fac reverse la un XOR si sa sparg cheia.
Am folosit principiul de utilizare al XOR-ului, ca A ^ B = C, dar si A ^ C = B.

Initial, am crezut ca trebuie sa folosesc parola de la natas11.
Dar am vazut dupa ca mi se tot returna mereu acelasi cookie, indiferent de ce trimiteam eu.
Astfel, mi-am dat seama ca se pleaca de la un JSON de start (fara spatii) si se construieste acel raspuns
Dupa, am vazut ca nu era base64 valid, dar punandu-l in gepeto, am vazut ca e URL-Base64, pentru ca `%3D` e de fapt `=`. `3D` e valoarea in hex care apare si pe tabela ASCII.

In final, am decodat outputul de la base64, l-am XOR-at cu ce un JSON de start, adica `{"showpassword":"no","bgcolor":"#ffffff"}` si dupa am vazut ca cheia se repeta.