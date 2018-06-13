pipeline {
  agent {
    docker {
      image 'golang:1.10-alpine'
      args '-v /Users/jean-marcelbelmont/jenkins_data'
    }

  }
  stages {
    stage('Build Information') {
      steps {
        sh '''go version

go fmt'''
      }
    }
    stage('Run Tests') {
      steps {
        sh 'go test'
      }
    }
  }
}