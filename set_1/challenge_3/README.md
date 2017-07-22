# Challenge 3 Write-up

Given Matasano's challenge brief, we learn a few crucial things:

- The Plaintext was XOR'ed with a single character key (translation: we can be lazy, CRUCIAL POINT)
- We could probably XOR it by hand (translation: we don't need to be complex)

The challenge also hinted at creating a scoring table and looking at character frequency. 

---

I decided to ignore the hint because we already know the length of the key. had the key length not been known, it would definitely be worth identifying the key-length and subsequently doing a lookup table scoring of characters to see the freqeuncy of occurences - hence the joke about _ETAOIN SHRDLU_ which are the most commonly occuring characters in the English language after _space_.

Given all the information above, we can easily bruteforce the key by knowing the following XOR operation which is 

Ciphertext = Plaintext ^ Key

The key is a single character and thus will occur anywhere between 1 - 255 because XOR with 0 produces the same text back.

So the solution proposes the following, 

Each byte of the integer (representing the full range of characters) will be rotated through as the key against each byte in the encoded hex string.

The result will be analyzed for a phrase that is readable as the full plaintext

which in this case is "Cooking MC's like a pound of bacon" which occurs for the int 88 #QED

---

To-Do

- Write a better solution to make the bruteforce look neater OR implement the scoring which accurately just scores in the background and spits out a single text line