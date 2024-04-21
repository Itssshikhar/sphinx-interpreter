# sphinx-lang

![](https://media.discordapp.net/attachments/1215265039892750427/1222187436281561150/shikhar_7985_sphinx_with_camera_angle_towards_top_from_lower_ri_3bb46794-d267-4374-80ac-0aec18350ae1.png?ex=6630fd3f&is=661e883f&hm=d01ca81f10e5b0d17a74158fd5c163c3b37455cff30fc29e5cde13b0c390bcd8&=&format=webp&quality=lossless&width=400&height=350)

sphinx lang is an interpreted and fully functional programming language. It supports mathematical expressions, variable
bindings, functions and the application of those functions, conditionals, return statements and
even advanced concepts like higher-order functions and closures. And it also supports different
data types like: integers, booleans, strings, arrays and hashes.

## Running locally
Having go **(version >= 1.21.5**) installed in the system is a must, after that, open the terminal and run the following command:

`go run main.go`

## Example usages
### puts function
puts prints the given arguments on new lines to STDOUT. It calls the Inspect() method on the objects passed in as arguments and prints the return value of these calls.
```
>> puts("Hello!")
Hello!
>> puts(1234)
1234
>> puts(fn(x) { x * x })
fn(x) {
(x * x)
}
```
### hashes
A hash is what’s sometimes called hash, map, hash map or dictionary in other programming languages. It maps keys to values.
```
>> let myHash = {"name": "Jimmy", "age": 72, "band": "Led Zeppelin"};
>> myHash["name"]
Jimmy
>> myHash["age"]
72
>> myHash["band"]
Led Zeppelin
```
### arrays
an array is an ordered list of elements of possibly different types. Each element in the array can be accessed individually. Accessing individual elements by their index in the array is done with a new operator, called the index operator: `array[index].`
```
>> let myArray = ["deez", "nuts", 69, fn(x) { x * x }];
>> myArray[0]
deez
>> myArray[2]
69
>> myArray[3](2);
4
```
### strings
strings are a sequence of characters. They are first-class values, can be bound to identifiers, used as arguments in functions calls and be returned by functions. They look just like the strings in many other programming languages: characters enclosed by double quotes.
```
>> let firstName = "shi";
>> let lastName = "pakdey hail";
>> let fullName = fn(first, last) { first + " " + last };
>> fullName(firstName, lastName);
shi pakdey hai
```
### functions & function calls
Functions are treated like any other value: we can bind them to names, use them in expressions, pass them to other functions, return them from functions and so on.
```
>> let add = fn(a, b, c, d) { return a + b + c + d };
>> add(1, 2, 3, 4);
10
>> let addThree = fn(x) { return x + 3 };
>> addThree(3);
6
>> let max = fn(x, y) { if (x > y) { x } else { y } };
>> max(5, 10)
10
```
Passing around functions, higher-order functions and closures will also work:
```
>> let callTwoTimes = fn(x, func) { func(func(x)) };
>> callTwoTimes(3, addThree);
9
>> callTwoTimes(3, fn(x) { x + 1 });
5
>> let newAdder = fn(x) { fn(n) { x + n } };
>> let addTwo = newAdder(2);
>> addTwo(2);
4
```
### these are many more things besides these that I'm too lazy to mention here

## References
### The OG pratt parser paper https://tdop.github.io/ (biggest help)
### Colin James explains TDOP https://youtu.be/2l1Si4gSb9A?si=Bh2PkYsXCpoy7tnY (biggest help)
### Thorsten ball https://interpreterbook.com/ (biggest help)
### Pixeled on yt https://youtube.com/playlist?list=PLUDlas_Zy_qC7c5tCgTMYq2idyyT241qs&si=eRyQgHKNaE8zp1K3
### Immo Landwerth https://youtube.com/playlist?list=PLRAdsfhKI4OWNOSfS7EUu5GRAVmze1t2y&si=J00dN1ywockH7IoZ
