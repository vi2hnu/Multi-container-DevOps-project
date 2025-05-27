# Multi-container-DevOps project

## Overview
This project is a URL shortener web application featuring a fully automated CI/CD pipeline that handles building, and deployment processes. The application is deployed on an AWS EC2 instance, providing scalable and secure hosting. This setup ensures efficient development workflows and delivers a seamless user experience. Main focus of this project is to showcase the devops cycle of a project.

![Image](https://github.com/user-attachments/assets/d1d25b7b-e25e-4444-8206-758721f3f8f2)

## Website

![Image](https://github.com/user-attachments/assets/f0f69831-98d9-4cf2-b77d-b8885644c645)

### Tech Stack

- **React**: Frontend.
- **GoLang**: Backend.
- **Redis**: Efficiently caches the most frequently accessed original long URLs to reduce memory usage and storage by avoiding duplicate URL shortening.
- **MongoDB**: Database that uses indexed queries to efficiently retrieve original URLs.


## CI Pipeline

Continous Integration pipeline triggered for every commit on main branch and checks for integration error before deploying.

### Steps:

1. Log in to GitHub Container Registry for publishing latest docker images
2. Dockerize, build, and push the latest backend Docker image.
3. Dockerize, build, and push the latest frontend Docker image.

``` yaml
steps:
    - uses: actions/checkout@v4

    - name: Log in to GitHub Container Registry
      uses: docker/login-action@v3
      with:
        registry: ghcr.io
        username: vi2hnu
        password: ${{ secrets.GH_PAT }}

    - name: Build & push backend
      run: |
        cd backend
        docker build . --tag ghcr.io/vi2hnu/devops_project-backend:${{ vars.VERSION_TAG}}
        docker push ghcr.io/vi2hnu/devops_project-backend:${{ vars.VERSION_TAG}}

    - name: Build & push frontend
      run: |
        cd frontend
        docker build . --tag ghcr.io/vi2hnu/devops_project-frontend:${{ vars.VERSION_TAG}}
        docker push ghcr.io/vi2hnu/devops_project-frontend:${{ vars.VERSION_TAG}}
```


## CD Pipeline
The Continuous Delivery pipeline is triggered after the CI pipeline to deploy the latest images to the EC2 instance. Each commit initiates the CI pipeline, which then triggers the CD pipeline, performing all deployment tasks directly on the EC2 instance. Docker Compose is used to manage and run multiple containers. In this project, the EC2 instance also functions as a self-hosted GitHub runner.


### Steps:
1. Create backend/.env if it doesnt exists
2. Stop running containers
3. Replace version tag in docker-compose
4. Pull latest docker image
5. Delete old docker image without any tags
6. Start up docker service by using docker compose

``` yaml
steps:
    - uses: actions/checkout@v4
 
    - name: Create backend/.env
      run: |
        mkdir -p backend
        echo "${{ secrets.ENV_SECRECT }}" > backend/.env
        
    - name: Stop running containers
      run: sudo docker compose down

    - name: Replace version tag in docker-compose
      run: |
        sed -i "s/\${VERSION_TAG}/${{ vars.VERSION_TAG }}/g" docker-compose.yml

    - name: Pull latest docker image
      run: sudo docker compose pull

    - name: Delete old docker image
      run: sudo docker image prune -f

    - name: Start up docker service
      run: sudo docker compose up -d

```

## DevOps Operations Done in this Project

| Techniques              | Status                               |
| ------------------------| -------------------------------------|
| Continuous Development  | Done by Developer                    |
| Continuous Integration  | Done via Github actions              |
| Continuous Testing      | Not Done                             |
| Continuous Delivery     | Done via Github actions              |
| Continuous Monitoring   | Handled through AWS Console for EC2  |
| Continuous Feedback     | Provided via AWS Console for EC2     |
| Continuous Operations   | Managed on AWS EC2 instance          |
