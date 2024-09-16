// Created by Jack Foley (C00274246)
// 13/09/2024

#include <stdio.h>
#include <stdlib.h>
#include <string.h>


/* 

How do run?

Compile the program using the following command:
gcc main.c -o main

Then run the output in terminal:
./main

 */


/* 

Function used to get the integer value of a single roman numeral.

Arguments:
- c: The roman numeral character to convert to an integer value.

Returns:
- int: The value of the character.

 */
int romanCharToInt(char c) {
    switch (c) {
        case 'I': return 1;
        case 'V': return 5;
        case 'X': return 10;
        case 'L': return 50;
        case 'C': return 100;
        case 'D': return 500;
        case 'M': return 1000;
        default: return 0;
    }
}


/* 

Function used to convert a roman numeral string to the integer value.

Arguments:
- s: The roman numeral string to co
nvert to the integer equivalent.

Returns:
- int: The value of the roman numeral string.


 */
int romanToInt(const char* s) {
    int result = 0;
    int length = strlen(s);
    
    for (int i = 0; i < length; i++) {
        int current = romanCharToInt(s[i]);
        int next = (i + 1 < length) ? romanCharToInt(s[i + 1]) : 0; // if next is out of bounds, set to 0
        
        if (current < next) {
            result -= current;
        } else {
            result += current;
        }
    }
    
    return result;
}

int main() {

    // Get user input, store in char array
    printf("Enter your Roman Numeral string: ");
    char* roman = malloc(15);
    scanf("%s", roman);

    // Validate that the str is not empty
    if (strlen(roman) == 0) {
        printf("No roman numeral provided\n");
        return 1;
    }

    // Validate that the size is less than 15 characters
    if (strlen(roman) > 15) {
        printf("Roman numeral is too long\n");
        return 1;
    }

    // Validate that the string has all valid roman numeral characters
    for (int i = 0; i < strlen(roman); i++) {
        if (romanCharToInt(roman[i]) == 0) {
            printf("Invalid roman numeral character: %c\n", roman[i]);
            return 1;
        }
    }

    printf("The integer value of %s is %d\n", roman, romanToInt(roman));

    // Free the allocated memory and exit
    free(roman);
    return 0;
}