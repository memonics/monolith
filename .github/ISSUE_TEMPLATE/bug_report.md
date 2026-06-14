---
name: 🐛 Bug Report
about: File a structured bug report regarding a layer boundary or runtime failure
title: '[BUG]: '
labels: bug, unverified
assignees: ''
---

## 🔍 Bug Description
A clear and concise description of what the runtime anomaly is.

## 🗺️ Architectural Context
Please isolate where the issue is occurring within the BDRA Lite topology:
* **Target Ring:** [e.g., Ring 0, Ring 1, Ring 2]
* **Target Layer:** [e.g., Pure, Protected, Public, Health]

## 🧪 Steps to Reproduce
1. Start the backing sandbox infrastructure via `make up`
2. Execute the following sequence or endpoint payload:
3. See error condition: 

## 📋 Expected Behavior
What should have happened according to the BDRA Lite specification? (e.g., "The Public layer should have gracefully caught the storage fault and triggered Tier 2 transparent cache lookup instead of panicking.")

## 💻 Environment Setup
* **Go Version:** [e.g., 1.22]
* **OS:** [e.g., Windows 11, macOS Sequoia, Ubuntu]
* **Backing DB/Cache Version:** [e.g., Postgres 15, Redis 7]