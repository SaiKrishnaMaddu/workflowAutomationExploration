# Workflow orchestration and task scheduling tools
- Workflow orchestration and task scheduling tools are software platforms that allow users to design, automate, and manage complex sequences of tasks across different systems
- essentially coordinating multiple operations within a larger workflow by scheduling them to run at specific times or based on dependencies; popular options include: 
    - **Apache Airflow** - opensource
    - **AWS Step Functions** - specific to aws
    - **Camunda** - opens source
    - **Prefect**
    - **Dagster** - opensource
    - **Microsoft Power Automate** - specific to microsoft
    - **Google Workflows** - specific to google
    - **Argo** - opensource container native
    - **Temporal** - open source but it has temporal cloud too
    - **Control M** - not an opensource too expensive
    - **Cadence** - opensource
- These tools enable users to define dependencies between tasks, set execution schedules, monitor progress, handle errors, and trigger actions based on specific conditions
- Scheduling is a sequential process that involves triggering tasks at specific times, while orchestration is a workflow-level process that involves managing and running services.


### Which has golag support
- Temporal has SDK
- Cadence has SDK
- Argo has SDK

### API support
- AWS step up functions
- Google workflows
- Control - M

### Python based but can be triggered using APIs
- Apcahe workflow
- Perfect
- Dagster

### Java based but can be triggered using APIs
- Camuda

### platform Api
- Microsoft power platform


### How to interact
- using SDKs will be good because
    - Easier to implement & maintain.
    - Direct function calls (low latency).

- using APIs
    - Works with all tools (even without Go SDKs).
    - More flexible, supports multiple languages
    - Higher latency than SDK calls

## Tool too pick up
    - ✅ State management (resumes tasks after failure).
    - ✅ Fault tolerance (handles failures, retries).
    - ✅ Highly scalable (supports millions of tasks).
    - ✅ Open-source (no vendor lock-in).
    - ✅ Cloud-agnostic (not tied to AWS, GCP, Azure).

## **Temporal**
- Temporal stands out as the strongest option across all your requirements:
- **State management**: Industry-leading persistent state tracking that can resume workflows exactly where they left off, even after long outages
- **Fault tolerance**: Exceptional error handling with configurable retry policies, timeout mechanisms, and automatic recovery
- **Scalability**: Designed to handle millions of workflows and tasks with a distributed architecture
_ **Open-source**: Fully open-source under MIT license
- **Cloud-agnostic**: Runs anywhere - on-premises or any cloud provider

## **Apache Airflow**
- **State management**: Tracks task states in a database and can resume failed workflows
- **Fault tolerance**: Provides retries, error handling, and SLA monitoring
- **Scalability**: Scales to millions of tasks with Celery or Kubernetes executors
- **Open-source**: Apache 2.0 licensed with active community
- **Cloud-agnostic**: Deployable anywhere, though requires more configuration for large-scale setups

# Best will be temporal
**Solution Architecture**
- Core Setup
    - **Temporal Server**: Deploy as your central orchestration engine
    - **PostgreSQL Integration**: Use your existing PostgreSQL for Temporal's persistence store
    - **Worker Services**: Deploy horizontally scalable worker nodes to process tasks
- Policy Creation Event → Temporal Workflow Starts
    - When a new insurance policy is created, a Temporal workflow is triggered to schedule recurring invoices, notifications, and payment tasks.
- Recurring Monthly Invoicing
    - Temporal cron workflows generate invoices every month for each policyholder.
    - Uses PostgreSQL to store invoice details and ensure idempotency.
- Handling Other Workflows
    - Document Requests → Asynchronous Temporal activities fetch requested documents.
    - Pending Cancellations → Automated workflows check pending cases & send reminders.
    - Payment Processing → Ensures payments are collected, retries if necessary.
    - Cancellation Notices → Sends alerts before policy cancellation.
- Scaling for 3M Requests on Peak Days
    - Deploy multiple Temporal Worker nodes to handle parallel invoice generation.
    - Use queue-based execution (Temporal’s Task Queues) to distribute workloads.
    - Use PostgreSQL sharding if needed for better performance.

