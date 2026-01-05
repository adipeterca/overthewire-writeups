
# General info

Username: natas5
Password: 0n35PkggAPm2zbEpOU802c0x0Msn1ToK
URL:      http://natas5.natas.labs.overthewire.org

# Special code section

```py
response = requests.get(
    url="http://natas5.natas.labs.overthewire.org/", 
    auth=("natas5", "0n35PkggAPm2zbEpOU802c0x0Msn1ToK"),
    cookies={"loggedin": "1"}
)
```

Headerul de `Authorization` poate fi:
* `Basic <base64-string>`, unde `base64-string` e `username:password`
* `Bearer <token>`, unde `token` e un base64url-encoded `header.payload.signature`

Headerul de `Cookie` contine un field de key-value separate prin `;`
    * exemplu: `session_id=abc123; user_id=456; remember=true`
    * poti seta `cookies` direct din python si e _mai indicat_ decat sa setezi din `headers`

E important de mentionat ca, la `cookies`, nu poti avea valori de tipul `int`, doar `str` !!