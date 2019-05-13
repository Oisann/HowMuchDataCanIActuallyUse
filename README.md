# How much data can I actually use?

How much data can I actually use per second, without having my Raspberry Pi Cluster exceeding my data cap on my mobile plan. I have a 1000GB/month plan, limited to 10Mbps download and upload on my phone, with a twin sim connected to my rpi cluster. If my fiber for some odd reason would go down, I also rely on this connection for the rest of my devices. I should probably rate limit myself at another percentage used then.

## What is this doing?

* Data cap: 1000GB
* Renew date: 1st
* Average Download speed: 10Mbps
* Average Upload speed: 2Mbps

Given an amount used, this calculates how much data can be downloaded per second without exceeding the data cap.

## Why?

My Raspberry Pi cluster is partly running off of my mobile data. Usually this is used by the tor relay and some uptime services.

## How?

This application is part of a cron job that checks the mobile plan usage and the average download speed of the server. I rate limit the server to the *min* value of this scripts *result* and the *actual download speed* (10mbps), but only if I have used more than 50% of my data.

## Examples

#### Under 50%
* Used 7.3GB out of 1000GB
* Date: May 13th ~22:35
* Renew date: 1st

*Result: 10000 Kbps*

#### Over 50%

* Used 501GB out of 1000GB
* Date: May 13th ~22:35
* Renew date: 1st

*Result: 2258 Kbps*