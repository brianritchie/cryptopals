# Challenge 1 WriteUp

---

Detailed Writeup can be found here: https://brianritchie.me/crytopals-solution-set-1-challenge-1

---

I have taken a shortcut to the challenge by using the pre-existing libraries in GoLang to produce the output which are the following:

- "encoding/hex"
- "endoding/base64"

The program takes the provided hexadecimal string of the challenge and checks if it fulfils the requirements for a hexadecimal string and if it does, assigns it to _hex\_string_ .
It then converts _hex\_string_ to its ASCII bytes form as preferred by the challenge and assings the result of the conversion to _byte\_string_ .

_byte\_string_ is then passed on to the base64 encoding library and the result is assigned to _base64\_encode_ .

The base of the challenge is considered solved.

## Optimal To-Dos

- Create the conversion manually without relying on the standard libraries
- Write tests to confirm the validity of the aforementioned manual conversions
- Publish an optimized and faster solution

