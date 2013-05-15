goperf
======

goperf is a little tool I've written to start experimenting with go...
oh and to test our REST api server response time, yeah that...

ATTENTION PLEASE!!!
this is my very first program written in go, so the code may not follow the
golang guidelines.

Usage:
---
goperf --uri http://domain.to.test/somepage.html --request 10 #the number of requests

it will (should) return the average loading time of the specified uri
