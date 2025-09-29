# 🌟 Most Useful AWS Services + Real-world Use Cases

## 📧 Amazon Simple Email Service (SES)

Scalable email service for sending transactional and marketing emails.  
**Use case:** Send OTPs, password reset links, or bulk promotional emails.

## 📢 Amazon Simple Notification Service (SNS)

Pub/Sub messaging to fanout notifications across multiple subscribers.  
**Use case:** Send push notifications to mobile users when a new order is placed.

## 📩 Simple Queue Service (SQS)

Reliable message queuing for decoupling microservices.  
**Use case:** Queue payment requests from an e-commerce site for async processing.

## ⚡ ElastiCache

In-memory cache (Redis/Memcached) for fast data access.  
**Use case:** Store frequently accessed product data or session tokens to reduce DB load.

## 📑 Amazon DocumentDB

Managed, MongoDB-compatible NoSQL database.  
**Use case:** Store JSON-based user profiles or IoT device data.

## 🗄️ Aurora and RDS

Relational databases with high availability and automated backups.  
**Use case:** Power transactional systems like banking, bookings, or ERP apps.

## 🛡️ CloudTrail

Logs all AWS API calls for auditing and compliance.  
**Use case:** Detect unauthorized access to critical AWS resources.

## ⏱️ CloudWatch

Monitoring and alerting for logs and metrics.  
**Use case:** Trigger an alert if CPU usage of EC2 spikes above 80%.

## ☸️ Elastic Kubernetes Service (EKS)

Managed Kubernetes for container orchestration.  
**Use case:** Run scalable microservices-based fintech or SaaS applications.

## 🚀 Elastic Container Service (ECS)

Fully managed container orchestration without Kubernetes complexity.  
**Use case:** Deploy a fleet of Dockerized background workers.

## 📦 Elastic Container Registry (ECR)

Private Docker image repository.  
**Use case:** Store and version your microservices’ Docker images for CI/CD pipelines.

## 📂 S3 (Simple Storage Service)

Durable, scalable object storage.  
**Use case:** Host static websites, store user-uploaded photos/videos, or backup logs.

## 💻 EC2 (Elastic Compute Cloud)

An EC2 instance is a virtual server in the cloud that comes with configurable CPU, memory (RAM), storage (hard disk/SSD), and operating system.
It gives you the flexibility to choose the right compute resources for your workload.

**Use case:** Run APIs, backend services, or ML model inference servers.
When launching an EC2 instance, you need to configure several settings that define its compute power, storage, security, and networking.

🖥️ Instance Type

Select the hardware capacity such as CPU, memory, and network performance.
👉 Example: t2.micro for free tier, m5.large for production workloads.

💿 AMI (Amazon Machine Image)

Choose the operating system and pre-installed software (Linux, Windows, Mac, etc.).
👉 Example: Ubuntu Server 22.04, Amazon Linux, Windows Server.

📦 Storage

Configure the type and size of storage (EBS volumes, SSD, or HDD).
👉 Example: 30GB General Purpose SSD for web apps.

🔒 Security Groups

Set up firewall rules to control inbound and outbound traffic.
👉 Example: Allow port 22 for SSH, port 80/443 for HTTP/HTTPS.

🔑 Key Pair

Create or use an existing key pair for SSH access into the instance.Like login on your computer.
👉 Example: Download .pem file and connect via terminal.

🌐 Network Settings

Configure VPC, subnet, and IP addresses (public/private).
👉 Example: Place instance in a private subnet with NAT Gateway for security.

🛡️ IAM Role

Attach an IAM role to give the instance permissions to access other AWS resources.
👉 Example: Grant S3 read access without embedding credentials.

⚙️ User Data

Add bootstrap scripts to run automatically when the instance starts.
👉 Example: Install Apache, set environment variables, configure apps.

📍 Elastic IP

Optionally associate a static IP address for consistent public access.
👉 Example: Assign an Elastic IP so your app server always has the same IP.

## 🔐 IAM (Identity and Access Management)

Used to securely manage users, groups, and roles, assign fine-grained permissions for accessing AWS services, and generate access keys, secret key for programmatic access in applications.

---

# 🛒 E-commerce App on AWS (High-Level Architecture)

```plaintext
                     ┌───────────────┐
                     │   Customers   │
                     └───────┬───────┘
                             │
                      HTTPS Requests
                             │
                       ┌─────▼─────┐
                       │   Route53 │  (DNS & domain routing)
                       └─────┬─────┘
                             │
                     ┌───────▼────────┐
                     │   CloudFront   │  (CDN caching static assets from S3)
                     └───────┬────────┘
                             │
                   ┌─────────▼─────────┐
                   │   ALB / API GW    │  (Load balancer / API Gateway)
                   └─────────┬─────────┘
                             │
                   ┌─────────▼─────────┐
                   │      EC2/ECS/EKS  │  (Backend services: Go, Node.js, etc.)
                   └───────┬──┬──┬─────┘
                           │  │  │
          ┌────────────────┘  │  └───────────────────┐
          │                   │                      │
   ┌──────▼─────┐      ┌──────▼─────┐        ┌───────▼──────┐
   │   RDS/Aurora│      │ ElastiCache│        │   SQS/SNS    │
   │ (User+Order │      │   (Caching)│        │ Messaging    │
   │   Database) │      └────────────┘        └───────┬──────┘
   └─────────────┘                                   │
                                                     │
                                              ┌──────▼───────┐
                                              │   Workers    │ (Async tasks: payment, emails)
                                              └──────┬───────┘
                                                     │
                                             ┌───────▼─────────┐
                                             │   SES (Emails)  │
                                             └─────────────────┘
```
