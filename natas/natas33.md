
# General info

Username: natas33
Password: 2v9nDlbSF7jvawaCncr5Z9kSzkmBeoCJ
URL:      http://natas33.natas.labs.overthewire.org

# Special code section

Ok, problema aici e ca trebuie sa facem reverse la un MD5 hash.
Stim ca textul incepe cu `<?php` si se termina cu `?>`.
Am incercat toate combinatiile de valori binare de dimensiune 3, nu a functionat.
Am mai incercat niste metode, tot n-au mers.

O alta idee: serverul totusi face upload la fisier.
Daca cumva _nu trebuie_ sa fac reverse la acel MD5, ci de fapt trebuie sa accesez acel fisier incarcat de mine?
Am incercat asta, totusi pentru niciun fisier incarcat, serverul nu il gaseste - oare il sterge asa repede?

Am incercat un cod cu doua request-uri, sa ma asigur ca nu e problema de viteza - dar tot nu a functionat.

Am incercat si un path traversal exploit, dar nu merge - nu am permisiuni sa scriu in `../../`.

---

Ok so, am mai incercat cateva idei si am inteles urmatoarele:
- exista folderul `/natas33/upload/` unde pot incarca fisiere (path-ul este absolut, adica pleaca de la `/`, nu de la current webserver directory)
- pot scrie in `/tmp/`
- poate trebuie sa gasesc un loc unde sa scriu un fisier care este executat automat - dar unde?

Nu poti face HTML injection - nu ai voie cu `/` in numele fisierului.
Nu poti face PHP injection - un simplu `echo` nu este reinterpretat.

---

Am decis sa ma uit pe un writeup, deoarece nu mai aveam idei.
Cel pe care [l-am gasit](https://learnhacking.io/overthewire-natas-level-33-walkthrough/) vorbeste despre un _atac Phar_, no idea what that is.
Dar aparent foloseste **Burp Suite**.