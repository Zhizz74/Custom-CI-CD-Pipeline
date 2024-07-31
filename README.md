CI / CD pipeline with local runner on GitHub Actions that builds, deploys, and pushes docker images to Docker Hub.

# Usage

1. Fork repository. 
2. Download actions-runner from GitHub Actions.
3. Register runner using your runner token
   ```./config.sh --url forked_repository_url --token token```
5. Start up runner in local folder using
   ```./run.sh```
6. Start up docker on your device.
7. Commit changes to production branch.
