function test()
	return "hello"
end


print(type("print"))
print(type(5))
print(type(3.4))
print(type(print))
print(type(1,2))
print(type(true))
print(type(x))
print(type(type(x)))
print(type(type(y)))
print(type(test()))


-- lua is dynamically typed programming language
--
print("type of a is ->",type(a))
a = 10
print("now type of a is ->",type(a))
a = "Hello world"
print("now type of a is ->",type(a))


-- in lua function is a first-class function(https://developer.mozilla.org/en-US/docs/Glossary/First-class_Function)
--
a = print
a(type(a))

