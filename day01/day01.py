
def captiasum(st):
    sum  = 0
    for index in range(len(st)):
        if index  == len(st) -1:
            if st[index] == st[0]:
                sum += int(st[index])
        elif st[index] == st[index+1]:
            sum += int(st[index])
    return sum

with open("day01.txt") as f:
    lines = f.read()
    print(captiasum(lines))