# ping-status

![Image](docs/img/ping-status-architecture.drawio.svg)

Functions:
- Using pro-bing library (ICMP echo), it sends ICMP Echo Request packet(s) and waits for an Echo Reply in response
- Printed variables atm: website adress, ping (avg) in ms, packet loss in % and TTLs
- Later it gets connected to a monitoring system with Grafana/Prometheus, through a /metrics API