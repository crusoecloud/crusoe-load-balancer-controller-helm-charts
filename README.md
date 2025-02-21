# Crusoe Load Balancer Controller Helm Chart

This repository contains the **official Crusoe Load Balancer Controller Helm Chart** for use with [Crusoe Managed Kubernetes (CMK)](https://www.crusoecloud.com).  
**Currently, the Crusoe Load Balancer Controller is in :construction: Alpha :construction:.**

---

## Overview

The **Crusoe Load Balancer Controller** automates the provisioning and management of Crusoe Cloud load balancers for Kubernetes `Service` objects of type `LoadBalancer`.  
This Helm chart is:

- **Supported only on CMK**. This guide assumes that the user has already set up CMK on Crusoe Cloud. Other configurations will be supported on a best-effort basis.

---

## Prerequisites

1. A [Crusoe Managed Kubernetes (CMK)](https://www.crusoecloud.com) cluster.
2. [Helm](https://helm.sh/docs/intro/install/) installed on your local machine or CI environment.

---

## Installation

### 1. Clone or Acquire the Helm Chart
Ensure you have the Helm chart for the Crusoe LB Controller available locally.

### 2. Configure the VPC ID
Edit your `values.yaml` file and set the `CRUSOE_VPC_ID` under `controller.env`:

```yaml
controller:
  env:
    CRUSOE_VPC_ID: "your-vpc-id-here"
```

This VPC ID must match the VPC where your backend nodes reside.

### 3. Choose the Namespace
We **strongly recommend** deploying the LB controller in the `crusoe-system` namespace. The service account secrets required to interact with the Crusoe Load Balancer API are stored in this namespace by default.  

If you wish to deploy the controller in a different namespace, you must copy the necessary secrets (`crusoe-secrets`) into that namespace.

### 4. Install with Helm
Navigate to the chart directory and install the controller:

```bash
helm install crusoe-lb-controller ./crusoe-lb-controller --namespace crusoe-system
```

### 5. Verify Installation
Check the installed release:

```bash
helm list -n crusoe-system
```

You should see something like:

```
NAME                	NAMESPACE    	REVISION	UPDATED                             	STATUS  	CHART                      	APP VERSION
crusoe-lb-controller	crusoe-system	1       	2025-02-06 16:10:54.608382 -0800 PST	deployed	crusoe-lb-controller-0.0.23	v0.0.23
```

Verify that the controller pod is running:

```bash
kubectl get pods -n crusoe-system -l app=crusoe-lb-controller
```