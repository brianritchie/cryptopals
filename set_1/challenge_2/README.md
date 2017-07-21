# Challenge 2 WriteUp

This was fun. Both hexadecimal string were assigned to respective variables.

They were then converted to their byte strings for processing.

The XOR Operations utilizes GoLang's slice feature to store all the resulting bytes, aptly names _store_ .

The solution XOR's  _byte\_string\_1_ by declaring a for loop with the length of the string and iterating through each individual character with _byte\_string\_2_ and stores the resulting output in the aforedeclared slice _store_ .

The output is then printed to screen in hexadecimal form using the _%x_ notation with _fmt.Printf_. 
