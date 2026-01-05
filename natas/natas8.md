
# General info

Username: natas8
Password: xcoXLmzMkoIP9D7hlgPlh9XD7OgLAe5Q
URL:      http://natas8.natas.labs.overthewire.org

# Special code section

Am gasit in source code linia asta `bin2hex(strrev(base64_encode($secret)))`
Dupa, am cautat sa vad ce inseamna `bin2hex` in PHP
dupa, am facut asa in python:

```py
import base64

output = "3d3d516343746d4d6d6c315669563362"

# gotta find the input
# bin2hex(strrev(base64_encode($secret)));

res = bytes.fromhex(output)
res = res[::-1]

res = base64.decodebytes(res)

res = res.decode(encoding="utf-8")
print(res)
```
