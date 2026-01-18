
# General info

Username: natas31

Password: m7bfjAHpJmSYgQWWeqRE2qVBuMiRNq0y

URL:      http://natas31.natas.labs.overthewire.org

# Special code section

Am inceput prin a citi _cum poti trimite fisiere prin `requests.py`_, am gasit [acest link](https://www.geeksforgeeks.org/python/how-to-upload-files-using-python-requests-library/).

Am incercat sa injectez niste PHP dar este safely escaped prin functia `escapeHTML`, dar nu a mers.

Urmatoarea idee a fost sa incarc un fisier `test.php` pe care dupa sa il accesez la `http://natas31.natas.labs.overthewire.org/tmp/test.php`, dar ma gandeam ca local va fi salvat cu un nume aleator.
Din pacate, nu a mers - fisierul e sters dupa ce scriptul este executat.

Cred ca e o problema de charsets, dar nu inteleg cum s-o exploatez - am citit [acest articol](https://zaynar.co.uk/posts/charset-encoding-xss/) dar nu vad cum s-ar aplica aici.

---

Am citit un [write-up](https://learnhacking.io/overthewire-natas-level-31-walkthrough/), iar solutia aparenta se bazeaza pe o vulnerabilitate foarte veche, care tine de cum anume interpreteaza Perl ARGV - nu am inteles foarte bine explicatia.

Am gasit [acest write-up](https://github.com/javiunzu/natas/blob/master/natas31.py) scris in Python, am incercat sa inteleg de ce
```py
url="http://natas31.natas.labs.overthewire.org/index.pl?ls | xargs echo |",
```
nu merge, dar
```py
url="http://natas31.natas.labs.overthewire.org/index.pl?cat /etc/natas_webpass/natas32 | xargs echo |",
```
merge.
LE: _aparent conteaza acel spatiu inainte de_ `|`

Nu am inteles concret ce este cu vulnerabilitatea asta.