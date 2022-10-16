k=int(input())

ans_h = '21'
colon = ':'
ans_m = ''

if k < 60:
    ans_m=str(k)
else:
    ans_h='22'
    ans_m=str(k-60)

ans_m='0'+ans_m
ans_m=ans_m[-2:]

ans=ans_h+colon+ans_m
print(ans)
