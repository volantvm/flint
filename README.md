# 🌀 Flint — KVM Management, Reimagined

<p align="center">
  <img src="https://i.ibb.co/yj2bFZG/flint-banner.jpg" alt="Flint Logo" width="300"/>
</p>

<p align="center">
  <strong>
    A single &lt;11MB binary with a modern Web UI, CLI, and API for KVM.
    <br/>No XML. No bloat. Just VMs.
  </strong>
</p>

<p align="center">
  <a href="https://github.com/ccheshirecat/flint/releases/latest">
    <img src="https://img.shields.io/github/v/release/ccheshirecat/flint" alt="Latest Release">
  </a>
  <a href="https://github.com/ccheshirecat/flint/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/ccheshirecat/flint" alt="License">
  </a>
  <a href="https://github.com/ccheshirecat/flint/actions/workflows/release.yml">
    <img src="https://img.shields.io/github/actions/workflow/status/ccheshirecat/flint/.github/workflows/release.yml" alt="Build Status">
  </a>
</p>

---
![Flint Dashboard](https://i.ibb.co/wN9H8WKX/Screenshot-2025-09-07-at-3-51-58-AM.png)
![Flint Library](https://i.ibb.co/Z1k9XBqQ/Screenshot-2025-09-08-at-4-59-46-AM.png)


Flint is a modern, self-contained KVM management tool built for developers, sysadmins, and home labs who want zero bloat and maximum efficiency. It was built in a few hours out of a sudden urge for something better.

---

### 🚀 One-Liner Install

**Prerequisites:** A Linux host with `libvirt` and `qemu-kvm` installed.

```bash
curl -fsSL https://raw.githubusercontent.com/ccheshirecat/flint/main/install.sh | bash
```
*Auto-detects OS/arch, installs to `/usr/local/bin`, and prompts for web UI passphrase setup.*

---

### 🔐 Security & Authentication

Flint implements a multi-layered security approach:

**Web UI Security:**
- **Passphrase Authentication**: Web interface requires a passphrase login
- **Session-Based**: Secure HTTP-only cookies with 1-hour expiry
- **No API Key Exposure**: Web UI never exposes API keys to browsers

**API Security:**
- **Bearer Token Authentication**: CLI and external tools use API keys
- **Protected Endpoints**: All API endpoints require authentication
- **Flexible Access**: Support for both session cookies and API keys

**Authentication Flow:**
```bash
# First run - set passphrase
flint serve
# 🔐 No web UI passphrase set. Let's set one up for security.
# Enter passphrase: ********

# Web UI access
# Visit http://your-server:5550 → Enter passphrase → Full access

# CLI access (uses API key)
flint vm list --all

# External API access
curl -H "Authorization: Bearer YOUR_API_KEY" http://localhost:5550/api/vms
```

---

### ✨ Core Philosophy

-   🖥️ **Modern UI** — A beautiful, responsive Next.js + Tailwind interface, fully embedded.
-   ⚡ **Single Binary** — No containers, no XML hell. A sub-8MB binary is all you need.
-   🛠️ **Powerful CLI & API** — Automate everything. If you can do it in the UI, you can do it from the command line or API.
-   📦 **Frictionless Provisioning** — Native Cloud-Init support and a simple, snapshot-based template system.
-   🔐 **Secure by Default** — Multi-layered authentication with passphrase protection.
-   💪 **Non-Intrusive** — Flint is a tool that serves you. It's not a platform that locks you in.

---

### 🏎️ Quickstart

**1. Start the Server**
```bash
# Interactive setup (recommended for first run)
flint serve --set-passphrase

# Or set passphrase directly
flint serve --passphrase "your-secure-password"

# Or use environment variable
export FLINT_PASSPHRASE="your-secure-password"
flint serve
```
*On first run, you'll be prompted to set a web UI passphrase for security.*
*   **Web UI:** `http://localhost:5550` (requires passphrase login)
*   **API:** `http://localhost:5550/api` (requires authentication)

**2. Web UI Access**
- Visit `http://localhost:5550`
- Enter your passphrase to access the management interface
- All API calls are automatically authenticated via session

**3. CLI Usage**
```bash
# VM Management
flint vm list                    # List all VMs
flint vm launch my-server        # Create and start a VM
flint vm ssh my-server          # SSH into a VM

# Cloud Images
flint image list                 # Browse cloud images
flint image download ubuntu-24.04 # Download an image

# Networks & Storage
flint network list               # List networks
flint storage volume list default # List storage volumes
```

**4. API Access (for external tools)**
```bash
# Get your API key (requires authentication)
curl -H "Authorization: Bearer YOUR_API_KEY" http://localhost:5550/api/vms
```
---

### 📖 Full Documentation

Complete CLI commands, API reference, and advanced usage:

➡️ **[docs.md](docs.md)** - Complete CLI & API Documentation

---

### 🔧 Tech Stack

-   **Backend:** Go 1.25+
-   **Web UI:** Next.js + Tailwind + Bun
-   **KVM Integration:** libvirt-go
-   **Binary Size:** ~11MB (stripped)

---

<p align="center">
  <b>🚀 Flint is young, fast-moving, and designed for builders.<br/>
  Try it. Break it. Star it. Contribute.</b>
</p>
