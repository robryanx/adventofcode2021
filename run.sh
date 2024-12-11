#!/bin/bash

go build -o ~/adventofcode/2021/bin/$1 ~/adventofcode/2021/days/$1/*

if [ $? -eq 0 ]
then
    ~/adventofcode/2021/bin/$1;
fi