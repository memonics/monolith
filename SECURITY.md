# Security Policy

## Supported Versions

We actively monitor and patch architectural vulnerabilities for the following baseline versions:


| Version | Supported |
| ------- | --------- |
| v1.5.x  | ✅ Yes     |
| < v1.5  | ❌ No      |


## Reporting a Vulnerability

**CRITICAL: Do not open public GitHub Issues for security-critical exploits.** If you discover a structural bypass, such as a configuration leak that exposes Ring 0 credentials to a Ring 2 public interface, or an exploit vector that corrupts thread synchronization inside the core service engines, please report it privately.

Please transmit your technical breakdown and minimal reproduction steps directly to **[bdralite@gmail.com](mailto:bdralite@gmail.com)**.

### Our Disclosure Commitment

- We will acknowledge receipt of your vulnerability report within 48 business hours.
- We will provide an estimated timeline for an architectural patch or remediation path.
- Once a resolution is compiled and validated through our AST linters, we will release a security advisory patch and properly credit your contributions to the safety of the ecosystem.

