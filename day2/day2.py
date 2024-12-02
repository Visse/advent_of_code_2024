from pathlib import Path

input = Path('day2.in').read_text()

input = [
    [int(n) for n in l.split()]
    for l in input.splitlines()
]

safe = 0
for r in input:
    if any((r[0]<r[1]) != (p<v) for p,v in zip(r, r[1:])):
        # print('Unsafe:', str(r).replace(',', '')) #, r[0]<r[1], [(p-v, r[0]<r[1], p < v, (r[0]<r[1]) != (p <v)) for p,v in zip(r, r[1:])], )
        continue
    if any(abs(p-v) not in (1, 2, 3) for p,v in zip(r, r[1:])):
        #print('Unsafe:', str(r).replace(',', ''))
        continue
    #print('Safe:', str(r).replace(',', ''))
    safe += 1
print('Part 1:', safe)
