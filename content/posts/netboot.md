+++ 
draft = false
date = 2026-10-31T23:59:45+01:00
title = "Zero-Touch Network Boot Infrastructure (PXE / iPXE / ProxyDHCP)"
description = "A lightweight, robust, and non-disruptive network boot solution that turns any PC on your local network bootable on-the-fly. Powered by netboot.xyz, dnsmasq (ProxyDHCP), and iVentoy, this repository provides automated startup scripts and configurations to serve live Linux distributions, installers, and rescue utilities (Clonezilla, Partition Magic, etc.) without altering your existing core network router or primary DHCP server."
slug = ""
authors = []
tags = ["linux", "docker", "git", "gitops", "netboot", "iventoy", "tftp", "pxe", "dhcp", "tcpdump", "dns", "nginx"]
categories = []
externalLink = ""
series = []
+++

[![Repository](https://img.shields.io/badge/GitHub-mozkaya1%2Fnetboot-blue?logo=github)](https://github.com/mozkaya1/netboot)
[![Stack](https://img.shields.io/badge/Stack-Linux%20%7C%20ProxyDHCP%20%7C%20TFTP%20%7C%20iPXE-orange)](#-tech-stack--skills-demonstrated)

A lightweight, robust, and non-disruptive network boot solution that turns any PC on your local network bootable on-the-fly. Powered by **netboot.xyz**, **dnsmasq (ProxyDHCP)**, and **iVentoy**, this repository provides automated startup scripts and configurations to serve live Linux distributions, installers, and rescue utilities (Clonezilla, Partition Magic, etc.) without altering your existing core network router or primary DHCP server.

---
## 👤 Author
GitHub: [@mozkaya1](https://github.com/mozkaya1)  
Repository Link: [https://github.com/mozkaya1/netboot](https://github.com/mozkaya1/netboot)

## 🏗️ Architecture & Core Concept

Modifying a production or home router's DHCP scope to support PXE booting is often risky, restricted, or impossible on ISP gear. This project solves that by deploying a **ProxyDHCP** configuration:

```
                  +-----------------------------------+
                  |      Primary Network Router       |
                  |     (Existing DHCP Server)        |
                  |  Provides IPs: 192.168.1.0/24     |
                  +-----------------+-----------------+
                                    |
                                    v
+-----------------------+   +-------+-------+   +-----------------------+
|  Client PC (PXE Boot) |<--| Local Network |-->| DevOps Netboot Server |
|  Broadcasting PXE     |   |   Subnet      |   | (dnsmasq ProxyDHCP +  |
|  Request on Port 67   |   +---------------+   |  TFTP / HTTP / iPXE)  |
+-----------------------+                       +-----------------------+
```

1. **Primary DHCP Server:** Handled by your local router—continues issuing normal IP leases (`192.168.x.x`).
2. **`dnsmasq` (ProxyDHCP):** Listens passively on `:67` and steps in *only* when a client broadcasts a PXE boot request, supplying the necessary `next-server` IP and bootfile name (`undionly.kpxe` or `ipxe.efi`).
3. **TFTP / HTTP Servers:** Deliver the actual bootloader files, kernels, and ISO payloads dynamically.

---

## 🛠️ Tech Stack & Skills Demonstrated

* **Linux Systems Engineering:** Advanced network configuration, daemon tuning, and service management.
* **Network Protocols:** DHCP, ProxyDHCP, TFTP, HTTP, PXE, iPXE, and TCP/IP stack diagnostics.
* **Automation & Orchestration:** Bash automation scripts (`.sh`), Docker integration.
* **Traffic Analysis & Observability:** Packet interception using `tcpdump` and Wireshark.
* **Tools & Software:** [dnsmasq](https://thekelleys.org.uk/dnsmasq/doc.html), [netboot.xyz](https://netboot.xyz/), [iVentoy](https://www.iventoy.com/).

---

## ⚠️ Port Collision & Stack Selection

Both **netboot.xyz + dnsmasq** and **iVentoy** handle network bootstrapping and contain embedded DHCP proxy and TFTP services.

> 🛑 **IMPORTANT:** Running both stacks at the same time on the same host machine will cause severe port collisions on **`:67` (DHCP/ProxyDHCP)**, **`:68` (DHCP Client)**, and **`:69` (TFTP)**.

Choose **one** of the two deployment stacks based on your use case:

| Feature | Option A: `netboot.xyz` + `dnsmasq` | Option B: `iVentoy` |
| :--- | :--- | :--- |
| **Primary Use Case** | Cloud/Online netbooting, live distros, OS installers | Local ISO repo repository, offline deployments |
| **DHCP Proxy** | Handled by `dnsmasq.conf` | Built-in to iVentoy daemon |
| **Flexibility** | Highly customizable iPXE boot scripts | Simple web GUI management for local ISO storage |

---

## 📂 Repository Structure

```text
.
├── dnsmasq.conf            # ProxyDHCP and TFTP configuration file
├── netbootxyz-start.sh     # Automation script to start netboot.xyz + dnsmasq
├── iventoy-start.sh        # Automation script to start iVentoy service
├── static/                 # Gallery folder containing interface screenshots
│   ├── netboot(*).png         # netboot.xyz operational screenshot
│   └── iventoy(*).png         # iVentoy operational screenshot
└── README.md               # Project documentation
```

---

## ⚙️ Quick Start Guide

### Step 1: Clone the Repository
```bash
git clone https://github.com/mozkaya1/netboot.git
cd netboot
```

### Option A: Launching netboot.xyz Stack
1. Inspect `dnsmasq.conf` and adjust your network interface bindings if necessary (e.g., `eth0`, `wlan0`).
2. Make the startup script executable and run it:
   ```bash
   chmod +x netbootxyz-start.sh
   ./netbootxyz-start.sh
   ```
3. Boot any client host on your LAN, select **Network Boot (PXE)** in BIOS/UEFI, and enjoy dynamic boot options.

### Option B: Launching iVentoy Stack
1. Ensure the `dnsmasq` service is stopped to clear ports `:67` and `:69`.
2. Make the script executable and launch iVentoy:
   ```bash
   chmod +x iventoy-start.sh
   ./iventoy-start.sh
   ```
3. Open your browser and navigate to the iVentoy Web Console (`http://<server-ip>:26000`) to load ISO images.

---

## 📸 Interface Gallery

The project assets and operational screenshots are stored in the `./static/` folder:

### iVentoy Stack
![iVentoy](/iventoy-1.png)
![iVentoy](/iventoy-2.png)
![iVentoy](/iventoy-3.png)
![iVentoy](/iventoy-4.png)
![iVentoy](/iventoy-5.png)

### netboot.xyz Stack
![netboot.xyz Interface](/netboot-1.png)
![netboot.xyz Interface](/netboot-2.png)
![netboot.xyz Interface](/netboot-3.png)
![netboot.xyz Interface](/netboot-4.png)

### DHCP Hadshake Debug Log -- tcpdump
![tcpdump](/dhcp-handshake.png)

---

## 🔍 Diagnostics & Packet Inspection

To verify ProxyDHCP responses and debug network handshake issues on your server interface, run `tcpdump`:

```bash
# Monitor DHCP & ProxyDHCP traffic (Ports 67 & 68)
sudo tcpdump -i any -nnv port 67 or port 68

# Monitor TFTP boot file transfers (Port 69)
sudo tcpdump -i any -nnv port 69
```

---



