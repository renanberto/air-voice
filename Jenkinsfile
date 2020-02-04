pipeline {
  agent {
    node {
      label 'master'
    }

  }
  stages {
    stage('ShellScript') {
      parallel {
        stage('ShellScript') {
          steps {
            sh '''#!/bin/bash

echo "Olá"'''
          }
        }

        stage('ShellScript2') {
          steps {
            sh '''#!/bin/bash

echo "Olá 2"'''
          }
        }

      }
    }

  }
}