with open('day02.txt','r') as f:
	listl=[]
	for line in f:
		strip_lines=line.strip()
		listli=strip_lines.split()
		m=listl.append(listli)


listl = [[int(j) for j in i] for i in listl]

def findsumofdiff(s):
    sum = 0 
    for i in listl:
        a = 0
        b = 0
        a = max(i)
        b = min(i)
        diff =  abs(int(a)-int(b))
        sum += diff
    return sum

print(findsumofdiff(listl))