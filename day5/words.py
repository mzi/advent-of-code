#!/usr/bin/env python

import re

text = open('input.txt').readlines()


def part1():
    n = 0
    for line in text:
        if re.compile(r'.*(ab|cd|pq|xy).*').match(line) is not None:
            continue
        if re.compile(r'(.*[aieou].*){3,}').match(line) is None:
            continue
        if re.compile(r'.*(\w)\1.*').match(line) is not None:
            n += 1
    return n


def part2():
    n = 0
    for line in text:
        if re.compile(r'.*(\w).\1.*').match(line) is None:
            continue
        if re.compile(r'.*(\w\w).*\1').match(line) is not None:
            n += 1
    return n

print part1()
print part2()