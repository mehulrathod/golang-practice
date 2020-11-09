A Go program consists of various tokens. A token is either a keyword, an identifier, a constant, a string literal, or a
symbol. For example, the following Go statement consists of six tokens −

```fmt.Println("Hello, World!")```

The individual tokens are −

- fmt
- .
- Println
- (
-   "Hello, World!"
- )
****
###Identifier:
```Identifiers are names for variables, methods, classes, or parameters. Identifiers can have alphanumerical characters, underscores and dollar signs ($). It is an error to begin a variable name with a number. White space in names is not permitted.```

###Literals:
```A literal is a textual representation of a particular value of a type. Literal types include boolean, integer, floating point, string, null, or character. Technically, a literal will be assigned a value at compile time, while a variable will be assigned at runtime.```

###Keywords:
```A keyword is a reserved word in Java language. Keywords are used to perform a specific task in the computer program. For example, to define variables, do repetitive tasks or perform logical operations.```

****
Defines a function that is named "SalesTax" and has one parameter named "price". The type of price is "double". The function's return type is also a double.

`double SalesTax(double price)
 {
   return 0.05 * price;
 }`