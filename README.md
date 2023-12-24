# Test task

## Description

Design and implement “Word of Wisdom” tcp server.

- TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained. 
- After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes. 
- Docker file should be provided both for the server and for the client that solves the POW challenge

## Solution

As the POW algorithm, **scrypt** was chosen in combination with **hash** generation that includes **few leading zeros**. The scrypt algorithm allows for the adjustment of parameters in such a way as to ensure both computational resources and memory resources on the client side. Moreover, it is resistant to automated hash generation through specialized mining equipment (ASIC), which complicates potential DDOS attacks.

## How to use
A Makefile can be used to start a server and to run a client.