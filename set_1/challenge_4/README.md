# Challenge 4 Write-up

Given Matasano's challenge brief, we can summize a few things:

- Given the 60 lines of hex encoded strings, only one has been encrypted by a single character XOR (translation: we know the key length and there is only one winner)
- Code from previous challenge should help (so we can reuse a working solution)

---

We could technically brute force it just like the previous challenge but there will be too much noise on screen, so we need a way to filter the output.

We achieve this by using the previous code to decode and XOR each string and subsequently pass it through a filter to identify which byte string has the largest amount of matches to the most commonly used characters in the English language - ETAOIN SHRDLU

There are multiple ways to do this, but the solution here achieves what we need using RegEx. We pass each decoded string through a RegEx filter to check for matches to the aforementioned characters and keep increasing the amount of matches until we arrive at the solution on screen:

The string reviewed is --> 7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f
b is 21 and the plaintext is -->  nOWTHATTHEPARTYISJUMPING*
The amount of matches to the English Language Common characters is --> 16

QED