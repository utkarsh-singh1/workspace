print("Hello World")

-- sample complex program

function fact (a)
    if a == 1 then
        return 1
    else
        return a * fact(a-1)
    end
end

print("Taking an input:")
a = io.read("*num")
print(fact(a))

