+++ 
draft = false
date = 2026-10-31T12:58:14+01:00
title = "Kubernetes & GitOps Infrastructure Orchestration with Microservices"
description = "This repository demonstrates an enterprise-grade, production-ready microservices architecture orchestrated on **Kubernetes (Rancher/K3s)**, provisioned via **Terraform (IaC)**, and continuously deployed using **ArgoCD (GitOps)**."
slug = ""
authors = []
tags = ["kubernetes", "linux", "terraform", "argoCD", "git", "go", "python", "gitOps","database",""]
categories = []
externalLink = ""
series = []
+++


# 🚀 Production-Ready Kubernetes & GitOps Infrastructure Setup

![Kubernetes](https://img.shields.io/badge/Kubernetes-v1.28%2B-blue?style=flat-square&logo=kubernetes)
![Terraform](https://img.shields.io/badge/Terraform-Infrastructure%20as%20Code-purple?style=flat-square&logo=terraform)
![ArgoCD](https://img.shields.io/badge/ArgoCD-GitOps-orange?style=flat-square&logo=argocd)
![Rancher](https://img.shields.io/badge/Rancher-2.9.3-lightblue?style=flat-square&logo=rancher)
![Golang](https://img.shields.io/badge/Golang-Microservices-00ADD8?style=flat-square&logo=go)

This repository demonstrates an enterprise-grade, production-ready microservices architecture orchestrated on **Kubernetes (Rancher/K3s)**, provisioned via **Terraform (IaC)**, and continuously deployed using **ArgoCD (GitOps)**.

---
## Code/Github 

[mozkaya1@github.com](https://github.com/mozkaya1/kubernetes-production-setup.git)

## 🏗️ Architecture & DevOps Workflow

```text
[ Git Repository ] 
       │ (Push/PR)
       ▼
[ ArgoCD (GitOps) ] ──(Sync)──► [ Kubernetes Cluster (Rancher / K3s) ]
                                      │
        ┌─────────────────────────────┼─────────────────────────────┐
        ▼                             ▼                             ▼
[ Traefik Ingress ]          [ Go Microservices ]          [ Databases ]
  ├── Live Portal              ├── Weather API               ├── MariaDB
  ├── Meira Accessories        ├── Currency API              └── PostgreSQL
  └── Musti App                └── Real-time WebSocket
```

---

## 📦 Deployed Microservices & Workloads

1. **Live Go Portal**
   - **Type:** Real-time dashboard frontend with WebSocket communication streaming live data.
2. **Weather API**
   - **Type:** Go REST API providing real-time weather data and health checks.
3. **Real Currency API**
   - **Type:** Multi-asset currency JSON data provider built in Go.
4. **Meira Accessories**
   - **Type:** E-commerce web application integrated with a dedicated **MariaDB** database layer.
5. **Musti App**
   - **Type:** Web application accessible internally via ClusterIP and externally via NodePort/Ingress.
6. **Observability Stack**
   - **Prometheus:** Metrics collection and resource monitoring.
   - **Grafana:** Visual dashboards for cluster performance and health tracking.

---

## 🛠️ Technical Stack

- **Orchestration & Management:** Kubernetes v1.28+, Rancher 2.9.3, K3s
- **Infrastructure as Code (IaC):** Terraform
- **Continuous Delivery & GitOps:** ArgoCD
- **Networking & Ingress:** Traefik Ingress Controller
- **Backend Services:** Golang (REST APIs, WebSockets), Python, Bash
- **Databases:** MariaDB, PostgreSQL
- **Monitoring:** Prometheus, Grafana

---

## 📂 Repository Structure

```text
kubernetes-production-setup/
├── terraform/          # Infrastructure provisioning scripts & modules
├── argocd/             # GitOps application manifests and sync policies
├── manifests/          # Raw Kubernetes YAML manifests
├── helm-charts/        # Reusable Helm charts for microservices
├── documentation/      # Architecture blueprints and runbooks
└── scripts/            # Deployment, automation, and maintenance scripts
```

---

## 🚀 Quick Start & Deployment

### 1. Provision Infrastructure (Terraform)
```bash
cd terraform/
terraform init
terraform apply
```

### 2. GitOps Continuous Deployment (ArgoCD)
```bash
# Apply ArgoCD application definitions to manage cluster state
kubectl apply -f argocd/
```

### 3. Manual Deployment (Alternative)
```bash
# Deploy raw manifests
kubectl apply -f manifests/

# Or deploy via Helm charts
helm install meira-accessory ./helm-charts/web-app
helm install weather-api ./helm-charts/go-api-app
```

---

## 📊 Visual Dashboards

### Rancher Cluster Management
![Rancher](/rancher.png)

### Grafana Observability
![Grafana](/monitoring.png)

### ArgoCD Continuous Deployment
![ArgoCD](/argocd.png)

### Kubernetes PODS 
![Pods](/pods.png)

### Kubernetes Service 
![Pods](/services.png)


