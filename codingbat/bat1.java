the aIsBigger() method should return true if the int parameter a 
is larger than b by 2 or more. This code uses an if with &&
 ("and") to return true if the condition is met. Alternately, 
 the run falls through to the return-false at the bottom.  

 public boolean aIsBigger(int a, int b) {
 	if (a >= (b+2)) 
 	 return true;

 	return false;
 }


Alternately it can be done with an if/else structure like this:

 public boolean aIsBigger(int a, int b) {
 	if (a >= (b+2)) 
 	 return true;

 	return false;
 }

 And in fact, since the boolean test is true when 
 we want to return true, and false when we want to 
 return false, it can be written as a one-liner like this:

 public boolean aIsBigger(int a, int b) {
 	return (a >= (b+2));
 }

========================================================================================================
========================================================================================================
========================================================================================================
========================================================================================================
========================================================================================================
 


 With a string, str.substring(i, j) returns the String that starts at index i
  and goes up to but not including j.

   The first char in the String is at  index 0, so str.substring(0, 2)
    returns a string of the first two chars. 

   The method str.length() returns the length of a string. Compare two strings
    like this: str1.equals(str2).

     Do not compare two strings with == which
     looks reasonable but does not work correctly in all cases.

. == should be used during reference comparison. == checks if
 both references points to same location or not. equals()
  method should be used for content comparison


This twoE() example method returns true if the string 
contains exactly two 'e' chars. The code:
"for (int i=0; i<str.length(); i++) { ..." is very 
standard code to look at each position in a String.

 public boolean twoE(String str) {
 	for (int i= 0; i< str.length() ; i++) {
 		
 	}
 	  	if ( 	 str.substring(i, j)  == ) {
 		return true; 
 	}
 }

 for int i = 0 -> length -1

 	if arr[i] == arr[i+1] {
 		return true;
 	}

This new6() method makes and returns a new int array of size N that is filled with the value 6.

new6(int N)  {
	        int[] array = {11,12,13,14,15};

 	return 	 arr
}


Here is a simple recursive method that counts the number of "A" in the given string.

// [A B C]
// [A A]
// [A]
// []


int counter(String str, int counter) {

	if str == "" || str.length ==0 {
		return 0;
	}

	return counter += counter(str.substring(1,str.length))
	
}



// 61c,162,164

Here is a simple recursive method that counts the number of "A" in the given string.

public int countA(String str) {
  // Base case -- return simple answer
  if (str.length() == 0) {
    return 0;
  }

  // Deal with the very front of the string (index 0) -- count "A" there.
  int count = 0;
  if (str.substring(0, 1).equals("A")) {
    count = 1;
  }

  // Make a recursive call to deal with the rest of string (the part beyond the front).
  // Add count to whatever the recursive call returns to make the final answer.


  // Note that str.substring(1) starts with char 1 and goes to the
  // end of the string.
  return count + countA(str.substring(1));
}

