import ctypes

# Bin
print(0b1010)
# Oct
print(oct(10))
print(int(oct(10)))

print(ctypes.c_byte(0b11110110).value)

# Bitwise https://realpython.com/python-bitwise-operators/#bitwise-and
print("Bitwise")
print(0b1010 & 0b1000) 
print(0b1010 | 0b0100)
print(0b1010 ^ 0b1000)

# Cifragen
print("Cifragem")
x = 2017
x = x ^ 0x51
print(x)
x = x ^ 0x51
print(x)

# shifting bits
print("shift Left")
x = 7
print(x)
x = x << 1
print(x)
x = x << 1
print(x)
x = x << 1
print(x)

# Cadeia de textos 
print(chr(97))

# no terminal echo 'Acesse mentebinaria.com.br'  | hexdump


