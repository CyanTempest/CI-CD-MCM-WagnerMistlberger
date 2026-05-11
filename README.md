![CI Status](https://github.com/CyanTempest/CI-CD-MCM-WagnerMistlberger/actions/workflows/ci.yml/badge.svg?branch=exercise%2F02-microservice-docker)
Sonar Test

# Continuous Delivery in Agile Software Development -- Exercises

This repository contains four progressive exercises for the Master course **Continuous Delivery in Agile Software Development**.

## Overview

Throughout this course, you will build and evolve a **Product Catalog API**—a RESTful web service for managing products. The project grows in complexity with each exercise, moving from basic Git usage to a fully orchestrated Kubernetes deployment.

| Exercise | Topic | Branch |
| --- | --- | --- |
| **1** | **Git Basics:** PRs, Interactive Rebase, Unit Tests | `exercise/01-git-basics` |
| **2** | **Microservices & Docker:** Multi-stage builds & Compose | `exercise/02-microservice-docker` |
| **3** | **CI Pipeline:** SonarCloud, Matrix Builds, Linting | `exercise/03-ci-pipeline` |
| **4** | **Security & K8s:** Vulnerability Scanning & Kubernetes | `exercise/04-security-k8s` |

## Technology Stack

* **Language:** Go 1.24+
* **Database:** PostgreSQL
* **Containerization:** Docker & Docker Compose
* **CI/CD:** GitHub Actions
* **Quality/Security:** SonarCloud, golangci-lint, Trivy
* **Deployment:** Kubernetes (Minikube)

## Prerequisites

* Go 1.24+ installed locally
* Git 2.30+
* A GitHub Account
* Docker Desktop installed
* Minikube (required for Exercise 4)

## Getting Started

1. **Fork** this repository to your own GitHub account. **Uncheck** "Copy the `main` branch only" to ensure you get all exercise branches.
2. **Clone** your fork:
```bash
git clone https://github.com/<your-username>/CI-CD-MCM.git
cd CI-CD-MCM

```


3. **Switch** to the branch of the exercise you are currently working on:
```bash
git checkout exercise/01-git-basics

```



> **Note:** Each branch contains its own specific `README.md` with detailed instructions and tasks for that particular exercise.

## Author

* **FH-Prof. Dr. Marc Kurz** (marc.kurz@fh-hagenberg.at)
