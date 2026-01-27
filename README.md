# MPaisa
Personal Finance Tracker App

## AWS Deployment Overview

This project is designed to run as a containerized Go backend with PostgreSQL on AWS.

### RDS PostgreSQL (recommended for production-like setups)

1. Create an RDS PostgreSQL instance in the same region as your compute (for example, `ap-south-1`).
2. Use a free-tier eligible instance class such as `db.t3.micro` with at least 20 GB of storage.
3. Place the instance in private subnets and restrict its security group to only allow inbound traffic on port `5432` from your application compute (EKS worker nodes or EC2 instance running the container).
4. Construct a connection string in the form:

   `postgresql://<user>:<password>@<rds-endpoint>:5432/MPaisa?sslmode=require`

5. Store this connection string (or its individual components) in AWS Secrets Manager and surface it to the application through the `DB_SOURCE` environment variable.
6. Ensure automated backups are enabled and set a reasonable retention period for learning/testing environments.
