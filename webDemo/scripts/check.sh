#!/bin/bash

proPath=./volume/$PROGRAM
outname=./volume/$PROGRAM-out
errname=./volume/$PROGRAM-err
go run $proPath>$outname 2>$errname